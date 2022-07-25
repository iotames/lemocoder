package webserver

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	ListenAndServe() error
}
type BaseWebServer struct{}

func (w BaseWebServer) ListenAndServe() error {
	return nil
}

func New() WebServer {
	log.Println("====Begin---webserver--Run")
	conf := config.GetWebServer()
	pid := util.GetPidByPort(conf.Port)
	if pid > 0 {
		return BaseWebServer{}
	}
	h := gin.Default()
	setRouters(h)

	log.Println("----WebServerPort:", conf.Port)
	server := http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: h}
	// http.HandleFunc("/debug", debug)
	return &server
}

// func debug(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "---Success--Debug----r:%+v--\n\n---os.Getenv(\"DB_HOST\"):(%v)---", *r, os.Getenv("DB_HOST"))
// }
