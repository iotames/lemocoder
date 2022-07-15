// @ts-ignore
/* eslint-disable */

declare namespace API {
  type CurrentUser = {
    id: number,
    account: string;
    name?: string;
    avatar?: string;
    mobile?: string;
  };

  type LoginResult = {
    Token?: string;
    ExpiresIn?: number;
  };

  type OptResult = {
    Code: number;
    Msg: string;
    Data?: {};
  }

  type ClientConfig = {
      Title?: string;
      Desc?: string;
  }

  type MenuItem = {
    component:string; 
    layout?:boolean; 
    name:string;
    icon?:string;
    path?:string;
    redirect?:string; 
  }

  type MenuData = {
    Items: MenuItem[]
  }

  type PageParams = {
    current?: number;
    pageSize?: number;
  };

  type RuleListItem = {
    key?: number;
    disabled?: boolean;
    href?: string;
    avatar?: string;
    name?: string;
    owner?: string;
    desc?: string;
    callNo?: number;
    status?: number;
    updatedAt?: string;
    createdAt?: string;
    progress?: number;
  };

  type RuleList = {
    data?: RuleListItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type LoginParams = {
    Username?: string;
    Password?: string;
    AutoLogin?: boolean;
    LoginWay?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };
}