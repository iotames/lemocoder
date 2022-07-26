import config from "@/utils/config"
import { post } from '@/services/api';
import {
  ProFormInstance,
  ProFormSelect,
  ProForm,
  ProFormText,
  ProFormUploadDragger,
} from '@ant-design/pro-components';
import { message, Col, Row } from 'antd';
import { useRef } from 'react';

export default () => {
  const baseApiUrl = config.BaseApiUrl
  const formRef = useRef<
    ProFormInstance<{
      spider: string;
      title: string;
      sheet_name: string;
      url_title: string;
      uploadfile: string;
    }>
  >();
  const uploadUrl = baseApiUrl + "/api/local/upload"
  return (
    <Row>
      <Col span={8}>

        <ProForm<{
          spider: string;
          title: string;
          sheet_name: string;
          url_title: string;
          uploadfile: string;
        }>
          onFinish={async (values) => {
            console.log(values)
            const resp = await post("/api/local/excelspider", {sheet_name:values.sheet_name, spider:values.spider, title:values.title, url_title:values.url_title, uploadfile: values.uploadfile})
            if (resp.Code != 200){
              message.error(resp.Msg);
              return
            }
            message.success(resp.Msg);
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
          <ProFormSelect width="md" name="spider" rules={[{required: true, message: "爬虫不能为空"}]} label="爬虫" valueEnum={{alibaba:"alibaba.com", 1688:"1688.com"}} /> 
          <ProFormText width="md" name="title" rules={[{required: true, message: "标题不能为空"}]} label="标题" />
          <ProFormText width="md" name="sheet_name" initialValue="Sheet1" rules={[{required: true, message: "Excel表名不能为空"}]} label="Excel表名" />
          <ProFormText width="md" name="url_title" initialValue="URL" rules={[{required: true, message: "URL标题不能为空"}]} label="URL标题" />
          <ProFormText name="uploadfile" rules={[{required: true, message: "必须上传EXCEL文件"}]} hidden />
          <ProFormUploadDragger max={1} rules={[{required: true, message: "必须上传EXCEL文件"}]} width="md" label="上传EXCEL文件" name="dragger" action={uploadUrl} accept=".xlsx"  onChange={(info)=>{
            // const resp = JSON.parse(info.file.response)
            if (info.file.response != undefined){
              const resp = info.file.response
              console.log(resp)
              console.log(resp.Data.Url)
              console.log(resp.Code)
              console.log(resp.Data.ID)
              formRef.current?.setFieldsValue({uploadfile:resp.Data.ID})
            }
          }} />

        </ProForm>

      </Col>
    </Row>



  );
};