package main

import (
	"lemocoder/database"
	"lemocoder/util"
	"lemocoder/webserver"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	daemon := false
	for k, v := range args {
		// log.Printf("\n----Arg--Key:%d---Val:%s----\n", k, v)
		if v == "-d" {
			daemon = true
			args[k] = ""
		}
		if v == "stop" {
			err := killWebServer()
			if err != nil {
				log.Println("killWebServer Error:", err)
			}
			return
		}
		if v == "init" {
			database.CreateTables()
			return
		}
		if v == "dbsync" {
			database.SyncTables()
			return
		}
	}

	if daemon {
		var newArgs []string
		if len(args) > 1 {
			newArgs = args[1:]
		}
		cmd := exec.Command(args[0], newArgs...)
		cmd.Env = os.Environ()
		cmd.Start()
		return
	}
	go func() {
		time.Sleep(1 * time.Second)
		err := startBrowser()
		if err != nil {
			log.Println("startBrowser Error:", err)
		}
	}()

	webserver.New().ListenAndServe()
}

func init() {
	if !util.IsPathExists(".env") {
		f, err := os.Create(".env")
		if err != nil {
			panic("Create .env Error: " + err.Error())
		}
		f.Close()
	}
	err := godotenv.Load(".env", "env.default")
	if err != nil {
		panic("godotenv Error: " + err.Error())
	}
}
