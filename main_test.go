package main

import (
	"bytes"
	"fmt"
	"io"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/generator"
	"lemocoder/model"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// AddApiRoutes 去除API路由的 `/api` 前缀
func TestAddApiRoutes(t *testing.T) {
	var routes []model.ApiRoute
	apiRoutes := []model.ApiRoute{
		{Method: "GET", Path: "/api/products/list", FuncName: "GetListProduct"},
		{Method: "POST", Path: "/api/product/create", FuncName: "CreateProduct"},
		{Method: "POST", Path: "/api/product/update", FuncName: "UpdateProduct"},
		{Method: "POST", Path: "/api/product/delete", FuncName: "DeleteProduct"},
	}
	for _, route := range apiRoutes {
		if strings.TrimSpace(route.Path) == "" {
			continue
		}
		if route.Method == "GET" && route.Path == "/table/demodata" {
			continue
		}
		if route.Method == "POST" && route.Path == "/demo/post" {
			continue
		}
		if strings.Index(route.Path, "/api/") == 0 {
			route.Path = strings.Replace(route.Path, "/api/", "/", 1)
		}
		if strings.Index(route.FuncName, generator.API_ROUTE_FUNC_PREFIX) != 0 {
			route.FuncName = generator.API_ROUTE_FUNC_PREFIX + route.FuncName
		}
		routes = append(routes, route)
	}
	data := map[string]interface{}{
		"Routes": routes,
	}
	var bf bytes.Buffer
	tplText := `<%{range .Routes }%>g.<%{.Method}%>("<%{.Path}%>", <%{.FuncName}%>)
	<%{end}%>`
	err := generator.SetContentByTplText(tplText, data, &bf)
	if err != nil {
		panic(err)
	}
	addCode := bf.String()
	filepath := config.ServerApiRoutesPath

	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	var before []byte
	before, err = io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	f.Close()
	commentEnd := generator.CommentEnd
	log.Println(strings.Replace(string(before), commentEnd, addCode+commentEnd, 1))
}

func TestSnakName(t *testing.T) {
	log.Println(database.ObjToTableCol("PageID"))
	log.Println(database.ObjToTableCol("PageId"))
	log.Println("-------------------------")
	log.Println(database.TableColToObj("project_id"))
	log.Println(database.TableColToObj("page_id"))
}

func TestTemplate(t *testing.T) {
	apiRoutes := []model.ApiRoute{
		{Method: "GET", Path: "/api/book/get", FuncName: "GetBook"},
		{Method: "POST", Path: "/api/book/add", FuncName: "AddBook"},
	}
	_, err := generator.AddApiRoutes(apiRoutes)
	if err != nil {
		panic(err)
	}

	// // faldskfal
	// type GameStatus struct {
	// 	Name  string
	// 	IsWin bool
	// }
	// // {{define "debug"}}
	// tpl := template.Must(template.ParseGlob("./resource/templates/debug.common.tpl"))
	// // t1, err := template.ParseFiles("./resource/templates/debug.tpl")
	// var userStatus = []GameStatus{
	// 	{"大春", true},
	// 	{"NiuBee", false},
	// 	{"球球", true},
	// }
	// for _, u := range userStatus {
	// 	err := tpl.ExecuteTemplate(os.Stdout, "debug", u)
	// 	if err != nil {
	// 		log.Println("executing template:", err)
	// 	}
	// }
}

func TestEnv(t *testing.T) {
	title := os.Getenv("APP_TITLE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	webServerPort := os.Getenv("WEB_SERVER_PORT")

	log.Printf("--title:%s--dbHost:%s----dbPort:%s--webServerPort:%s--", title, dbHost, dbPort, webServerPort)
}

func TestExec(t *testing.T) {
	var outBuffer bytes.Buffer

	cmd := exec.Command("yarn", "--version")
	// cmd := exec.Command("echo", "%myvar%")
	// cmd.Env = append(os.Environ(), "myvar=hellomyvar")
	cmd.Stdout = &outBuffer
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	log.Println("Exec.Stdout:", outBuffer.String())
	log.Println("cmd.Path", cmd.Path)
	// log.Println("os.Env:", os.Environ())
	os.Stderr.Write(outBuffer.Bytes())
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

func TestDbQuery(t *testing.T) {
	// table := database.DataTable{PageID: 1578323968463474688}
	// ts := database.DataTable{}
	result, err := database.Query("SELECT * FROM data_tables where page_id = 1578318193674424320")
	if err != nil {
		log.Println("-----Error----", err)
		return
	}
	log.Printf("----len(%+v)------", len(result))
	for _, l := range result {
		for k, v := range l {
			log.Printf("---%s--(%s)---", k, string(v))
		}
	}

}

func TestStatus(t *testing.T) {
	hostname, _ := os.Hostname()
	cpuNum := runtime.NumCPU() // maxProcs := runtime.GOMAXPROCS(0)
	os := runtime.GOOS         // 操作系统
	arch := runtime.GOARCH     // 架构
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Error(err)
	}
	netInfo := ""
	for i, inter := range interfaces {
		addrs, _ := inter.Addrs()
		addrInfo := ""
		for _, addr := range addrs {
			addrInfo += fmt.Sprintf("[%s -> %s]", addr.Network(), addr.String())
		}
		netInfo += fmt.Sprintf("\n--[%d]--Name(%s)-HardwareAddr(%s)--Addrs(%s)--\n", i, inter.Name, inter.HardwareAddr, addrInfo)
	}

	t.Logf("\n--hostname(%s)--cpuNum(%d)--os(%s)-arch(%s)--%s-\n", hostname, cpuNum, os, arch, netInfo)

	// https://blog.csdn.net/whatday/article/details/109620192
	cpup, _ := cpu.Percent(time.Second, false)
	vm, _ := mem.VirtualMemory()
	parts, _ := disk.Partitions(true)
	t.Log(parts)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	cpuInfo, _ := cpu.Info()
	cpuNumReal, _ := cpu.Counts(false)
	cpuLogicNum, _ := cpu.Counts(true)
	t.Logf("cpu percent:%.2f%%", cpup[0])
	t.Log("VirtualMemory", vm.UsedPercent, vm.Total, vm.Free, vm.Active, vm.Available)
	t.Logf("diskInfo--%.2f%%--Total(%.2f)--Free(%.2f)--", diskInfo.UsedPercent, float64(diskInfo.Total/(1024*1024*1024)), float64(diskInfo.Free/(1024*1024*1024)))
	t.Log("CPU info:", cpuInfo, "real CPU num:", cpuNumReal, "CPU logic num:", cpuLogicNum)
	// [{"cpu":0,"vendorId":"GenuineIntel","family":"205","model":"","stepping":0,"physicalId":"BFEBFBFF000A0653","coreId":"","cores":12,"modelName":"Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz","mhz":2904,"cacheSize":0,"flags":[],"microcode":""}]
	// for{
	// 	info, _ := cpu.Percent(time.Duration(time.Second), false)
	// 	fmt.Println(info)
	// }
}
