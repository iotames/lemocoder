export default [
    // { path: '/', component: '@/pages/index' },
    // { path: '/welcome',name: 'welcome',icon: 'smile',component: './Welcome',},
    // {path: '/tabledemo',name: '数据表格示例',component: './TableDemo',},
    <%{range .Routes}%>{ layout: <%{.Layout}%>, <%{if ne .Path "" }%> path: '<%{.Path}%>', <%{end}%> name: '<%{.Name}%>', <%{if ne .Component "" }%> component: '<%{.Component}%>', <%{end}%> <%{if ne .Redirect "" }%> redirect: '<%{.Redirect}%>', <%{end}%> },
    <%{end}%>
    // { path: '/', redirect: '/welcome' },
    // { component: './404' },
];