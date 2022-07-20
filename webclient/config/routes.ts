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
        path: '/excelspider',
        name: 'Excel爬虫',
        component: './ExcelSpider',
    },
    { path: '/', redirect: '/welcome' },
    { component: './404' },
];