package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type WebServer interface {
	ListenAndServe() error
}

func debug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---Success--Debug----r:%+v--\n\n---os.Getenv(\"DB_HOST\"):(%v)----os.Getenv(\"DB_TYPE\"):(%v)---", *r, os.Getenv("DB_HOST"), os.Getenv("DB_TYPE"))
}

func New() WebServer {
	log.Println("====Begin---webserver--Run")
	server := http.Server{Addr: "127.0.0.1:8085"}
	http.HandleFunc("/debug", debug)
	return &server
}
