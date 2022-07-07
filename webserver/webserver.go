package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type WebServer interface {
	ListenAndServe() error
}

// func debug(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "---Success--Debug----r:%+v--\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_TYPE\"):(%v)---", *r, os.Getenv("DB_HOST"), os.Getenv("DB_TYPE"))
// }

func gdebug(c *gin.Context) {
	c.String(200, fmt.Sprintf("---Success--Debug-----\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_TYPE\"):(%v)---", os.Getenv("DB_HOST"), os.Getenv("DB_TYPE")))
}

func New() WebServer {
	log.Println("====Begin---webserver--Run")
	h := gin.Default()
	h.GET("/gdebug", gdebug)
	server := http.Server{Addr: ":8085", Handler: h}
	// http.HandleFunc("/debug", debug)
	return &server
}
