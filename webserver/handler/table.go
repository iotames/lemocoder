package handler

import (
	"lemocoder/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTable(c *gin.Context) {
	table := database.DataTable{}
	b := c.Bind(&table)
	if b != nil {
		ErrorArgs(c)
		return
	}
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
		ErrorArgs(c)
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
	pageID := c.GetInt64("page_id")
	table := database.DataTable{PageID: int64(pageID)}
	has, err := database.GetModel(&table)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail(err.Error(), 500))
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	c.JSON(http.StatusOK, Response(table, "success", 200))
}
