import { defineConfig } from 'umi';
import routes from './routes';
import defaultSettings from './defaultLayoutSetting';

export default defineConfig({
    nodeModulesTransform: {type: "none"},
    fastRefresh: {},
    routes: routes,
    layout: {
        // https://umijs.org/zh-CN/plugins/plugin-layout
        locale: true,
        siderWidth: 208,
        ...defaultSettings,
    },
    mfsu: {},
});


// export default defineConfig({
//   nodeModulesTransform: {
//     type: 'none',
//   },
//   fastRefresh: {},
// });