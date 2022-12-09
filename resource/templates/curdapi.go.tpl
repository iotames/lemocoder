package handler

import (
	"fmt"
	"lemocoder/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	pg := database.WebPage{}
	err := CheckBindArgs(&pg, c)
	if err != nil {
		return
	}
	_, err = database.CreateModel(&pg)
	msg := "数据创建成功"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, ResponseOk(msg))
}

func DeleteItem(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	postID := data.GetID()
	if postID == 0 {
		ErrorArgs(c, fmt.Errorf("删除对象的ID不能为0"))
		return
	}

	m := new(database.WebPage)
	m.ID = data.GetID()
	result, err := database.DeleteModel(m)
	if err != nil {
		ErrorServer(c, err)
		return
	}

	c.JSON(http.StatusOK, ResponseOk(fmt.Sprintf("删除成功%d", result)))
}

func UpdateItem(c *gin.Context) {
	postData := PostData{}
	err := postData.ParseBody(c.Request.Body)
	if err != nil {
		ErrorServer(c, fmt.Errorf("request body parse error:%w", err))
		return
	}	
	if postData.GetID() == 0 {
		ErrorArgs(c, fmt.Errorf("操作对象的ID不能为0"))
		return
	}
	modelFind := database.WebPage{}
	modelFind.ID = postData.GetID()
	has, err := database.GetModel(&modelFind)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	updateModel := database.WebPage{}
	postData.ParseTo(&updateModel)
	_, err = database.UpdateModel(&updateModel, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("数据更新成功"))
}

func GetListItem(c *gin.Context) {
	items := make([]database.WebPage, 0)
	limitStr := c.DefaultQuery("limit", "30")
	pageStr := c.DefaultQuery("page", "1")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	log.Printf("----GetListItem--limit(%d)---page(%d)----", limit, page)
	err := database.GetAll(&items, limit, page, "id > ?", 0)
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