package generator

import (
	"fmt"
	"lemocoder/util"
	"os"
	"os/exec"
)

func runCmdInClient(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = "webclient"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func BuildWebClient() error {
	var err error
	var result []byte
	result, err = util.RunCmd("yarn", "--version")
	if err != nil {
		return err
	}
	fmt.Println("yarn version:", string(result))

	err = runCmdInClient("yarn")
	if err != nil {
		return err
	}
	err = runCmdInClient("yarn", "build")
	if err != nil {
		return err
	}

	return util.CopyDir("webclient/dist", "resource/client")
}
