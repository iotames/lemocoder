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
	database.GetAll(&pages, 1000, 1, "project_id = ?", 0)

	// 重建客户端路由数据
	var rts []initial.ClientRoute
	addRoute := true
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
	return CreateFile(config.ClientConfigDir+"/routes.ts", config.TplDirPath+"/routes.ts.tpl", dt1)
}

// CreateTableServer 创建服务端源码文件
func CreateTableServer(t model.TableSchema, p database.WebPage) error {
	return nil
}
