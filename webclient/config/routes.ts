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
    { path: '/', redirect: '/welcome' },
    { component: './404' },
];