const webPath = process.env.NODE_ENV == "production" ? "<%{.WebPath}%>" : "/";
// const webPath = process.env.NODE_ENV == "production" ? "/client/" : "/";
export default {
    webPath: webPath,
    BaseApiUrl: "<%{.BaseApiUrl}%>", // "http://127.0.0.1:8888"
}