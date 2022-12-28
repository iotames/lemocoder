package generator

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/initial"
	"lemocoder/model"
	"lemocoder/util"
	"strings"
)

// CreateTableClient 创建客户端源码文件
func CreateTableClient(t model.TableSchema, p database.WebPage) error {
	// 创建数据表格页面的tsx文件
	err := CreateFile(fmt.Sprintf("%s/%s.tsx", config.ClientSrcPagesDir, p.Component), config.TplDirPath+"/table.tsx.tpl", t)
	if err != nil {
		return err
	}

	// 添加客户端路由
	rts := []initial.ClientRoute{
		{Name: p.Name, Path: p.Path, Component: "./" + p.Component, Layout: true},
	}
	return AddClientRoutes(rts)
}

// CreateTableServer 创建服务端API源码文件
func CreateTableServer(t model.TableSchema) error {
	dataTypeName := database.TableColToObj(t.ItemDataTypeName)
	// 创建ORM数据表模型文件，并添加到数据列表中
	err := AddDbModel(t.Items, dataTypeName)
	if err != nil {
		return err
	}

	apiRoutes := []model.ApiRoute{
		// {Method: "GET", Path: "", FuncName: "Get"+dataTypeName,
		{Method: "GET", Path: t.ItemsDataUrl, FuncName: "GetList" + dataTypeName},
		{Method: "POST", Path: t.ItemCreateUrl, FuncName: "Create" + dataTypeName},
		{Method: "POST", Path: t.ItemUpdateUrl, FuncName: "Update" + dataTypeName},
		{Method: "POST", Path: t.ItemDeleteUrl, FuncName: "Delete" + dataTypeName},
	}
	hasPaths := []string{t.ItemsDataUrl, t.ItemCreateUrl, t.ItemUpdateUrl, t.ItemDeleteUrl}
	// type: edit,action,form,redirect
	for _, opt := range t.ItemOptions {
		if opt.Type != model.TABLE_ITEM_OPT_ACTION {
			continue
		}
		if util.GetIndexOf(opt.Url, hasPaths) > -1 {
			continue
		}
		funcName := database.TableColToObj(opt.Key) + dataTypeName
		apiRoutes = append(apiRoutes, model.ApiRoute{Method: "POST", Path: opt.Url, FuncName: "OptItem" + funcName})
		hasPaths = append(hasPaths, opt.Url)
	}
	for _, batch := range t.BatchOptButtons {
		if util.GetIndexOf(batch.Url, hasPaths) > -1 {
			continue
		}
		funcName := getFuncName(batch.Url, dataTypeName)
		apiRoutes = append(apiRoutes, model.ApiRoute{Method: "POST", Path: batch.Url, FuncName: "BatchOptItem" + funcName})
	}
	for _, fh := range t.ItemForms {
		form := fh.Form
		if util.GetIndexOf(form.SubmitUrl, hasPaths) > -1 {
			continue
		}
		funcName := getFuncName(form.SubmitUrl, dataTypeName)
		apiRoutes = append(apiRoutes, model.ApiRoute{Method: "POST", Path: form.SubmitUrl, FuncName: "FormSubmit" + funcName})
	}

	// 添加路由到服务端路由文件中 routesadd.go
	routes, err := AddApiRoutes(apiRoutes)
	if err != nil {
		return err
	}

	// 创建CURD代码
	return CreateCurdCode(routes, t)
}

func getFuncName(url, dataTypeName string) string {
	urlSplit := strings.Split(url, `/`)
	fkey := urlSplit[len(urlSplit)-1]
	return database.TableColToObj(fkey) + dataTypeName
}
