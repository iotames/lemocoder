import webserverConf from "../../config/webserver";
import type { ProFormInstance } from '@ant-design/pro-components';
import {
  ProForm,
  ProFormText,
  ProFormUploadDragger,
} from '@ant-design/pro-components';
import { message, Col, Row } from 'antd';
import { useRef } from 'react';

export default () => {
  const formRef = useRef<
    ProFormInstance<{
      title: string;
      sheet_name?: string;
      url_title?: string;
      uploadfile?: string;
    }>
  >();
  const uploadUrl = webserverConf.baseUrl + "/api/public/upload"
  return (
    <Row>
      <Col span={8}>

        <ProForm<{
          title: string;
          sheet_name?: string;
          url_title?: string;
          uploadfile?: string;
        }>
          onFinish={async (values) => {
            console.log("--------------values----", values);
            const val1 = await formRef.current?.validateFields();
            console.log('------validateFields:', val1);
            const val2 = await formRef.current?.validateFieldsReturnFormatValue?.();
            console.log('------validateFieldsReturnFormatValue:', val2);
            message.success('提交成功');
          }}
          formRef={formRef}
          dateFormatter={(value, valueType) => {
            console.log('---->', value, valueType);
            return value.format('YYYY/MM/DD HH:mm:ss');
          }}
          // request={async () => {
          //   return {
          //     title: '蚂蚁设计有限公司',
          //     sheet_name: 'chapter',
          //   };
          // }}
          autoFocusFirstInput
        >

    <ProFormText width="md" name="title" required label="标题" />
    <ProFormText width="md" name="sheet_name" initialValue="Sheet1" required label="Excel表名" />
    <ProFormText width="md" name="url_title" initialValue="URL" required label="Excel链接栏标题" />
    <ProFormText name="uploadfile" hidden />
    <ProFormUploadDragger width="md" label="Dragger" name="dragger" action={uploadUrl} onChange={(info)=>{
      // const resp = JSON.parse(info.file.response)
      if (info.file.response != undefined){
        const resp = info.file.response
        console.log(resp)
        console.log(resp.Data.Url)
        console.log(resp.Code)
        formRef.current?.setFieldsValue({uploadfile:resp.Data.Url})
      }
    }} />

        </ProForm>

      </Col>
    </Row>



  );
};