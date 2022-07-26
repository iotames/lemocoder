package main

import (
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/util"
	"lemocoder/webserver"
	"log"
	"os"
	"os/exec"
	"time"
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

	if util.IsPathExists(config.ClientFilepath) {
		go func() {
			time.Sleep(1 * time.Second)
			err := startBrowser()
			if err != nil {
				log.Println("startBrowser Error:", err)
			}
		}()
	}

	log.Println("WebServer Run Error:", webserver.New().ListenAndServe().Error())
}

func init() {
	// time.LoadLocation("Asia/Shanghai")
	config.LoadEnv()
}
