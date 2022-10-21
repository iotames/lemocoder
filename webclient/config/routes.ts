export default [
    // { path: '/', component: '@/pages/index' },
    // { path: '/welcome',name: 'welcome',icon: 'smile',component: './Welcome',},
    // {path: '/tabledemo',name: '数据表格示例',component: './TableDemo',},
    { layout: false,  path: '/public/login',  name: '',  component: './public/Login',   },
    { layout: false,  path: '/public/init',  name: '',  component: './public/AppInit',   },
    { layout: true,  path: '/welcome',  name: 'Welcome',  component: './Welcome',   },
    { layout: true,  path: '/tabledemo',  name: '数据表格示例',  component: './TableDemo',   },
    { layout: true,  path: '/codemaker',  name: '代码生成器',  component: './CodeMaker',   },
    { layout: true,  path: '/excelspider',  name: 'Excel爬虫',  component: './ExcelSpider',   },
    { layout: true,  path: '/tableschema',  name: '',  component: './TableSchema',   },
    { layout: true,  path: '/test',  name: '测试菜单1',  component: './Test',   },
    { layout: true,  path: '/',  name: '',   redirect: '/welcome',  },
    { layout: true,  name: '',  component: './404',   },
    
    // { path: '/', redirect: '/welcome' },
    // { component: './404' },
];