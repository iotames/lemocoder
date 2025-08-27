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
	"strings"
	"sync"

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

func getVersion(b []byte, rege string) string {
	v := string(b)
	if rege != "" && len(b) != 0 {
		vstr := (util.NewStrfind(v)).SetRegexp(rege).DoFind().GetOne(false)
		if vstr != "" {
			v = strings.TrimSpace(vstr)
		}
	}
	return v
}

type CmdInfo struct {
	Name, Url, Version string
	Args               []string
	Result             []byte
	TextRege           string
}

func GetOsStatus(c *gin.Context) {
	var cmdlist []CmdInfo = []CmdInfo{
		{Name: "go", Args: []string{"version"}, TextRege: `(\d+\.\d+)`, Url: "https://golang.google.cn/doc/install"},
		{Name: "node", Args: []string{"--version"}, TextRege: `(\d+\.\d+\.\d+)`, Url: "https://nodejs.org/zh-cn/download/"},
		{Name: "yarn", Args: []string{"--version"}, TextRege: "", Url: "https://yarn.bootcss.com/docs/install"},
		{Name: "git", Args: []string{"--version"}, TextRege: `(\d+\.\d+\.\d+)`, Url: "https://git-scm.com/downloads"},
	}
	wg := sync.WaitGroup{}
	var devTools []DevTool
	for i, cmd := range cmdlist {
		wg.Add(1)
		go func(cmd CmdInfo, ii int) {
			cmd.Result, _ = util.RunCmd(cmd.Name, cmd.Args...)
			cmd.Version = getVersion(cmd.Result, cmd.TextRege)
			fmt.Println("-------", ii, cmd.Name, cmd.Version)
			// cmdlist[ii] = cmd
			devTools = append(devTools, DevTool{Name: cmd.Name, Version: cmd.Version, Url: cmd.Url})
			wg.Done()
		}(cmd, i)
	}
	wg.Wait()
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
