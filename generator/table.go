package generator

import (
	"fmt"
	"lemocoder/config"
	"lemocoder/database"
	"lemocoder/initial"
	"lemocoder/model"
)

func CreateTableClient(t model.TableSchema, p database.WebPage) error {
	err := CreateFile(fmt.Sprintf("%s/%s.tsx", config.ClientSrcPagesDir, p.Component), config.TplDirPath+"/table.tsx.tpl", t)
	if err != nil {
		return err
	}
	pages := make([]database.WebPage, 0)
	database.GetAll(&pages, 1000, 1, "project_id = ?", 0)
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
	return CreateFile(config.ClientConfigDir+"/routes.ts", config.TplDirPath+"/routes.ts.tpl", dt1)
}
