package generator

import (
	"bytes"
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/initial"
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

// AddClientRoutes 添加客户端路由
func AddClientRoutes(addRoutes []initial.ClientRoute) error {
	tplText := `<%{range . }%>{ layout: <%{.Layout}%>, <%{if ne .Path "" }%> path: '<%{.Path}%>', <%{end}%> name: '<%{.Name}%>', <%{if ne .Component "" }%> component: '<%{.Component}%>', <%{end}%> <%{if ne .Redirect "" }%> redirect: '<%{.Redirect}%>', <%{end}%> },
    <%{end}%>`
	var bf bytes.Buffer
	err := SetContentByTplText(tplText, addRoutes, &bf)
	if err != nil {
		return err
	}
	return AddCodeToFile(config.ClientRoutesPath, bf.String())
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
	tplText := `<%{range .Routes }%>g.<%{.Method}%>("<%{.Path}%>", <%{.FuncName}%>)
	<%{end}%>`
	err = SetContentByTplText(tplText, data, &bf)
	if err != nil {
		return
	}
	err = AddCodeToFile(config.ServerApiRoutesPath, bf.String())
	return
}

type ApiTplData struct {
	ItemDataTypeName, GetList, GetOne, Create, Update, Delete string
	FuncsItemOpt, FuncsItemsBatchOpt, FuncsFormSubmit         []string
}

// CreateCurdCode 创建CURD代码
func CreateCurdCode(routes []model.ApiRoute, schema model.TableSchema) error {
	tplData := ApiTplData{ItemDataTypeName: schema.ItemDataTypeName}
	for _, route := range routes {
		funcName := strings.Replace(route.FuncName, API_ROUTE_FUNC_PREFIX, "", 1)
		if strings.Index(funcName, "GetList") == 0 {
			tplData.GetList = funcName
		}
		if strings.Index(funcName, "GetOne") == 0 {
			tplData.GetOne = funcName
		}
		if strings.Index(funcName, "Create") == 0 {
			tplData.Create = funcName
		}
		if strings.Index(funcName, "Update") == 0 {
			tplData.Update = funcName
		}
		if strings.Index(funcName, "Delete") == 0 {
			tplData.Delete = funcName
		}
		if strings.Index(funcName, "OptItem") == 0 {
			tplData.FuncsItemOpt = append(tplData.FuncsItemOpt, funcName)
		}
		if strings.Index(funcName, "BatchOptItem") == 0 {
			tplData.FuncsItemsBatchOpt = append(tplData.FuncsItemsBatchOpt, funcName)
		}
		if strings.Index(funcName, "FormSubmit") == 0 {
			tplData.FuncsFormSubmit = append(tplData.FuncsFormSubmit, funcName)
		}
		fmt.Println(funcName)
	}
	targetFile := fmt.Sprintf("%s/%s.go", config.ServerHandlerDir, database.ObjToTableCol(schema.ItemDataTypeName))
	tplFile := config.TplDirPath + "/curdapi.go.tpl"
	return CreateFile(targetFile, tplFile, tplData)
}
