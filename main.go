package main

import (
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/generator"
	"lemocoder/util"
	"lemocoder/webserver"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var err error
	args := os.Args
	daemon := false
	for k, v := range args {
		// log.Printf("\n----Arg--Key:%d---Val:%s----\n", k, v)
		if v == "-d" {
			daemon = true
			args[k] = ""
		}
		if v == "stop" {
			err = killWebServer()
			if err != nil {
				log.Println("killWebServer Error:", err)
			}
			return
		}
		if v == "init" {
			database.CreateTables()
			err = generator.BuildWebClient()
			if err != nil {
				log.Println("build Error:", err)
			}
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
			err = startBrowser()
			if err != nil {
				log.Println("startBrowser Error:", err)
			}
		}()
	}
	err = webserver.New().ListenAndServe()
	if err != nil {
		log.Println("WebServer Run Error:", err)
	}
}

func init() {
	// time.LoadLocation("Asia/Shanghai")
	config.LoadEnv()
}
