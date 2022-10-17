package handler

import (
	"encoding/json"
	"fmt"
	"lemocoder/database"
	"lemocoder/generator"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PageID              int64                 `xorm:"notnull default(0) 'page_id'"`
// Name, Title, Remark string                `xorm:"varchar(64) notnull"`
// StructSchema        generator.TableSchema `xorm:"TEXT notnull"`
// type TableSchema struct {
// 	ItemDataTypeName, ItemsDataUrl, ItemUpdateUrl, ItemDeleteUrl, ItemCreateUrl, RowKey string
// 	Items                                                                               []TableItemSchema
// 	ItemOptions                                                                         []TableItemOptionSchema
// 	ItemForms                                                                           []ModalFormSchema
// 	ToolBarForms                                                                        []ModalFormSchema
// 	BatchOptButtons                                                                     []BatchOptButtonSchema
// }

type FormTableSchema struct {
	PageID              int64
	Name, Title, Remark string
	generator.TableSchema
}

func AddTable(c *gin.Context) {
	table := database.DataTable{}
	f := FormTableSchema{}
	b := c.Bind(&f)
	if b != nil {
		ErrorArgs(c)
		return
	}
	table.PageID = f.PageID
	table.StructSchema.BatchOptButtons = f.BatchOptButtons
	table.StructSchema.ItemDataTypeName = f.ItemDataTypeName
	table.StructSchema.ItemDeleteUrl = f.ItemDeleteUrl
	table.StructSchema.ItemOptions = f.ItemOptions
	table.StructSchema.ItemUpdateUrl = f.ItemUpdateUrl
	table.StructSchema.Items = f.Items
	table.StructSchema.ItemsDataUrl = f.ItemsDataUrl
	table.PageID = f.PageID
	table.StructSchema.RowKey = f.RowKey
	has, err := database.GetModelWhere(new(database.DataTable), "page_id = ?", table.PageID)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if has {
		c.JSON(http.StatusOK, ResponseFail("该页面已包含一个数据表格，请勿重复添加", http.StatusBadRequest))
		return
	}
	_, err = database.CreateModel(&table)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
}

func UpdateTable(c *gin.Context) {
	table := database.DataTable{}
	b := c.Bind(&table)
	if b != nil {
		fmt.Println("error for Bind:", b)
		ErrorArgs(c, b)
		return
	}
	findModel := database.DataTable{}
	findModel.ID = table.ID
	has, err := database.GetModel(&findModel)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	_, err = database.UpdateModel(&table, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("提交成功"))
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
			vv := generator.TableSchema{}
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
