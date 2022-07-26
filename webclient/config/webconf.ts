const webPath = process.env.NODE_ENV == "production" ? "/client/" : "/";
export default {
    webPath: webPath,
    BaseApiUrl: "http://127.0.0.1:8888"
}

// import { request } from 'umi';
// import webconf from '../../config/webconf';
// const webPath = webconf.webPath;
// const conf = await request<{BaseApiUrl: string}>(webPath + "config.json")
// export default conf