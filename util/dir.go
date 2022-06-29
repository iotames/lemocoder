package util

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func ReadDir(path string, callback func(fileinfo fs.FileInfo)) error {
	filelist, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, fileinfo := range filelist {
		if fileinfo.Mode().IsRegular() {
			callback(fileinfo)
		}
	}
	return nil
}

func Mkdir(path string) error {
	if IsPathExists(path) {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
