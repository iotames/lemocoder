package generator

import (
	"bytes"
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

func runCmd(name string, arg ...string) (string, error) {
	var bf bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &bf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	os.Stdout.Write(bf.Bytes())
	return bf.String(), err
}

func BuildWebClient() error {
	var err error
	var result string
	result, err = runCmd("yarn", "--version")
	if err != nil {
		return err
	}
	fmt.Println("yarn version:", result)

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
