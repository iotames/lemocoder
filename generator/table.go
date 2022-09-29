package generator

import (
	"lemocoder/config"
	"lemocoder/initial"
)

func (t TableSchema) Create() error {
	err := CreateFile(config.ClientSrcPagesDir+"/Test.tsx", config.TplDirPath+"/table.tsx.tpl", t)
	routes := initial.GetClientRoutes(initial.ClientRoute{Name: "test", Path: "/test", Component: "./Test", Layout: true})
	dt1 := map[string]interface{}{"Routes": routes}
	if err != nil {
		return err
	}
	return CreateFile(config.ClientConfigDir+"/routes.ts", config.TplDirPath+"/routes.ts.tpl", dt1)
}
