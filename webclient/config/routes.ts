export default [
    // { path: '/', component: '@/pages/index' },
    {
        layout: false,
        path: '/public/login',
        // name: 'ligin',
        component: './public/Login',
    },

    {
        layout: false,
        path: '/public/init',
        component: './public/AppInit',
    },

    {
        path: '/welcome',
        name: 'welcome',
        icon: 'smile',
        component: './Welcome',
    },
    {
        path: '/tabledemo',
        name: '数据表格示例',
        component: './TableDemo',
    },
    {
        path: '/codemaker',
        name: '代码生成器',
        component: './CodeMaker',
    },
    {
        path: '/excelspider',
        name: 'Excel爬虫',
        component: './ExcelSpider',
    },
    { path: '/', redirect: '/welcome' },
    { component: './404' },
];