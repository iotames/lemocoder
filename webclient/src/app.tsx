// import Footer from '@/components/Footer';
import RightContent from '@/components/RightContent';
import { LinkOutlined } from '@ant-design/icons';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';
import { SettingDrawer, PageLoading } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from 'umi';
import { history, Link } from 'umi';
import defaultSettings from '../config/defaultLayoutSetting';

import { getClientConfig, getMenuData, getCurrentUser } from '@/services/api';

const isDev = process.env.NODE_ENV === 'developmentNOT';
const loginPath = '/public/login';
const initPath = "/public/init";

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  settings: Partial<LayoutSettings>;
  config: API.ClientConfig;
  currentUser: API.CurrentUser;
  menuItems?: API.MenuItem[];
  loading?: boolean;
}> {
  const currentUser = {Id: 0, Account: ""}
  const config = (await getClientConfig()).Data
  if (history.location.pathname !== loginPath) {
    // 如果不是登录页面，执行
    const currentUser = (await getCurrentUser()).Data
    const menuItems = (await getMenuData()).Data.Items
    return {
      menuItems,
      currentUser,
      config,
      settings: defaultSettings,
    };
  }
  return {
    config,
    settings: defaultSettings,
    currentUser,
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState, setInitialState }) => {
  return {
    menu: {
      // 每当 initialState?.currentUser?.id 发生修改时重新执行 request
      params: initialState,
      request: async (params, defaultMenuData) => {
        return initialState?.menuItems;
      },
    },
    title: initialState?.config.Title,
    logo: initialState?.config.Logo,
    
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    waterMarkProps: {
      content: initialState?.currentUser?.Account,
    },
    // footerRender: () => <Footer />,
    onPageChange: () => {
      const { location } = history;
      console.log(history);
      console.log("initialState?.currentUser?.Account---", initialState?.currentUser?.Account)
      // 如果没有登录，重定向到 login
      if (location.pathname != loginPath && location.pathname != initPath){
        if (!initialState?.currentUser?.Account){
          history.push(loginPath);
          return
          // if (!initialState){
          //   history.push(loginPath);
          //   return
          // }

          // getCurrentUser().then((resp)=>{
          //   if (resp.Code!=200){
          //     history.push(loginPath);
          //     return
          //   }
          //   const currentUser = resp.Data;
          //   initialState.currentUser = currentUser
          //   getMenuData().then((res)=>{
          //     const menuItems = res.Data.Items
          //     initialState.menuItems = menuItems
          //     setInitialState(initialState)
          //   })
          // })

        }
      }
      if (initialState?.config && initialState.config.IsLocked && location.pathname == initPath){
        history.push("/")
      }
      if (initialState?.config && !initialState.config.IsLocked) {
        history.push(initPath);
      }
    },
    links: isDev
      ? [
          <Link key="openapi" to="/umi/plugin/openapi" target="_blank">
            <LinkOutlined />
            <span>OpenAPI 文档</span>
          </Link>,
        ]
      : [],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children, props) => {
      if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          {!props.location?.pathname?.includes('/login') && (
            <SettingDrawer
              disableUrlParams
              enableDarkTheme
              settings={initialState?.settings}
              onSettingChange={(settings) => {
                setInitialState((preInitialState) => ({
                  ...preInitialState,
                  settings,
                }));
              }}
            />
          )}
        </>
      );
    },
    ...initialState?.settings,
  };
};
