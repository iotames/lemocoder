package handler

import (
	"lemocoder/database"
	"log"
	"strconv"

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
	pageIDstr := c.DefaultQuery("page_id", "0")
	pageID, _ := strconv.ParseInt(pageIDstr, 10, 64)
	log.Println("---GetTable---", pageID)
	// table := database.DataTable{}
	// has, err := database.GetModelWhere(&table, "page_id = ?", pageID)
	table := database.DataTable{PageID: pageID}
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
