package main

import (
	"log"
	"os"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	// faldskfal
	type GameStatus struct {
		Name  string
		IsWin bool
	}
	// {{define "debug"}}
	tpl := template.Must(template.ParseGlob("./resource/templates/debug.common.tpl"))
	// t1, err := template.ParseFiles("./resource/templates/debug.tpl")
	var userStatus = []GameStatus{
		{"大春", true},
		{"NiuBee", false},
		{"球球", true},
	}
	for _, u := range userStatus {
		// err = t1.Execute(w, userStatus)
		err := tpl.ExecuteTemplate(os.Stdout, "debug", u)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

func TestEnv(t *testing.T) {
	title := os.Getenv("APP_TITLE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	webServerPort := os.Getenv("WEB_SERVER_PORT")

	log.Printf("--title:%s--dbHost:%s----dbPort:%s--webServerPort:%s--", title, dbHost, dbPort, webServerPort)
}
