package handler

import (
	"fmt"
	"lemocoder/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// TODO 数值型字段在POST时，会变为字符串
<%{ if ne .Create "" }%>
func Create<%{$.ItemDataTypeName}%>(c *gin.Context) {
	item := database.<%{$.ItemDataTypeName}%>{}
	err := CheckBindArgs(&item, c)
	if err != nil {
		return
	}
	_, err = database.CreateModel(&item)
	msg := "数据创建成功"
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, ResponseOk(msg))
}<%{end}%>

<%{ if ne .Delete "" }%>
func Delete<%{$.ItemDataTypeName}%>(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	var result int64
	items, ok := data["items"]
	m := new(database.<%{$.ItemDataTypeName}%>)
	if ok {
		var codes []string
		for _, v := range items.([]interface{}) {
			code := v.(map[string]interface{})["ID"].(string)
			codes = append(codes, code)
		}
		result, err = database.BatchDelete(m, codes)
	} else {
		postID := data.GetID()
		if postID == 0 {
			ErrorArgs(c, fmt.Errorf("删除对象的ID不能为0"))
			return
		}
		m.ID = data.GetID()
		result, err = database.DeleteModel(m)
	}
	if err != nil {
		ErrorServer(c, err)
		return
	}

	c.JSON(http.StatusOK, ResponseOk(fmt.Sprintf("删除成功%d", result)))
}<%{end}%>

<%{ if ne .Update "" }%>
func Update<%{.ItemDataTypeName}%>(c *gin.Context) {
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
	modelFind := database.<%{.ItemDataTypeName}%>{}
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
	updateModel := database.<%{.ItemDataTypeName}%>{}
	postData.ParseTo(&updateModel)
	_, err = database.UpdateModel(&updateModel, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("数据更新成功"))
}<%{end}%>

<%{ if ne .GetList "" }%>
func GetList<%{.ItemDataTypeName}%>(c *gin.Context) {
	items := make([]database.<%{.ItemDataTypeName}%>, 0)
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
}<%{end}%>