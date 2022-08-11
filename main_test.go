package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
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
