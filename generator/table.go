package generator

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/initial"
	"lemocoder/model"
)

// CreateTableClient 创建客户端源码文件
func CreateTableClient(t model.TableSchema, p database.WebPage) error {
	// 创建数据表格页面的tsx文件
	err := CreateFile(fmt.Sprintf("%s/%s.tsx", config.ClientSrcPagesDir, p.Component), config.TplDirPath+"/table.tsx.tpl", t)
	if err != nil {
		return err
	}
	pages := make([]database.WebPage, 0)
	// 获取所有页面
	database.GetAll(&pages, 1000, 1, "project_id = ? AND state = ?", 0, 1)

	// 重建客户端路由数据
	var rts []initial.ClientRoute
	addRoute := true
	if p.Path == "/test" {
		addRoute = false
	}
	for _, page := range pages {
		rts = append(rts, initial.ClientRoute{Name: page.Name, Path: page.Path, Component: "./" + page.Component, Layout: true})
		if page.Path == p.Path {
			addRoute = false
		}
	}
	if addRoute {
		rts = append(rts, initial.ClientRoute{Name: p.Name, Path: p.Path, Component: "./" + p.Component, Layout: true})
	}

	routes := initial.GetClientRoutes(rts...)
	dt1 := map[string]interface{}{"Routes": routes}

	// 重建routes.ts客户端路由文件
	return CreateFile(config.ClientRoutesPath, config.TplDirPath+"/routes.ts.tpl", dt1)
}

// CreateTableServer 创建服务端源码文件
func CreateTableServer(t model.TableSchema) error {
	// 创建服务端API源码文件
	AddDbModel(t.Items, t.ItemDataTypeName)

	// 重建服务端路由文件 routesadd.go
	var apiRoutes []model.ApiRoute
	// type: edit,action,form,redirect
	for _, opt := range t.ItemOptions {
		// TODO
		if opt.Type == "edit" {

		}
		if opt.Type == "action" {

		}
	}
	if len(apiRoutes) > 0 {
		return AddApiRoutes(apiRoutes)
	}
	return nil
}
