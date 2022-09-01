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

func GetClientRoutes(routes ...ClientRoute) []ClientRoute {
	// { path: '/', component: '@/pages/index' },
	r1 := []ClientRoute{
		{false, "/public/login", "./public/Login", "", ""},
		{false, "/public/init", "./public/AppInit", "", ""},
		{true, "/welcome", "./Welcome", "Welcome", ""},
		{true, "/tabledemo", "./TableDemo", "数据表格示例", ""},
		{true, "/codemaker", "./CodeMaker", "代码生成器", ""},
		{true, "/excelspider", "./ExcelSpider", "Excel爬虫", ""},
	}
	r1 = append(r1, routes...)
	r1 = append(r1, ClientRoute{true, "/", "", "", "/welcome"}, ClientRoute{true, "", "./404", "", ""})
	return r1
}

// export default [
//     <%{range .Routes}%>
//     { layout: <%{.Layout}%>, path: '<%{.Path}%>', name: '<%{.Name}%>', component: '<%{.Component}%>', },
//     <%{end}%>
//     // { path:"/test", name:"test", component:"./Test"},
//     { path: '/', redirect: '/welcome' },
//     { component: './404' },
// ];
