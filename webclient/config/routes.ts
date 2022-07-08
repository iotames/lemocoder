export default [
    // { path: '/', component: '@/pages/index' },
    {
        path: '/welcome',
        name: 'welcome',
        icon: 'smile',
        component: './Welcome',
    },
    { path: '/', redirect: '/welcome' },
    { component: './404' },
];