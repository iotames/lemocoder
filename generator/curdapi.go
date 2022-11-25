package generator

import (
	"bytes"
	"fmt"
	"io"
	"lemocoder/config"
	"lemocoder/model"
	"lemocoder/util"
	"os"
	"strings"
)

// const commentStart = "// Code generated Begin; DO NOT EDIT."
const commentEnd = "// Code generated End; DO NOT EDIT."

func AddDbModel(fields []model.TableItemSchema, structName string) error {
	j := 0
	skipFields := []string{"id", "created_at", "updated_at"}
	for _, v := range fields {
		if util.GetIndexOf(v.DataName, skipFields) == -1 {
			fields[j] = v
			j++
		}
	}
	fields = fields[:j]
	tableName := strings.ToLower(structName)
	path := fmt.Sprintf("database/%s.go", tableName)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	tplText := `package database

type <%{.StructName}%> struct {
	BaseModel           "xorm:\"extends\""
	<%{range .Fields}%><%{.DataName|toObjStr}%>    <%{.DataType|gotype}%>  "xorm:\"<%{.DataType|dbtype}%> notnull <%{.DataType|dbdefault}%> '<%{.DataName}%>'\""
	<%{end}%>
}

func (m <%{.StructName}%>) TableName() string {
	return "<%{.TableName}%>"
}
`

	data := map[string]interface{}{
		"StructName": structName,
		"TableName":  tableName,
		"Fields":     fields,
	}

	err = SetContentByTplText(tplText, data, f)
	if err != nil {
		return err
	}
	return AddDbModelToTables(structName)
}

func AddDbModelToTables(structName string) error {
	// 读取原文件内容
	f, err := os.OpenFile(config.ServerTablesPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	var before []byte
	before, err = io.ReadAll(f)
	if err != nil {
		return err
	}
	f.Close()
	replactstr := fmt.Sprintf(`    new(%s),
	`+commentEnd, structName)
	replaceBytes := []byte(replactstr)

	// 写入变更后的内容
	f, err = os.OpenFile(config.ServerApiRoutesPath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	f.Write(bytes.Replace(before, []byte(commentEnd), replaceBytes, 1))
	return f.Close()
}
