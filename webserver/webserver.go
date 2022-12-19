package webserver

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	ListenAndServe() error
}
type BaseWebServer struct {
	Pid, Port int
}

func (w BaseWebServer) ListenAndServe() error {
	if w.Pid > 0 {
		return fmt.Errorf("端口%d已被进程%d占用", w.Port, w.Pid)
	}
	return nil
}

func New() WebServer {
	conf := config.GetWebServer()
	pid := util.GetPidByPort(conf.Port)
	if pid > 0 {
		return BaseWebServer{Pid: pid, Port: conf.Port}
	}
	h := gin.Default()
	setRouters(h)

	addr := fmt.Sprintf(":%d", conf.Port)
	server := http.Server{
		Addr:           addr,
		Handler:        h,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf(`
	欢迎使用 LemoCoder
	当前版本:V1.0.1
	服务端接口运行地址:http://127.0.0.1%s
	客户端资源运行地址:http://127.0.0.1%s/client
`, addr, addr)

	if !util.IsPathExists(config.ClientFilepath) {
		fmt.Println("警告! 缺少客户端资源文件: " + config.ClientFilepath)
	}

	return &server
}
