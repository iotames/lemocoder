package handler

import (
	"errors"
	"fmt"
	"lemocoder/database"
	gen "lemocoder/generator"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RebuildProject(c *gin.Context) {
	var err error
	postData := PostData{}
	err = postData.ParseBody(c.Request.Body)
	if err != nil {
		ErrorServer(c, fmt.Errorf("request body parse error:%w", err))
		return
	}
	pageIDstr, ok := postData["PageID"]
	if !ok {
		ErrorArgs(c, errors.New("缺少PageID参数"))
		return
	}
	if pageIDstr.(string) == "" {
		ErrorArgs(c, errors.New("参数PageID不能为空"))
		return
	}
	pageID, _ := strconv.ParseInt(pageIDstr.(string), 10, 64)

	// state 状态不对不允许重新编译
	wpage := database.WebPage{}
	wpage.ID = pageID
	err = mustFind(c, &wpage)
	if err != nil {
		return
	}
	if wpage.State < database.PAGE_STATE_CREATED {
		c.JSON(http.StatusOK, ResponseFail("代码未生成，请勿编译", 400))
		return
	}
	if wpage.State == database.PAGE_STATE_BUILT {
		c.JSON(http.StatusOK, ResponseFail("请勿重复操作", 400))
		return
	}

	// 编译客户端
	err = gen.BuildWebClient("resource/client")
	if err != nil {
		ErrorServer(c, err)
		return
	}

	// 编译主程序
	cmd := exec.Command("go", "build", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		ErrorServer(c, err)
		return
	}

	// 更新数据表结构 database.SyncTables()
	cmd = exec.Command("go", "run", ".", "dbsync")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		ErrorServer(c, err)
		return
	}

	// 更新数据库
	database.UpdateModel(&wpage, map[string]interface{}{"state": database.PAGE_STATE_BUILT})
	c.JSON(http.StatusOK, ResponseOk("编译完成, 请重启应用"))
}
