package generator

import (
	"io"
	"os"
)

// SetContentByTplFile
// f, err := os.OpenFile(targetFilepath, os.O_CREATE|os.O_WRONLY, 0o755)
// SetContentByTplFile(tplFilepath, f, data)
// f.Close()
func SetContentByTplFile(tplFilepath string, wr io.Writer, data interface{}) error {
	// t, err := template.ParseFiles(tplFilepath)
	t, err := parseFiles(tplFilepath)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func CreateFile(targetFilepath, tplFilepath string, data interface{}) error {
	f, err := os.OpenFile(targetFilepath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	err = SetContentByTplFile(tplFilepath, f, data)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
