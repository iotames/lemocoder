package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

// GetContentByTpl
// f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
// GetContentByTpl(tplFilepath, f, data)
// f.Close()
func GetContentByTpl(tplFilepath string, wr io.Writer, data interface{}) error {
	// t, err := template.ParseFiles(tplFilepath)
	t, err := parseFiles(tplFilepath)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func parseFiles(filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("template: no files named in call to ParseFiles")
	}
	var t *template.Template
	for _, filename := range filenames {
		name, b, err := readFileOS(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
			t.Delims("<%{", "}%>")
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
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
	f, err := os.OpenFile(targetFilepath, os.O_RDWR|os.O_CREATE, 0755)
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

func readFileOS(file string) (name string, b []byte, err error) {
	name = filepath.Base(file)
	b, err = os.ReadFile(file)
	return
}
