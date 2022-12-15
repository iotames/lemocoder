package generator

import (
	"fmt"
	"lemocoder/database"
	"lemocoder/model"
	"lemocoder/util"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func readFileOS(file string) (name string, b []byte, err error) {
	name = filepath.Base(file)
	b, err = os.ReadFile(file)
	return
}

func getFormFieldHtml(field model.FormFieldSchema) string {
	html := ""
	grouplen := len(field.Group)
	if grouplen > 0 {
		html += "<ProForm.Group>"
		for _, f := range field.Group {
			html += getFormFieldHtml(f)
		}
		html += "</ProForm.Group>"
	}
	if grouplen == 0 {
		html += fmt.Sprintf(`<%s name="%s" label="%s"`, field.Component, field.Name, field.Label)
		if field.Width != "" {
			html += ` width="` + field.Width + `"`
		}
		if field.Component == "ProFormSelect" {
			html += ` request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]}`
		}
		html += fmt.Sprintf(` placeholder="%s" />`, field.Placeholder)
	}
	return html
}

func getFormFieldsHtml(fields []model.FormFieldSchema) string {
	html := ""
	for _, v := range fields {
		html += getFormFieldHtml(v)
	}
	return html
}

func getDataTypeForJS(t string) string {
	numbers := []string{"int", "float"}
	if util.GetIndexOf(t, numbers) > -1 {
		return "number"
	}
	return t
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
			t = newTpl(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name) // .Funcs(tplFuncs)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func dbtype(t string) string {
	result := "VARCHAR"
	switch strings.ToUpper(t) {
	case "INT":
		result = "INT"
	case "SMALLINT":
		result = "SMALLINT"
	case "BIGINT":
		result = "BIGINT"
	case "FLOAT":
		result = "FLOAT"
	case "STRING":
		result = "VARCHAR"
	case "TEXT":
		result = "TEXT"
	}
	return result
}

func dbdefault(t string) string {
	result := ""
	switch t {
	case "float":
		result = "default(0)"
	case "int":
		result = "default(0)"
	}
	return result
}

func gotype(t string) string {
	switch t {
	case "float":
		t = "float64"
	}
	return t
}

func newTpl(name string) *template.Template {
	tplFuncs := template.FuncMap{
		"getDataTypeForJS":  getDataTypeForJS,
		"getFormFieldsHtml": getFormFieldsHtml,
		"toObjStr":          database.TableColToObj,
		"dbtype":            dbtype,
		"dbdefault":         dbdefault,
		"gotype":            gotype,
	}
	return template.New(name).Funcs(tplFuncs).Delims("<%{", "}%>")
}
