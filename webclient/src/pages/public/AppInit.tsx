import {
  ProCard,
  ProFormDigit, 
  ProFormInstance,
  ProFormSelect,
  ProFormText,
  StepsForm,
  ProFormGroup,
} from '@ant-design/pro-components';
import { message, Spin } from 'antd';
import { useRef, useState } from 'react';
import { useModel, history } from 'umi';
import { post } from "@/services/api"

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export default () => {
  const { initialState, loading, setInitialState } = useModel('@@initialState');

  if (loading || initialState==undefined){
    return <Spin />
  }
  const initConf = initialState.config
  const [isSqlite, setIsSqlite] = useState(true)
  const formRef = useRef<ProFormInstance>();
  const dbDrivers = [{value: 'sqlite3', label: 'Sqlite3'},{value: 'mysql', label: 'Mysql'}]
  const changeDbDriver = (val: any) => {
    const isSqlite = val == "sqlite3"
    setIsSqlite(isSqlite)
  }

  return (
    <div>
    <ProCard>
      <StepsForm<{
        name: string;
      }>
        formRef={formRef}
        onFinish={async (values) => {
          await waitTime(1000);
          console.log(values)
          const resp = (await post("/api/client/init", values))
          if (resp.Code == 200){
            message.success('提交成功');
            initialState.config.IsLocked = true
            setInitialState({...initialState})
            history.push("/")
          }else{
            message.error(resp.Msg);
          }
          
        }}
        formProps={{
          validateMessages: {
            required: '此项为必填项',
          },
        }}
      >
        <StepsForm.StepForm<{}>
          name="database"
          title="数据库设置"
          stepProps={{
            description: 'XORM引擎',
          }}
          onChange={(values) => {
            console.log(values)
          }}
        >
          <ProFormSelect label="数据库类型" width="sm" name="DbDriver" rules={[{required: true}]} initialValue={initConf.DbDriver} options={dbDrivers} onChange={changeDbDriver} />
          
<ProFormGroup >
  <ProFormText hidden={isSqlite} name="DbHost" label="主机名" width="sm" tooltip="填IP或网址" placeholder="请输入数据库主机地址" initialValue={initConf.DbHost} rules={[{ required: true }]} />
  <ProFormDigit hidden={isSqlite} name="DbPort" min={1000} max={9999} label="端口号" width="xs" tooltip="填数字" initialValue={initConf.DbPort} rules={[{ required: true }]} />
  <ProFormDigit hidden={isSqlite} name="DbNodeId" width="xs" label="节点ID" initialValue={initConf.DbNodeId} rules={[{ required: true }]} />
</ProFormGroup>

<ProFormGroup >
    <ProFormText hidden={isSqlite} name="DbName" label="数据库名" width="sm" initialValue={initConf.DbName} rules={[{ required: true }]} />
    <ProFormText hidden={isSqlite} name="DbUsername" label="数据库账号名" width="sm" initialValue={initConf.DbUsername} rules={[{ required: true }]} />
    <ProFormText hidden={isSqlite} name="DbPassword" label="数据库密码" width="sm" initialValue={initConf.DbPassword} rules={[{ required: true }]} />
    
</ProFormGroup>

          {/* <ProFormTextArea name="remark" label="备注" width="lg" placeholder="请输入备注" /> */}
        </StepsForm.StepForm>


        <StepsForm.StepForm<{}>
          name="webserver"
          title="Web服务器"
          stepProps={{
            description: '基于Gin框架',
          }}
          // onFinish={async () => {
          //   console.log(formRef.current?.getFieldsValue());
          //   return true;
          // }}
        >
          <ProFormDigit name="WebServerPort" min={1000} max={9999} label="Web端口号" width="xs" tooltip="填数字" initialValue={initConf.WebServerPort} rules={[{ required: true }]} />

          {/* <ProFormCheckbox.Group name="checkbox" label="迁移类型" width="lg" options={['结构迁移', '全量迁移', '增量迁移', '全量校验']} /> */}
        </StepsForm.StepForm>


        <StepsForm.StepForm
          name="webclient"
          title="Web客户端"
          stepProps={{
            description: 'Ant Design Pro',
          }}
        >
          <ProFormText name="Title" label="站点标题" width="md" initialValue={initConf.Title} rules={[{ required: true }]} />
          <ProFormText name="LoginAccount" label="登录名" width="md" initialValue={initConf.LoginAccount} rules={[{ required: true }]} />
          <ProFormText.Password name="LoginPassword" label="登录密码" width="md" initialValue={initConf.LoginPassword} rules={[{ required: true }]} />

        {/* <ProFormCheckbox.Group name="checkbox" label="部署单元" rules={[{required: true}]} options={['部署单元1', '部署单元2', '部署单元3']} /> */}
        </StepsForm.StepForm>

      </StepsForm>
    </ProCard>
    </div>
  );
};