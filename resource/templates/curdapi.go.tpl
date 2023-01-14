package handler

import (
	"fmt"
	"lemocoder/database"
	"lemocoder/util"
	"log"
	"net/http"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
)
// TODO 数值型字段在POST时，会变为字符串
<%{ if ne .Create "" }%>
func <%{.Create}%>(c *gin.Context) {
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
func <%{.Delete}%>(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	var result int64
	m := new(database.<%{$.ItemDataTypeName}%>)
	codes, ok := data.GetCodeList()
	if ok {
		if len(codes) == 0 {
			ErrorArgs(c, fmt.Errorf("删除对象ID列表为空"))
			return
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
func <%{.Update}%>(c *gin.Context) {
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

<%{ if ne .GetOne "" }%>
func <%{.GetOne}%>(c *gin.Context) {
	idstr := c.DefaultQuery("id", "0")
	if idstr == "0"{
		ErrorArgs(c, fmt.Errorf("id参数错误."))
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64) // strconv.Atoi(idstr)
	if err != nil{
		ErrorArgs(c, fmt.Errorf("id参数解析错误:%w", err))
		return
	}

	modelFind := database.<%{.ItemDataTypeName}%>{}
	modelFind.ID = id
	has, err := database.GetModel(&modelFind)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if !has {
		ErrorNotFound(c)
		return
	}
	c.JSON(http.StatusOK, Response(modelFind, "success", 200))
}<%{end}%>

<%{ if ne .GetList "" }%>
func <%{.GetList}%>(c *gin.Context) {
	ignoreFields := []string{"current", "pageSize", "page", "limit", "sort"}
	var err error
	var items []database.<%{.ItemDataTypeName}%>

	limitStr := c.DefaultQuery("limit", "30")
	pageStr := c.DefaultQuery("page", "1")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	log.Printf("----GetListItem--limit(%d)---page(%d)----", limit, page)

	searchMap := make(map[string]string)
	querys := c.Request.URL.Query()
	for k, v := range querys {
		if util.GetIndexOf(k, ignoreFields) == -1 && len(v) == 1 {
			searchMap[k] = v[0]
		}
	}
	if len(searchMap) == 0 {
		err = database.GetAll(&items, limit, page, "id > ?", 0)
	} else {
		qlike, qargs := database.GetWhereLikeArgs(searchMap)
		fmt.Printf("\n-----query(%+v)----args(%+v)----\n", qlike, qargs)
		err = database.GetAll(&items, limit, page, qlike, qargs...)
	}

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

<%{range .FuncsItemOpt }%>
func <%{ . -}%>(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	var result int64
	m := new(database.<%{$.ItemDataTypeName}%>)
	postID := data.GetID()
	if postID == 0 {
		ErrorArgs(c, fmt.Errorf("删除对象的ID不能为0"))
		return
	}
	m.ID = data.GetID()
	fmt.Println("TODO: 请填充服务端代码, 操作数据: ", m)
	// TODO 操作单条数据 result, err = database.DeleteModel(m)
	// TODO 操作单条数据 result, err = database.UpdateModel(m, data.GetUpdateData())

	// msg := fmt.Sprintf("%d条记录操作成功", result)
	msg := fmt.Sprintf("TODO: 请填充【服务端代码】.(已操作%d条记录)", result)
	c.JSON(http.StatusOK, ResponseOk(msg))
}
<%{end}%>

<%{range .FuncsItemsBatchOpt }%>
func <%{ . -}%>(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	var result int64
	codes, ok := data.GetCodeList()
	if ok {
		if len(codes) == 0 {
			ErrorArgs(c, fmt.Errorf("删除对象ID列表为空"))
			return
		}
		fmt.Println("TODO: 请填充服务端代码, 操作数据. ID列表: ", codes)
		// TODO 批量操作 result, err = database.BatchDelete(m, codes)
		// TODO 批量操作 result, err = database.BatchUpdate(m, codes)
		// TODO 批量操作 result, err = database.Exec(fmt.Sprintf("UPDATE <%{$.ItemDataTypeName}%> set state = 1 where id IN (%s)", strings.Join(codes, ",")))
	} else {
		ErrorArgs(c, fmt.Errorf("删除对象ID列表不存在"))
		return
	}
	// msg := fmt.Sprintf("%d条记录操作成功", result)
	msg := fmt.Sprintf("TODO: 请填充【服务端代码】.(已操作%d条记录)", result)
	c.JSON(http.StatusOK, ResponseOk(msg))
}
<%{end}%>

<%{range .FuncsFormSubmit }%>
func <%{ . -}%>(c *gin.Context) {
	data := PostData{}
	err := CheckBindArgs(&data, c)
	if err != nil {
		return
	}
	var result int64
	m := new(database.<%{$.ItemDataTypeName}%>)
	postID := data.GetID()
	if postID == 0 {
		ErrorArgs(c, fmt.Errorf("删除对象的ID不能为0"))
		return
	}
	m.ID = data.GetID()
	fmt.Println("TODO: 请填充服务端代码, 操作数据: ", m)
	// TODO 操作单条数据 result, err = database.UpdateModel(&updateModel)

	// msg := fmt.Sprintf("%d条记录操作成功", result)
	msg := fmt.Sprintf("TODO: 请填充【服务端代码】.(已操作%d条记录)", result)
	c.JSON(http.StatusOK, ResponseOk(msg))
}
<%{end}%>