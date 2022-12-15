package handler

import (
	"fmt"
	"lemocoder/database"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddWebPage(c *gin.Context) {
	pg := database.WebPage{}
	b := c.Bind(&pg)
	if b != nil {
		ErrorArgs(c)
		return
	}
	has, err := database.GetModelWhere(new(database.WebPage), "project_id = ? AND path = ?", pg.ProjectID, pg.Path)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if has {
		c.JSON(http.StatusOK, ResponseFail(pg.Path+"路径已存在", http.StatusBadRequest))
		return
	}
	_, err = database.CreateModel(&pg)
	msg := "提交成功"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, ResponseOk(msg))
}

func DeleteWebPage(c *gin.Context) {
	data := PostData{}
	b := c.Bind(&data)
	if b != nil {
		ErrorArgs(c)
		return
	}
	m := new(database.WebPage)
	m.ID = data.GetID()

	// https://xorm.io/zh/docs/chapter-10/readme/
	// ADD SESSION 添加事务
	_, err := database.DeleteModel(m)
	if err != nil {
		ErrorServer(c, err)
		return
	}

	mm := new(database.DataTable)
	mm.PageID = m.ID
	result, err := database.DeleteModel(mm)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	// ADD SESSION 添加事务

	c.JSON(http.StatusOK, ResponseOk(fmt.Sprintf("删除成功%d", result)))
}

func GetWebPagesList(c *gin.Context) {
	// var items []database.WebPage
	items := make([]database.WebPage, 0)
	limitStr := c.DefaultQuery("limit", "30")
	pageStr := c.DefaultQuery("page", "1")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	log.Printf("----GetWebPages--limit(%d)---page(%d)----", limit, page)
	err := database.GetAll(&items, limit, page, "project_id = ?", 0)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("请求错误"+err.Error(), 404))
		return
	}
	itemsStr, err := ItemsIDtoString(items)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.String(200, ResponseItems(itemsStr).(string))
}

func GetWebPage(c *gin.Context) {
	idStr := c.DefaultQuery("id", "0")
	if idStr == "0" {
		idStr = c.DefaultQuery("ID", "0")
	}
	pageID, _ := strconv.ParseInt(idStr, 10, 64)
	wpage := database.WebPage{}
	wpage.ID = pageID
	has, err := database.GetModel(&wpage)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	resp := wpage.ToMap(&wpage)
	c.JSON(http.StatusOK, Response(resp, "success", 200))
}
