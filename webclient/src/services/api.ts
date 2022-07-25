// import { request } from 'umi';
import request from "@/utils/request";

/** 获取当前的用户 GET /api/ */
export async function getClientConfig(options?: { [key: string]: any }) {
  return request<{
    Data: API.ClientConfig;
  }>('/api/client/config', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function post(url: string, body: {[key: string]: any}, options?: { [key: string]: any }) {
    return request<{Code: number; Data: API.OptResult; Msg: string}>(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    });
}

/** 获取当前的用户 GET /api/user/info */
export async function getCurrentUser(options?: { [key: string]: any }) {
  return request<{
    Data: API.CurrentUser;
    Code: number;
    Msg: string;
  }>('/api/user/info', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function getMenuData(options?: { [key: string]: any }) {
  return request<{
    Data: API.MenuData;
  }>('/api/user/menu', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 发送验证码 POST /api/login/captcha */
export async function getFakeCaptcha(
  params: {
    // query
    /** 手机号 */
    phone?: string;
  },
  options?: { [key: string]: any },
) {
  return request<API.FakeCaptcha>('/api/login/captcha', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 退出登录接口 POST /api/user/logout */
export async function outLogin(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/logout', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 登录接口 POST /api/login/account */
export async function login(body: API.LoginParams, options?: { [key: string]: any }) {
  return request<{Code: number; Data: API.LoginResult; Msg: string}>('/api/public/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
