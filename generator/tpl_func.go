package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func readFileOS(file string) (name string, b []byte, err error) {
	name = filepath.Base(file)
	b, err = os.ReadFile(file)
	return
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
