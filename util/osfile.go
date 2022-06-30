package util

import (
	"io/ioutil"
	"os"
)

//判断文件或文件夹是否存在
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		// fmt.Println(stat.IsDir())
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// ReadFileString. return string, error
// 最简: ioutil.ReadFile()直接读取到[]byte
// 最优: os.Open()读取文件到f(记得defer f.Close()), ioutil.ReadAll(r io.Reader)读取数据到[]byte
// os.Open()到f; var chunk []byte; buf := make([]byte, 1024); f.Read(buf); chunk = append(chunk, buf[:n]...)
func ReadFileToString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}
