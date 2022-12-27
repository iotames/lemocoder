package handler

import (
	"errors"
	"fmt"
	"lemocoder/database"
	gen "lemocoder/generator"
	"lemocoder/status"
	"lemocoder/util"
	"net/http"
	"os"
	"os/exec"
	"runtime"
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

type DevTool struct {
	Name, Version, Url string
}

type PlatformStatus struct {
	HostName, OS, Arch, CpuName, MemoryTotalText, MemoryFreeText string
	CpuUsedPercent, MemoryUsedPercent                            float64
	CpuNum, MemoryTotal, MemoryFree                              int
	DevTools                                                     []DevTool
	DiskInfo                                                     []status.DiskPartInfo
}

func GetOsStatus(c *gin.Context) {
	vGo, _ := util.RunCmd("go", "version")
	gov := ""
	if len(vGo) != 0 {
		gov = (util.NewStrfind(string(vGo))).SetRegexp(`(\d+\.\d+)`).DoFind().GetOne(false)
	}

	vNode, _ := util.RunCmd("node", "--version")
	nodev := ""
	if len(vNode) != 0 {
		nodev = (util.NewStrfind(string(vNode))).SetRegexp(`(\d+\.\d+\.\d+)`).DoFind().GetOne(false)
	}

	vYarn, _ := util.RunCmd("yarn", "--version")
	yarnv := string(vYarn)

	vGit, _ := util.RunCmd("git", "--version")
	gitv := ""
	if len(vGit) != 0 {
		gitv = (util.NewStrfind(string(vGit))).SetRegexp(`(\d+\.\d+\.\d+)`).DoFind().GetOne(false)
	}

	devTools := []DevTool{
		{Name: "Go", Version: gov, Url: "https://golang.google.cn/doc/install"},
		{Name: "Node", Version: nodev, Url: "https://nodejs.org/zh-cn/download/"},
		{Name: "Yarn", Version: yarnv, Url: "https://yarn.bootcss.com/docs/install"},
		{Name: "Git", Version: gitv, Url: "https://git-scm.com/downloads"},
	}
	hostname, _ := os.Hostname()
	// memStat := new(runtime.MemStats)
	// runtime.ReadMemStats(memStat)
	cpuInfo := status.GetCpuInfo()
	minfo := status.GetMemoryInfo()
	data := PlatformStatus{
		HostName:          hostname,
		CpuName:           cpuInfo.Name,
		CpuNum:            cpuInfo.Num,
		CpuUsedPercent:    cpuInfo.UsedPercent,
		MemoryUsedPercent: minfo.UsedPercent,
		MemoryTotal:       int(minfo.Total),
		MemoryFree:        int(minfo.Free),
		MemoryTotalText:   minfo.TotalText(),
		MemoryFreeText:    minfo.FreeText(),
		DiskInfo:          status.GetDiskInfo(),
		OS:                runtime.GOOS,   // 操作系统 win
		Arch:              runtime.GOARCH, // 体系架构 amd64
		DevTools:          devTools,
	}
	c.JSON(http.StatusOK, Response(data, "success", http.StatusOK))
}
