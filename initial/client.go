package initial

type ClientMenuItem struct {
	// Layout    bool   `json:"layout"`
	// Component string `json:"component"`
	Name string `json:"name"`
	Path string `json:"path"`
	// Redirect string `json:"redirect"`
	// Icon      string `json:"icon"`
}
type ClientRoute struct {
	Layout                          bool
	Path, Component, Name, Redirect string
}

// func GetClientRoutes(routes ...ClientRoute) []ClientRoute {
// 	// { path: '/', component: '@/pages/index' },
// 	r1 := []ClientRoute{
// 		{false, "/public/login", "./public/Login", "", ""},
// 		{false, "/public/init", "./public/AppInit", "", ""},
// 		{true, "/welcome", "./Welcome", "Welcome", ""},
// 		{true, "/tabledemo", "./TableDemo", "数据表格示例", ""},
// 		{true, "/codemaker", "./CodeMaker", "代码生成器", ""},
// 		{true, "/test", "./Test", "测试页面", ""},
// 		{true, "/excelspider", "./ExcelSpider", "Excel爬虫", ""},
// 		{true, "/tableschema", "./TableSchema", "", ""},
// 	}
// 	r1 = append(r1, routes...)
// 	r1 = append(r1, ClientRoute{true, "/", "", "", "/welcome"}, ClientRoute{true, "", "./404", "", ""})
// 	return r1
// }

func GetClientMenu(menu ...ClientMenuItem) []ClientMenuItem {
	initMenu := []ClientMenuItem{
		{Path: "/welcome", Name: "首页"},
		// {Path: "/tabledemo", Name: "数据表格示例"},
		{Path: "/codemaker", Name: "代码生成器"},
		// {Path: "/excelspider", Name: "Excel爬虫"},
		// {Path: "/test", Name: "测试1"},
	}
	initMenu = append(initMenu, menu...)
	return initMenu
}
