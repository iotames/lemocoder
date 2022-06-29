package generator

import (
	"bytes"
	"io"
	"lemocoder/util"
	"text/template"
)

func setContentByTpl(tplFilepath string, buff io.Writer, data interface{}) error {
	t, err := template.ParseFiles(tplFilepath)
	if err != nil {
		return err
	}
	return t.Execute(buff, data)
}

func CreateFileByTpl(tplFilepath, targetPath string, data interface{}) error {
	var bf bytes.Buffer
	err := setContentByTpl(tplFilepath, &bf, data)
	if err != nil {
		return err
	}
	return util.CreateFileByString(bf.String(), targetPath)
}
