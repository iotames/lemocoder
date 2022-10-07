package handler

import (
	"lemocoder/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AddWebPage(c *gin.Context) {
	pg := database.WebPage{}
	b := c.Bind(&pg)
	if b != nil {
		c.JSON(http.StatusOK, ResponseFail("请求参数解析错误", 404))
		return
	}
	has, err := database.GetModelWhere(new(database.WebPage), "project_id = ? AND path = ?", pg.ProjectID, pg.Path)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail(err.Error(), 500))
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

func GetWebPages(c *gin.Context) {
	// var items []database.WebPage
	items := make([]database.WebPage, 0)
	err := database.GetAll(&items, 30, 1, "project_id = ?", 0)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFail("请求错误"+err.Error(), 404))
		return
	}
	c.JSON(200, ResponseItems(items))
}