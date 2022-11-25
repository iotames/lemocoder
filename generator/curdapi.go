package generator

import (
	"fmt"

	"lemocoder/model"
	"lemocoder/util"
	"os"
	"strings"
)

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
	return nil
}
