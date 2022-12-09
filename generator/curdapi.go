package generator

import (
	"bytes"
	"fmt"
	"lemocoder/config"
	"lemocoder/model"
	"lemocoder/util"
	"os"
	"strings"
)

const API_ROUTE_FUNC_PREFIX = "handler."

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
	addCode := fmt.Sprintf(`new(%s),
		`, structName)
	return AddCodeToFile(config.ServerTablesPath, addCode)
}

// AddApiRoutes 去除API路由的 `/api` 前缀
func AddApiRoutes(apiRoutes []model.ApiRoute) (routes []model.ApiRoute, err error) {
	for _, route := range apiRoutes {
		if strings.TrimSpace(route.Path) == "" {
			continue
		}
		if route.Method == "GET" && route.Path == "/table/demodata" {
			continue
		}
		if route.Method == "POST" && route.Path == "/demo/post" {
			continue
		}
		if strings.Index(route.Path, "/api/") == 0 {
			route.Path = strings.Replace(route.Path, "/api/", "/", 1)
		}
		if strings.Index(route.FuncName, API_ROUTE_FUNC_PREFIX) != 0 {
			route.FuncName = API_ROUTE_FUNC_PREFIX + route.FuncName
		}
		routes = append(routes, route)
	}
	data := map[string]interface{}{
		"Routes": routes,
	}
	var bf bytes.Buffer
	tplText := `<%{range .Routes}%>
	g.<%{.Method}%>("<%{.Path}%>", <%{.FuncName}%>)
	<%{end}%>`
	err = SetContentByTplText(tplText, data, &bf)
	if err != nil {
		return
	}
	err = AddCodeToFile(config.ServerApiRoutesPath, bf.String())
	return
}

// CreateCurdCode 创建CURD代码
func CreateCurdCode(routes []model.ApiRoute, schema model.TableSchema) error {
	for _, route := range routes {
		funcName := strings.Replace(route.FuncName, API_ROUTE_FUNC_PREFIX, "", 1)
		fmt.Println(funcName)
	}

	// {Method: "GET", Path: t.ItemsDataUrl, FuncName: "GetList" + t.ItemDataTypeName},
	// {Method: "POST", Path: t.ItemCreateUrl, FuncName: "Create" + t.ItemDataTypeName},
	// {Method: "POST", Path: t.ItemUpdateUrl, FuncName: "Update" + t.ItemDataTypeName},
	// {Method: "POST", Path: t.ItemDeleteUrl, FuncName: "Delete" + t.ItemDataTypeName},

	return nil
}
