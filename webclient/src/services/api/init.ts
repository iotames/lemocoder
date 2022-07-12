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
