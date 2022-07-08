import { defineConfig } from 'umi';
import routes from './routes';

export default defineConfig({
    nodeModulesTransform: {type: "none"},
    fastRefresh: {},
    routes: routes,
    layout: {},
    mfsu: {},
});


// export default defineConfig({
//   nodeModulesTransform: {
//     type: 'none',
//   },
//   fastRefresh: {},
// });