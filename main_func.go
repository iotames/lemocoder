package main

import (
	"fmt"
	"lemocoder/util"
	"os"
	"strconv"
)

func killWebServer() error {
	portStr := os.Getenv("WEB_SERVER_PORT")
	port, _ := strconv.Atoi(portStr)
	pid := util.GetPidByPort(port)
	if pid > -1 {
		return util.KillPid(fmt.Sprintf("%d", pid))
	}
	return fmt.Errorf("PID is %d. No Process Running", pid)
}

func startBrowser() error {
	return util.StartBrowserByUrl("http://127.0.0.1:" + os.Getenv("WEB_SERVER_PORT") + "/client")
}
