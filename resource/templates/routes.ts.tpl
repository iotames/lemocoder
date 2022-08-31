export default [
    // { path: '/', component: '@/pages/index' },
    // {layout: false,path: '/public/login',component: './public/Login',},
    // { layout: false,path: '/public/init',component: './public/AppInit',},
    // { path: '/welcome',name: 'welcome',icon: 'smile',component: './Welcome',},
    // {path: '/tabledemo',name: '数据表格示例',component: './TableDemo',},
    // {path: '/codemaker', name: '代码生成器',component: './CodeMaker'},
    // { path: '/excelspider',name: 'Excel爬虫',component: './ExcelSpider',},
    <%{range .Routes}%>
    { layout: <%{.Layout}%>, <%{if ne .Path "" }%> path: '<%{.Path}%>', <%{end}%> name: '<%{.Name}%>', <%{if ne .Component "" }%> component: '<%{.Component}%>', <%{end}%> <%{if ne .Redirect "" }%> redirect: '<%{.Redirect}%>', <%{end}%> },
    <%{end}%>
    // { path:"/test", name:"test", component:"./Test"},
    // { path: '/', redirect: '/welcome' },
    // { component: './404' },
];