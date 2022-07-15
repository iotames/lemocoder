import { LockOutlined, MobileOutlined, UserOutlined } from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import { ProFormCaptcha, ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { history, useModel } from 'umi';
import Footer from '@/components/Footer';
import { login } from '@/services/antdprodemo/api';
import { getFakeCaptcha } from '@/services/antdprodemo/login';

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
  const [userLoginState, setUserLoginState] = useState<{Code: number; Msg: string; Data: API.LoginResult}>({Code: 500, Msg: "unknown", Data:{}});
  const [LoginWay, setLoginWay] = useState<string>('account');
  const { initialState, setInitialState } = useModel('@@initialState');

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();
    console.log('----fetchUserInfo-----------');
    console.log(userInfo);
    if (userInfo) {
      await setInitialState((s) => ({
        ...s,
        currentUser: userInfo,
      }));
    }
  };

  const handleSubmit = async (values: API.LoginParams) => {
    try {
      // 登录
      const msg = await login({ ...values, LoginWay});
      if (msg.Code === 200) {
        if (msg.Data.Token) {
          localStorage.setItem('AuthToken', msg.Data.Token);
        } else {
          message.error('AuthToken error');
          return;
        }

        const defaultLoginSuccessMessage = "登录成功";
        message.success(defaultLoginSuccessMessage);
        await fetchUserInfo();
        /** 此方法会跳转到 redirect 参数所在的位置 */
        if (!history) return;
        const { query } = history.location;
        const { redirect } = query as { redirect: string };
        history.push(redirect || '/');
        return;
      }

      // 如果失败去设置用户错误信息
      setUserLoginState(msg);
      message.error(msg.Msg);
    } catch (error) {
      const defaultLoginFailureMessage = "登录失败，请重试！";
      message.error(defaultLoginFailureMessage);
    }
  };

  const { Code, Msg } = userLoginState;
  // let config = initialState?.config
  // if (config === undefined) {
  //   config = initialState.fetchConfig()
  // }

  console.log(initialState?.config)

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
            <Tabs.TabPane
              key="mobile"
              tab="手机号登录"
            />
          </Tabs>

          {Code == 400 && Msg === 'account' && (
            <LoginMessage content="账户或密码错误(admin/ant.design)" />
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

          {Code === 400 && LoginWay === 'mobile' && <LoginMessage content="验证码错误" />}
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