package util

import (
	"io/ioutil"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFileByString(content string, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

//判断文件或文件夹是否存在
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		// fmt.Println(stat.IsDir())
		// fmt.Println(stat.Name())
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// ReadFileString. return string(btyes)
// 最简: ioutil.ReadFile()直接读取到[]byte
// 最优: os.Open()读取文件到f(记得defer f.Close()), ioutil.ReadAll()读取数据到[]byte
// os.Open()到f; var chunk []byte; buf := make([]byte, 1024); f.Read(buf); chunk = append(chunk, buf[:n]...)
func ReadFileString(path string) string {
	bytes, err := ioutil.ReadFile(path)
	checkError(err)
	return string(bytes)
}
