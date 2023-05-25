package main

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/generator"
	"lemocoder/util"
	"lemocoder/webserver"
	"log"
	"os"
	"os/exec"
	"runtime"
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
			stopApp()
			return
		}
		if v == "init" {
			if util.IsPathExists("app.lock") {
				// 只有调用 、/api/client/init 接口才会生成 app.lock 文件
				log.Println("Warning!! 不可重复初始化。请删除app.lock文件后继续")
				return
			}
			database.CreateTables()
			err = generator.BuildWebClient("resource/client")
			if err != nil {
				log.Println("build Error:", err)
			}
			return
		}
		if v == "dbsync" {
			database.SyncTables()
			return
		}
		if v == "dbinit" {
			database.CreateTables()
			return
		}
		if v == "clientinit" {
			err = generator.BuildWebClient("resource/client")
			if err != nil {
				panic(err)
			}
			return
		}
	}
	if daemon {
		var newArgs []string
		if len(args) > 1 {
			newArgs = args[1:]
		}
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			newArgs = append([]string{"/c"}, args...)
			cmd = exec.Command("cmd", newArgs...)
		} else {
			cmd = exec.Command(args[0], newArgs...)
		}
		cmd.Env = os.Environ()
		err = cmd.Start()
		if err != nil {
			fmt.Println("app run fail:", err, args)
		}
		return
	}

	if util.IsPathExists(config.ClientFilepath) {
		go func() {
			app := config.GetApp()
			if app.Env == config.ENV_DEV {
				return
			}
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
