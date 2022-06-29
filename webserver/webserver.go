package webserver

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func tpl(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
	fmt.Println(r)

	t1, err := template.ParseFiles("./resource/templates/debug.tpl")
	if err != nil {
		panic(err)
	}
	type GameStatus struct {
		Name  string
		IsWin bool
	}
	var userStatus = GameStatus{"大春", true}
	// err = t1.Execute(w, userStatus)
	err = t1.Execute(w, userStatus)
	if err != nil {
		panic(err)
	}
}

// func debugg(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello world")
// }

func Run() {
	log.Println("====Begin---webserver--Run")
	server := http.Server{Addr: "127.0.0.1:8085"}
	http.HandleFunc("/debug", tpl)
	server.ListenAndServe()
}
