import { LockOutlined, MobileOutlined, UserOutlined } from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import { ProFormCaptcha, ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { history, useModel } from 'umi';
import Footer from '@/components/Footer';
import {login, getFakeCaptcha, getCurrentUser} from "@/services/api"
import styles from './index.less';

const LoginMessage: React.FC<{
  content: string;
}> = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Login: React.FC = () => {
  const defaultLoginState = {Code: 500, Msg: "unknown", Data:{Token:"", ExpiresIn:0}}
  const [userLoginState, setUserLoginState] = useState<{Code: number; Msg: string; Data: API.LoginResult}>(defaultLoginState);
  const [LoginWay, setLoginWay] = useState<string>('account');
  const {initialState, setInitialState } = useModel('@@initialState');

  const handleSubmit = async (values: API.LoginParams) => {
    try {
      // 登录
      const resp = await login({ ...values, LoginWay});
      setUserLoginState(resp);
      if (resp.Code === 200) {
        if (resp.Data.Token && initialState) {
          localStorage.setItem('AuthToken', resp.Data.Token);
          const defaultLoginSuccessMessage = "登录成功";
          message.success(defaultLoginSuccessMessage);
          const currentUser = (await getCurrentUser()).Data
          initialState.currentUser = currentUser
          await setInitialState(initialState)
          if (!history){
            console.log("-----History fALSE", history)
            return
          }
          /** 此方法会跳转到 redirect 参数所在的位置 */
          const { query } = history.location;
          const { redirect } = query as { redirect: string };
          history.push(redirect || '/');
          return;
        } else {
          // TODO 登录后未自动跳转到首页。 initialState == undefined
          location.reload();
          message.error('AuthToken or initialState is undefined');
          return;
        }
      }else{
        message.error(resp.Msg);
      }

    } catch (error) {
      const defaultLoginFailureMessage = "登录失败，请重试！";
      message.error(defaultLoginFailureMessage);
      console.log(error);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <LoginForm
          logo={<img alt="logo" src="logo.svg" />}
          title="LemoCoder"
          subTitle="代码生成器"
          initialValues={{
            autoLogin: true,
          }}
          
          onFinish={async (values) => {
            await handleSubmit(values as API.LoginParams);
          }}
        >
          <Tabs activeKey={LoginWay} onChange={setLoginWay}>
            <Tabs.TabPane
              key="account"
              tab={"账户密码登录"}
            />
{/* 
            <Tabs.TabPane
              key="mobile"
              tab="手机号登录"
            />
*/}
          </Tabs>

          {userLoginState.Code == 400 && LoginWay === 'account' && (
            <LoginMessage content="账户或密码错误" />
          )}
          {LoginWay === 'account' && (
            <>
              <ProFormText
                name="Username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={'用户名:'}
                rules={[
                  {
                    required: true,
                    message: "请输入用户名!",
                  },
                ]}
              />
              <ProFormText.Password
                name="Password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder={'密码:'}
                rules={[
                  {
                    required: true,
                    message: "请输入密码！",
                  },
                ]}
              />
            </>
          )}

          {userLoginState.Code == 400 && LoginWay === 'mobile' && <LoginMessage content="验证码错误" />}

          {LoginWay === 'mobile' && (
            <>
              <ProFormText
                fieldProps={{
                  size: 'large',
                  prefix: <MobileOutlined className={styles.prefixIcon} />,
                }}
                name="mobile"
                placeholder={"手机号"}
                rules={[
                  {
                    required: true,
                    message: "请输入手机号！",
                  },
                  {
                    pattern: /^1\d{10}$/,
                    message: "手机号格式错误！",
                  },
                ]}
              />
              <ProFormCaptcha
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                captchaProps={{
                  size: 'large',
                }}
                placeholder={"请输入验证码"}
                captchaTextRender={(timing, count) => {
                  if (timing) {
                    return `${count} 获取验证码`;
                  }
                  return "获取验证码";
                }}
                name="captcha"
                rules={[
                  {
                    required: true,
                    message: "请输入验证码！",
                  },
                ]}
                onGetCaptcha={async (phone) => {
                  const result = await getFakeCaptcha({
                    phone,
                  });
                  if (result === false) {
                    return;
                  }
                  message.success('获取验证码成功！验证码为：1234');
                }}
              />
            </>
          )}
          <div
            style={{
              marginBottom: 24,
            }}
          >
            <ProFormCheckbox noStyle name="AutoLogin">
            自动登录
            </ProFormCheckbox>
            {/* <a style={{float: 'right'}}>忘记密码</a> */}
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Login;
