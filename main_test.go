package main

import (
	"bytes"
	"lemocoder/database"
	"lemocoder/generator"
	"lemocoder/model"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestSnakName(t *testing.T) {
	log.Println(database.ObjToTableCol("PageID"))
	log.Println(database.ObjToTableCol("PageId"))
	log.Println("-------------------------")
	log.Println(database.TableColToObj("project_id"))
	log.Println(database.TableColToObj("page_id"))

}

func TestTemplate(t *testing.T) {
	apiRoutes := []model.ApiRoute{
		{Method: "GET", Path: "/api/book/get", FuncName: "GetBook"},
		{Method: "POST", Path: "/api/book/add", FuncName: "AddBook"},
	}
	_, err := generator.AddApiRoutes(apiRoutes)
	if err != nil {
		panic(err)
	}

	// // faldskfal
	// type GameStatus struct {
	// 	Name  string
	// 	IsWin bool
	// }
	// // {{define "debug"}}
	// tpl := template.Must(template.ParseGlob("./resource/templates/debug.common.tpl"))
	// // t1, err := template.ParseFiles("./resource/templates/debug.tpl")
	// var userStatus = []GameStatus{
	// 	{"大春", true},
	// 	{"NiuBee", false},
	// 	{"球球", true},
	// }
	// for _, u := range userStatus {
	// 	err := tpl.ExecuteTemplate(os.Stdout, "debug", u)
	// 	if err != nil {
	// 		log.Println("executing template:", err)
	// 	}
	// }
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

func TestDbQuery(t *testing.T) {
	// table := database.DataTable{PageID: 1578323968463474688}
	// ts := database.DataTable{}
	result, err := database.Query("SELECT * FROM data_tables where page_id = 1578318193674424320")
	if err != nil {
		log.Println("-----Error----", err)
		return
	}
	log.Printf("----len(%+v)------", len(result))
	for _, l := range result {
		for k, v := range l {
			log.Printf("---%s--(%s)---", k, string(v))
		}
	}

}
