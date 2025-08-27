package generator

import (
	"fmt"
	"lemocoder/util"
	"os"
	"os/exec"
)

func runCmdInClient(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Env = cmd.Environ()
	// 解决yarn build报错：ERR_OSSL_EVP_UNSUPPORTED。因一些旧模块可能使用了 OpenSSL 3.0 默认不再允许的算法或密钥大小，从而引发此错误。
	cmd.Env = append(cmd.Env, "NODE_OPTIONS=--openssl-legacy-provider")
	// fmt.Printf("----runCmdInClient----env(%+v)---\n", cmd.Env)
	cmd.Dir = "webclient"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func BuildWebClient(target string) error {
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

	return util.CopyDir("webclient/dist", target)
}
