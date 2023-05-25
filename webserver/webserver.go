package webserver

import (
	"context"
	"fmt"
	"lemocoder/config"
	"lemocoder/util"
	"lemocoder/webserver/handler"
	"net/http"
	"time"

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
	go func() {
		<-handler.StopChan
		time.Sleep(5 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			panic(err)
		}
	}()
	app := config.GetApp()

	fmt.Printf(`
	欢迎使用 LemoCoder
	当前版本: %s
	服务端接口运行地址: http://127.0.0.1%s
	客户端资源运行地址: http://127.0.0.1%s/client
`, app.Version, addr, addr)

	if !util.IsPathExists(config.ClientFilepath) {
		fmt.Println("警告! 缺少客户端资源文件: " + config.ClientFilepath)
	}

	return &server
}
