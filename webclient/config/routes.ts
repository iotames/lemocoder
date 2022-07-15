export default [
    // { path: '/', component: '@/pages/index' },
    {
        layout: false,
        path: '/public/login',
        // name: 'ligin',
        component: './public/Login',
    },
    {
        path: '/welcome',
        name: 'welcome',
        icon: 'smile',
        component: './Welcome',
    },
    {
        path: '/forms',
        name: '表单管理',
        component: './FormsList',
    },
    {
        path: '/formgen',
        name: '表单生成器',
        component: './Formgen',
    },
    { path: '/', redirect: '/welcome' },
    { component: './404' },
];