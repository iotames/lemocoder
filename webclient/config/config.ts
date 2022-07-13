import { defineConfig } from 'umi';
import routes from './routes';
import defaultSettings from './defaultLayoutSetting';

// https://v3.umijs.org/zh-CN/docs/env-variables
// define in command: PORT=8000 umi dev   use in source code: process.env.PORT

const webPath = process.env.NODE_ENV == "production" ? "/client/" : "/";

export default defineConfig({
    base: webPath,
    publicPath: webPath,
    nodeModulesTransform: {type: "none"},
    history: {type:"hash"},
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