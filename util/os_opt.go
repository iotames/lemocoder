package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

func StartBrowserByUrl(url string) error {
	var cmdMap = map[string]*exec.Cmd{
		"windows": exec.Command("cmd", "/c", "start", url),
		"darwin":  exec.Command("/bin/bash", "-c", "open", url),
		"linux":   exec.Command("/bin/bash", "-c", "xdg-open", url),
	}
	eCmd, ok := cmdMap[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	return eCmd.Start()
}

func KillPid(pid string) error {
	fmt.Println("Kill---PID:", pid)
	unixCmd := exec.Command("/bin/bash", "-c", "kill "+pid)
	cmdMap := map[string]*exec.Cmd{
		"windows": exec.Command("cmd", "/c", fmt.Sprintf("taskkill -pid %s -F", pid)), // MUST ADD arg -F
		"linux":   unixCmd,
		"darwin":  unixCmd,
	}
	cmd, ok := cmdMap[runtime.GOOS]
	if !ok {
		return fmt.Errorf("Do Not Support platform: %s", runtime.GOOS)
	}
	return cmd.Run()
}
