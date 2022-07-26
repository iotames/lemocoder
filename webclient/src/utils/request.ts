// An highlighted block
/**
 * request 网络请求工具
 * 更详细的 api 文档: https://github.com/umijs/umi-request
 * @link https://blog.csdn.net/weixin_41753520/article/details/98317567
 */
 import { extend, RequestOptionsInit } from 'umi-request';
 import { notification } from 'antd';
 import config from "../../config/webconf"
 
 const codeMessage: {[key: number]: string} = {
   200: '服务器成功返回请求的数据。',
   201: '新建或修改数据成功。',
   202: '一个请求已经进入后台排队（异步任务）。',
   204: '删除数据成功。',
   400: '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
   401: '用户没有权限（令牌、用户名、密码错误）。',
   403: '用户得到授权，但是访问是被禁止的。',
   404: '发出的请求针对的是不存在的记录，服务器没有进行操作。',
   406: '请求的格式不可得。',
   410: '请求的资源被永久删除，且不会再得到的。',
   422: '当创建一个对象时，发生一个验证错误。',
   500: '服务器发生错误，请检查服务器。',
   502: '网关错误。',
   503: '服务不可用，服务器暂时过载或维护。',
   504: '网关超时。',
 };
 
 /**
  * 异常处理程序
  */
 const errorHandler = (error: { response: any; }) => {
   const { response } = error;
   if (response && response.status) {
     const errorText = codeMessage[response.status] || response.statusText;
     const { status, url } = response;
     notification.error({
       message: `请求错误 ${status}: ${url}`,
       description: errorText,
     });
   }
 
   return response;
 };
 
 const request = extend({
   errorHandler,
   // 默认错误处理
   // credentials: 'include', // 跨域去除 include 。默认请求是否带上cookie
   prefix: config.BaseApiUrl,
   // getResponse: true,
 });
 
 // request拦截器, 改变url 或 options.
 request.interceptors.request.use(async (url: string, options: RequestOptionsInit) => {
  // request.interceptors.request.use(async (url, options) => {
   // console.log(options);
   let c_token = localStorage.getItem('AuthToken');
   // 设置请求头可以解决跨域问题
 
   if (c_token) {
     let headers = {
       'Content-Type': 'application/json',
       Accept: 'application/json',
       'Auth-Token': c_token,
       // 'Sec-Fetch-Site': 'same-origin',
       // 'Sec-Fetch-Mode': 'cors',
       // 'Sec-Fetch-Dest': 'empty',
     };
     // headers['x-auth-token'] = c_token; A BUG heppend
     return {
       url: url,
       options: { ...options, headers: headers },
     };
   }
   return {
     url: url,
     options: { ...options },
   };
 });

 
 
 // response拦截器, 处理response
 // request.interceptors.response.use((response, options) => {
 //   console.log('-------------interceptors------response start-----');
 //   console.log(response.headers.get('x-auth-token'));
 //   console.log(response.headers.values);
 //   console.log(response.data);
 //   console.log('----------interceptors---------response end-----');
 
 //   return response;
 // });
 
 export default request;