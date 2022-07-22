package webserver

import (
	"fmt"
	"lemocoder/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	ListenAndServe() error
}

func New() WebServer {
	log.Println("====Begin---webserver--Run")
	h := gin.Default()
	setRouters(h)
	conf := config.GetWebServer()
	log.Println("----WebServerPort:", conf.Port)
	server := http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: h}
	// http.HandleFunc("/debug", debug)
	return &server
}

// func debug(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "---Success--Debug----r:%+v--\n\n---os.Getenv(\"DB_HOST\"):(%v)---", *r, os.Getenv("DB_HOST"))
// }
