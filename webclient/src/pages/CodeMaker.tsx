import { ModalForm, ProFormText, PageContainer } from '@ant-design/pro-components';
import { Button, message, Space } from 'antd';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export default () => {
  return (
    <PageContainer>
    <Space>

      <ModalForm
        title="新建表单"
        trigger={<Button type="primary">自定义文字</Button>}
        submitter={{
          searchConfig: {
            submitText: '确认',
            resetText: '取消',
          },
        }}
        onFinish={async (values) => {
          await waitTime(2000);
          console.log(values);
          message.success('提交成功');
          return true;
        }}
      >
        <ProFormText
          width="md"
          name="name"
          label="签约客户名称"
          tooltip="最长为 24 位"
          placeholder="请输入名称"
        />

        <ProFormText width="md" name="company" label="我方公司名称" placeholder="请输入名称" />
      </ModalForm>

    </Space>
    </PageContainer>
  );
};