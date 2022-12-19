package handler

import (
	"fmt"
	"lemocoder/database"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWebPage(c *gin.Context) {
	pg := database.WebPage{}
	b := c.Bind(&pg)
	if b != nil {
		ErrorArgs(c)
		return
	}
	has, err := database.GetModelWhere(new(database.WebPage), "project_id = ? AND (path = ? OR component = ?)", pg.ProjectID, pg.Path, pg.Component)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	if has {
		c.JSON(http.StatusOK, ResponseFail(fmt.Sprintf("路径(%s)或组件(%s)已存在", pg.Path, pg.Component), http.StatusBadRequest))
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
	err := mustFind(c, &wpage)
	if err != nil {
		return
	}

	resp := wpage.ToMap(&wpage)
	c.JSON(http.StatusOK, Response(resp, "success", 200))
}

func UpdateWebPage(c *gin.Context) {
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
	err = mustFind(c, &modelFind)
	if err != nil {
		return
	}
	updateModel := database.WebPage{}
	postData.ParseTo(&updateModel)
	if checkWebPage(c, modelFind, updateModel) != nil {
		return
	}
	_, err = database.UpdateModel(&updateModel, nil)
	if err != nil {
		ErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, ResponseOk("数据更新成功"))
}

func checkWebPage(c *gin.Context, origin, update database.WebPage) error {
	var has bool
	var err error
	path := update.Path
	pid := update.ProjectID
	component := update.Component
	m := new(database.WebPage)
	if path != origin.Path {
		has, err = database.GetModelWhere(m, "project_id = ? AND path = ?", pid, path)
		if has {
			msg := fmt.Sprintf("路径(%s)已存在", path)
			c.JSON(http.StatusOK, ResponseFail(msg, http.StatusBadRequest))
			return fmt.Errorf(msg)
		}
	}
	if component != origin.Component {
		has, err = database.GetModelWhere(m, "project_id = ? AND component = ?", pid, component)
		if has {
			msg := fmt.Sprintf("组件(%s)已存在", component)
			c.JSON(http.StatusOK, ResponseFail(msg, http.StatusBadRequest))
			return fmt.Errorf(msg)
		}
	}
	if err != nil {
		ErrorServer(c, err)
		return err
	}
	return nil
}
