import type { ProFormInstance } from '@ant-design/pro-components';
import {
  ProCard,
  ProForm,
  ProFormCheckbox,
  ProFormDatePicker,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  StepsForm,
} from '@ant-design/pro-components';
import { message } from 'antd';
import { useRef, useState } from 'react';


const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export default () => {
  const [isSqlite, setIsSqlite] = useState(true)
  const formRef = useRef<ProFormInstance>();

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
          message.success('提交成功');
        }}
        formProps={{
          validateMessages: {
            required: '此项为必填项',
          },
        }}
      >
        <StepsForm.StepForm<{
          name: string;
        }>
          name="DbSetting"
          title="数据库设置"
          stepProps={{
            description: '这里初始化数据库设置',
          }}
          onChange={(values) => {
            console.log(values)
          }}
          // onFinish={async () => {
          //   console.log(formRef.current?.getFieldsValue());
          //   await waitTime(2000);
          //   return true;
          // }}
        >

          <ProFormSelect
            label="数据库类型"
            name="DbDriver"
            rules={[
              {
                required: true,
              },
            ]}
            initialValue="sqlite"
            options={[
              {
                value: 'sqlite',
                label: 'Sqlite3',
              },
              { value: 'mysql', label: 'Mysql' },
            ]}
            onChange={(val)=>{
              const isSqlite = val == "sqlite"
              setIsSqlite(isSqlite)
            }}
          />
          
          <ProFormText
          hidden={isSqlite}
            name="DbHost"
            label="主机名"
            width="md"
            tooltip="填IP或网址"
            placeholder="请输入数据库主机地址"
            initialValue="127.0.0.1"
            rules={[{ required: true }]}
          />

          <ProFormText hidden={isSqlite} name="DbPort" label="端口号" width="md" tooltip="数据库端口号" initialValue="3306" rules={[{ required: true }]} />
          <ProFormText hidden={isSqlite} name="DbName" label="数据库名" width="md" initialValue="lemocoder" rules={[{ required: true }]} />
          <ProFormText hidden={isSqlite} name="DbUsername" label="数据库账号名" width="md" initialValue="root" rules={[{ required: true }]} />
          <ProFormText hidden={isSqlite} name="DbPassword" label="数据库密码" width="md" initialValue="root" rules={[{ required: true }]} />
          <ProFormText hidden={isSqlite} name="DbNodeId" label="数据库节点" width="md" initialValue="1" rules={[{ required: true }]} />



# Server
WEB_SERVER_PORT = 8888
WEB_SERVER_URL = "http://127.0.0.1:8888"


          {/* <ProFormTextArea name="remark" label="备注" width="lg" placeholder="请输入备注" /> */}

        </StepsForm.StepForm>

        <StepsForm.StepForm<{
          checkbox: string;
        }>
          name="checkbox"
          title="设置参数"
          stepProps={{
            description: '这里填入运维参数',
          }}
          // onFinish={async () => {
          //   console.log(formRef.current?.getFieldsValue());
          //   return true;
          // }}
        >
          <ProFormText name="dbname" label="业务 DB 用户名" />
          <ProFormCheckbox.Group
            name="checkbox"
            label="迁移类型"
            width="lg"
            options={['结构迁移', '全量迁移', '增量迁移', '全量校验']}
          />

        </StepsForm.StepForm>

        <StepsForm.StepForm
          name="time"
          title="发布实验"
          stepProps={{
            description: '这里填入发布判断',
          }}
        >
          <ProFormCheckbox.Group
            name="checkbox"
            label="部署单元"
            rules={[
              {
                required: true,
              },
            ]}
            options={['部署单元1', '部署单元2', '部署单元3']}
          />
          <ProFormSelect
            label="部署分组策略"
            name="remark"
            rules={[
              {
                required: true,
              },
            ]}
            initialValue="1"
            options={[
              {
                value: '1',
                label: '策略一',
              },
              { value: '2', label: '策略二' },
            ]}
          />
        </StepsForm.StepForm>

      </StepsForm>
    </ProCard>
    </div>
  );
};