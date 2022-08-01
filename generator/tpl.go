package generator

import (
	"io"
	"os"
	"text/template"
)

// GetContentByTpl
// f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
// GetContentByTpl(tplFilepath, f, data)
// f.Close()
func GetContentByTpl(tplFilepath string, wr io.Writer, data interface{}) error {
	t, err := template.ParseFiles(tplFilepath)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

// func CreateFile(targetPath string, content []byte) (n int, err error) {
// 	var file *os.File
// 	file, err = os.Create(targetPath)
// 	if err != nil {
// 		return 0, err
// 	}
// 	n, err = file.Write(content)
// 	if err != nil {
// 		file.Close()
// 		return n, err
// 	}
// 	err = file.Close()
// 	return n, err
// }

func CreateFile(targetFilepath, tplFilepath string, data interface{}) error {
	f, err := os.OpenFile(targetFilepath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	err = GetContentByTpl(tplFilepath, f, data)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
