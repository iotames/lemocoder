package util

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// GetPidByPort. 传入查询的端口号
// 返回端口号对应的进程PID，若没有找到相关进程，返回-1
func GetPidByPort(portNumber int) int {
	supportOSs := map[string]bool{
		"windows": true,
		"linux":   true,
		"darwin":  false,
	}
	support, ok := supportOSs[runtime.GOOS]
	if !ok || !support {
		panic("GetPidByPort Not Support " + runtime.GOOS)
	}

	res := -1
	var outBytes bytes.Buffer
	var cmdStr string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// fmt.Printf("---resStr---%s---resStr---", resStr)
		// 以换行符结尾`TCP    127.0.0.1:9222         0.0.0.0:0              LISTENING       9700
		//`
		cmdStr = fmt.Sprintf("netstat -ano -p tcp | findstr %d", portNumber)
		cmd = exec.Command("cmd", "/c", cmdStr)
	}

	if runtime.GOOS == "linux" {
		// processInfo := exec.Command("/bin/sh", "-c",`lsof -i:8299 | awk '{print $2}' | awk  'NR==2{print}'`)
		// 直接返回端口号, 但是以换行符结尾。直接转换字符串为数字会出BUG。使用 strings.TrimSpace 函数转换之。
		cmdStr = fmt.Sprintf(`lsof -i:%d | awk '{print $2}' | awk  'NR==2{print}'`, portNumber)
		cmd = exec.Command("/bin/sh", "-c", cmdStr)
	}

	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	log.Printf("----Executed---For---GetPidByPort---%s----", cmdStr)
	log.Printf("----Result--outString:%s-------", resStr)
	if len(outBytes.Bytes()) == 0 {
		return res
	}

	portStr := ""
	if runtime.GOOS == "linux" {
		portStr = strings.TrimSpace(resStr)
	}
	if runtime.GOOS == "windows" {
		r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
		if len(r) == 0 {
			return res
		}
		portStr = strings.TrimSpace(r[0])
	}
	pid, err := strconv.Atoi(portStr)
	if err == nil {
		res = pid
	}
	return res
}
