// import { request } from 'umi';
import request from "@/utils/request";

/** 获取当前的用户 GET /api/ */
export async function getClientConfig(options?: { [key: string]: any }) {
  return request<{
    data: API.ClientConfig;
  }>('/api/client/config', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function post(url: string, body: API.OptResult, options?: { [key: string]: any }) {
    return request<{Code: number; Data: API.LoginResult; Msg: string}>(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    });
  }