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

const (
	TYPE_DB_INT      = "INT"
	TYPE_DB_SMALLINT = "SMALLINT"
	TYPE_DB_BIGINT   = "BIGINT"
	TYPE_DB_FLOAT    = "FLOAT"
	TYPE_DB_STRING   = "STRING"
	TYPE_DB_TEXT     = "TEXT"
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

func getDataTypeForJS(t string) string {
	numbers := []string{"int", "float", "smallint"}
	if util.GetIndexOf(t, numbers) > -1 {
		return "number"
	}
	return t
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
		if field.Name == "ID" {
			html += fmt.Sprintf(`<%s name="ID" hidden`, FORM_COMPONENT_TEXT)
		} else {
			html += fmt.Sprintf(`<%s name="%s" label="%s"`, field.Component, field.Name, field.Label)
		}
		if field.Width != "" {
			html += ` width="` + field.Width + `"`
		}
		if field.Component == "ProFormSelect" {
			html += ` request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]}`
		}
		if field.Placeholder != "" {
			html += fmt.Sprintf(` placeholder="%s"`, field.Placeholder)
		}
		html += " />"
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

func dbtype(t string) string {
	result := "VARCHAR(255)"
	switch strings.ToUpper(t) {
	case TYPE_DB_INT:
		result = "INT"
	case TYPE_DB_SMALLINT:
		result = "SMALLINT"
	case TYPE_DB_BIGINT:
		result = "BIGINT"
	case TYPE_DB_FLOAT:
		result = "FLOAT"
	case TYPE_DB_STRING:
		result = "VARCHAR"
	case TYPE_DB_TEXT:
		result = "TEXT"
	}
	return result
}

func dbdefault(t string) string {
	result := ""
	switch strings.ToUpper(t) {
	case TYPE_DB_FLOAT, TYPE_DB_BIGINT, TYPE_DB_INT, TYPE_DB_SMALLINT:
		result = "default(0)"
	}
	return result
}

func gotype(t string) string {
	switch strings.ToUpper(t) {
	case TYPE_DB_FLOAT:
		t = "float64"
	case TYPE_DB_BIGINT:
		t = "int64"
	case TYPE_DB_INT, TYPE_DB_SMALLINT:
		t = "int"
	case TYPE_DB_TEXT, TYPE_DB_STRING:
		t = "string"
	}
	return t
}
