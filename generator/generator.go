package generator

import (
	"bytes"
	"io"
	"os"
)

// const commentStart = "// Code generated Begin; DO NOT EDIT."
const commentEnd = "// Code generated End; DO NOT EDIT."

func AddCodeToFile(filepath string, addCode string) error {
	// 读取原文件内容
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	var before []byte
	before, err = io.ReadAll(f)
	if err != nil {
		return err
	}
	f.Close()
	// 写入变更后的内容
	f, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	f.Write(bytes.Replace(before, []byte(commentEnd), []byte(addCode + commentEnd), 1))
	return f.Close()
}
