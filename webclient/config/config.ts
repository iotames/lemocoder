import { defineConfig } from 'umi';
import routes from './routes';
import defaultSettings from './defaultLayoutSetting';
import webconf from './webconf';

// https://v3.umijs.org/zh-CN/docs/env-variables
// define in command: PORT=8000 umi dev   use in source code: process.env.PORT
const webPath = webconf.webPath

export default defineConfig({
    base: webPath,
    publicPath: webPath,
    // webpack5: {mode:"development"}, // experiments:{topLevelAwait: true}
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
    // mfsu: {}, // https://github.com/umijs/umi/issues/7746
});
// yarn add @umijs/plugin-webpack-5
// yarn remove @umijs/plugin-webpack-5