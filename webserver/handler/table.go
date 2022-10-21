package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"lemocoder/database"
	"lemocoder/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormTableSchema struct {
	ID, PageID, Name, Title, Remark string
	model.TableSchema
}

func AddTable(c *gin.Context) {
	f := FormTableSchema{}
	berr := CheckArgs(&f, c)
	if berr != nil {
		return
	}
	if f.PageID == "" {
		ErrorArgs(c, errors.New("PageID不能为空"))
		return
	}
	pageID, _ := strconv.ParseInt(f.PageID, 10, 64)
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

	fmt.Printf("\n----table.SetStructSchema(%+v)--\n", table.StructSchema)
	err = setTableSchema(f, &table)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	fmt.Printf("\n----table.SetStructSchema(%+v)--\n", table.StructSchema)
	table.PageID = pageID
	table.Name = f.Name
	table.Title = f.Title
	table.Remark = f.Remark
	_, err = database.CreateModel(&table)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

func UpdateTable(c *gin.Context) {
	f := FormTableSchema{}
	berr := CheckArgs(&f, c)
	if berr != nil {
		return
	}
	if f.ID == "" {
		ErrorArgs(c, errors.New("ID不能为空"))
		return
	}
	model := database.DataTable{}
	model.ID, _ = strconv.ParseInt(f.ID, 10, 64)
	has, err := database.GetModel(&model)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	setTableSchema(f, &model)
	model.Name = f.Name
	model.Title = f.Title
	model.Remark = f.Remark
	_, err = database.UpdateModel(&model, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

func setTableSchema(f FormTableSchema, table *database.DataTable) error {
	ts, err := table.GetStructSchema()
	if err != nil {
		return err
	}
	ts.BatchOptButtons = f.BatchOptButtons
	ts.ItemDataTypeName = f.ItemDataTypeName
	ts.ItemDeleteUrl = f.ItemDeleteUrl
	ts.ItemOptions = f.ItemOptions
	ts.ItemUpdateUrl = f.ItemUpdateUrl
	ts.Items = f.Items
	ts.ItemsDataUrl = f.ItemsDataUrl
	ts.RowKey = f.RowKey
	return table.SetStructSchema(ts)
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
	fmt.Printf("---GetTable(%+v)-----", result[0])
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
