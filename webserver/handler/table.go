package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"lemocoder/database"
	gen "lemocoder/generator"
	"lemocoder/model"
	"lemocoder/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormTableSchema struct {
	ID, PageID, Name, Title, Remark string
	model.TableSchema
}

func CreateTable(c *gin.Context) {
	postData, err := ParsePostData(c)
	if err != nil {
		return
	}
	log.Println("-----------Log--postData--------------", postData)
	pageIDstr, ok := postData["PageID"]
	if !ok {
		ErrorArgs(c, errors.New("缺少PageID参数"))
		return
	}
	if pageIDstr.(string) == "" {
		ErrorArgs(c, errors.New("参数PageID不能为空"))
		return
	}
	pageID, _ := strconv.ParseInt(pageIDstr.(string), 10, 64)

	table := database.DataTable{}
	has, err := database.GetModelWhere(&table, "page_id = ?", pageID)
	fmt.Printf("\n--has(%+v)---err(%+v)---\n", has, err)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if has {
		c.JSON(http.StatusOK, ResponseFail("该页面已包含一个数据表格，请勿重复添加", http.StatusBadRequest))
		return
	}

	table.PageID = pageID
	err = setDataTable(postData, &table, true)
	if err != nil {
		ErrorServer(c, err)
		return
	}

	log.Printf("----Log--tableModel-<%+v>-\n", table)

	_, err = database.CreateModel(&table)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	wpage := database.WebPage{}
	wpage.ID = pageID
	database.UpdateModel(&wpage, map[string]interface{}{"state": database.PAGE_STATE_READY})

	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

func UpdateTable(c *gin.Context) {
	postData, err := ParsePostData(c)
	if err != nil {
		return
	}
	if postData.GetID() == 0 {
		ErrorArgs(c, errors.New("ID不能为空"))
		return
	}
	modelFind := database.DataTable{}
	modelFind.ID = postData.GetID()

	err = mustFind(c, &modelFind)
	if err != nil {
		return
	}

	err = setDataTable(postData, &modelFind, false)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	_, err = database.UpdateModel(&modelFind, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

// 生成源代码文件
func CreateTablePageCode(c *gin.Context) {
	postData, err := ParsePostData(c)
	if err != nil {
		return
	}
	pageIDstr, ok := postData["PageID"]
	if !ok {
		ErrorArgs(c, errors.New("缺少PageID参数"))
		return
	}
	if pageIDstr.(string) == "" {
		ErrorArgs(c, errors.New("参数PageID不能为空"))
		return
	}
	pageID, _ := strconv.ParseInt(pageIDstr.(string), 10, 64)
	table := database.DataTable{}
	has, err := database.GetModelWhere(&table, "page_id = ?", pageID)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		c.JSON(http.StatusOK, ResponseFail("该页面不包含数据表格", http.StatusBadRequest))
		return
	}
	page := database.WebPage{}
	page.ID = pageID
	has, err = database.GetModel(&page)
	if !has {
		logger := util.GetLogger()
		logger.Error("Error for CreateCode: Not Found WebPage: ", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	if page.State == database.PAGE_STATE_CREATED {
		c.JSON(http.StatusOK, ResponseFail("请勿重复生成代码", 400))
		return
	}

	// 生成源代码文件 BEGIN
	t, err := table.GetStructSchema()
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for GetStructSchema:", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}

	err = gen.CreateTableClient(t, page)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for CreateTableClient:", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	err = gen.CreateTableServer(t)
	if err != nil {
		logger := util.GetLogger()
		logger.Error("Error for CreateTableClient:", err)
		c.JSON(200, ResponseFail(err.Error(), 500))
		return
	}
	// 生成源代码文件 END

	database.UpdateModel(&page, map[string]interface{}{"state": database.PAGE_STATE_CREATED})
	c.JSON(http.StatusOK, ResponseOk("创建代码成功, 请进行后续操作!"))
}

func setDataTable(postData PostData, dtable *database.DataTable, isCreate bool) error {
	// TODO 判断路由, DataTypeName 是否重复。
	tschema, err := dtable.GetStructSchema()
	if err != nil {
		return err
	}
	tschemaJson, err := json.Marshal(postData["StructSchema"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(tschemaJson, &tschema) // POST tschemaJson 缺少部分字段, 不改变 &tschema 对应字段的原有值
	if err != nil {
		return err
	}
	if tschema.RowKey == "" {
		tschema.RowKey = "ID"
	}
	if tschema.ItemsDataUrl == "" {
		tschema.ItemsDataUrl = fmt.Sprintf("/api/%s/list", database.ObjToTableCol(tschema.ItemDataTypeName))
	}
	if tschema.ItemCreateUrl == "" {
		tschema.ItemCreateUrl = fmt.Sprintf("/api/%s/create", database.ObjToTableCol(tschema.ItemDataTypeName))
	}
	if tschema.ItemUpdateUrl == "" {
		tschema.ItemUpdateUrl = fmt.Sprintf("/api/%s/update", database.ObjToTableCol(tschema.ItemDataTypeName))
	}
	if tschema.ItemDeleteUrl == "" {
		tschema.ItemDeleteUrl = fmt.Sprintf("/api/%s/delete", database.ObjToTableCol(tschema.ItemDataTypeName))
	}
	if postData["Name"] != nil {
		dtable.Name = postData["Name"].(string)
	}
	if postData["Title"] != nil {
		dtable.Title = postData["Title"].(string)
	}
	if postData["Remark"] != nil {
		dtable.Remark = postData["Remark"].(string)
	}
	for i, v := range tschema.Items {
		tschema.Items[i].DataName = database.TableColToObj(v.DataName)
	}
	log.Printf("\n------tschema.ItemForms(%+v)----tschema.ToolBarForms(%+v)------\n", tschema.ItemForms, tschema.ToolBarForms)
	if len(tschema.ItemForms) == 0 {
		tschema.ItemForms = []model.ModalFormSchema{
			gen.GetUpdateForm(tschema.Items, tschema.ItemUpdateUrl),
		}
	}
	if len(tschema.ToolBarForms) == 0 {
		tschema.ToolBarForms = []model.ModalFormSchema{
			gen.GetCreateForm(tschema.Items, tschema.ItemCreateUrl),
		}
	}
	if isCreate {
		if len(tschema.ItemOptions) == 0 {
			tschema.ItemOptions = []model.TableItemOptionSchema{
				// {Key: "lineedit", Type: "edit", Title: "编辑", Url: tschema.ItemUpdateUrl}, // 添加行快捷编辑
				{Key: gen.TABLE_ITEM_OPT_KEY_UPDATE, Type: model.TABLE_ITEM_OPT_FORM, Title: "编辑", Url: tschema.ItemUpdateUrl},
			}
		}
		if len(tschema.BatchOptButtons) == 0 {
			// 添加批量操作
			tschema.BatchOptButtons = []model.BatchOptButtonSchema{
				{Title: "批量删除", Url: tschema.ItemDeleteUrl},
			}
		}
	}
	log.Println("------ItemOptions-----BatchOptButtons---------", tschema.ItemOptions, tschema.BatchOptButtons)
	return dtable.SetStructSchema(tschema)
}

func GetTable(c *gin.Context) {
	pageIDstr := c.DefaultQuery("page_id", "0")
	pageID, _ := strconv.ParseInt(pageIDstr, 10, 64)
	log.Println("---GetTable---", pageID)
	t := database.DataTable{}
	result, err := database.Query(fmt.Sprintf("SELECT * FROM %s WHERE page_id = %d", t.TableName(), pageID))
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail(err.Error(), 500))
		return
	}
	if len(result) == 0 {
		ErrorNotFound(c)
		return
	}

	resp := make(map[string]interface{}, len(result[0]))
	for k, v := range result[0] {
		nk := database.TableColToObj(k)
		if k == "struct_schema" {
			vv := model.TableSchema{}
			json.Unmarshal(v, &vv)

			resp[nk] = vv
		} else {
			resp[nk] = string(v)
		}

		log.Println(nk, string(v))
	}
	log.Println(resp)
	c.JSON(http.StatusOK, Response(resp, "success", 200))
}
