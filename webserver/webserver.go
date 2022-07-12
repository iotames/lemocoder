package webserver

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	ListenAndServe() error
}

func New() WebServer {
	log.Println("====Begin---webserver--Run")
	h := gin.Default()
	setRouters(h)
	portStr := os.Getenv("WEB_SERVER_PORT")
	server := http.Server{Addr: ":" + portStr, Handler: h}
	// http.HandleFunc("/debug", debug)
	return &server
}

// func debug(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "---Success--Debug----r:%+v--\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_TYPE\"):(%v)---", *r, os.Getenv("DB_HOST"), os.Getenv("DB_TYPE"))
// }
