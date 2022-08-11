package util

import (
	"fmt"
	"io"
	"os"
	"path"
)

// 判断文件或文件夹是否存在
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
	bytes, err := os.ReadFile(path)
	// bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

// https://blog.csdn.net/whatday/article/details/109709845

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// CopyDir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	// var fds []os.FileInfo
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	// if fds, err = ioutil.ReadDir(src); err != nil {
	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
