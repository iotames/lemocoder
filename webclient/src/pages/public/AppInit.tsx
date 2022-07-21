import { ProFormDigit, ProFormInstance } from '@ant-design/pro-components';
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
  ProFormGroup,
} from '@ant-design/pro-components';
import { message } from 'antd';
import { useRef, useState } from 'react';
import { useModel } from 'umi';

const { initialState, loading} = useModel('@@initialState');


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
  const dbDrivers = [{value: 'sqlite3', label: 'Sqlite3'},{value: 'mysql', label: 'Mysql'}]
  const submitFormFunc = (val: any) => {
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
          message.success('提交成功');
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
          // onFinish={async () => {
          //   console.log(formRef.current?.getFieldsValue());
          //   await waitTime(2000);
          //   return true;
          // }}
        >

<ProFormSelect label="数据库类型" width="sm" name="DbDriver" rules={[{required: true}]} initialValue="sqlite3" options={dbDrivers} onChange={submitFormFunc} />
          
<ProFormGroup >
  <ProFormText hidden={isSqlite} name="DbHost" label="主机名" width="sm" tooltip="填IP或网址" placeholder="请输入数据库主机地址" initialValue="127.0.0.1" rules={[{ required: true }]} />
  <ProFormDigit hidden={isSqlite} name="DbPort" min={1000} max={9999} label="端口号" width="xs" tooltip="填数字" initialValue="3306" rules={[{ required: true }]} />
  <ProFormDigit hidden={isSqlite} name="DbNodeId" width="xs" label="节点ID" initialValue="1" rules={[{ required: true }]} />
</ProFormGroup>

<ProFormGroup >
    <ProFormText hidden={isSqlite} name="DbName" label="数据库名" width="sm" initialValue="lemocoder" rules={[{ required: true }]} />
    <ProFormText hidden={isSqlite} name="DbUsername" label="数据库账号名" width="sm" initialValue="root" rules={[{ required: true }]} />
    <ProFormText hidden={isSqlite} name="DbPassword" label="数据库密码" width="sm" initialValue="root" rules={[{ required: true }]} />
    
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
          <ProFormDigit name="WebPort" min={1000} max={9999} label="Web端口号" width="xs" tooltip="填数字" initialValue="8888" rules={[{ required: true }]} />

          {/* <ProFormCheckbox.Group name="checkbox" label="迁移类型" width="lg" options={['结构迁移', '全量迁移', '增量迁移', '全量校验']} /> */}
        </StepsForm.StepForm>


        <StepsForm.StepForm
          name="webclient"
          title="Web客户端"
          stepProps={{
            description: 'Ant Design Pro',
          }}
        >
          <ProFormText name="ClientTitle" label="站点标题" width="md" initialValue="QingCoder" rules={[{ required: true }]} />
        {/* <ProFormCheckbox.Group name="checkbox" label="部署单元" rules={[{required: true}]} options={['部署单元1', '部署单元2', '部署单元3']} /> */}
        </StepsForm.StepForm>

      </StepsForm>
    </ProCard>
    </div>
  );
};