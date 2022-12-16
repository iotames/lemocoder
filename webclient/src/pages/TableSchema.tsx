import { ProCard, ProFormItem, ProFormSwitch, ProFormGroup, ProFormInstance, ProFormDigit, ProFormList, ProTable, ModalForm, ProFormText, ProForm, StepsForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Row, Col, Button, Typography, message, Modal, InputNumber, Input, Select } from 'antd';
import { useRef, useState, useEffect } from 'react';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import type { TableSchema } from "@/components/TableSchemaForm"
import { TableItemFormFields, TableItemOptFormFields } from "@/components/Form"
import { PlayCircleOutlined } from '@ant-design/icons';

import { history } from 'umi';

export default () => {
  // const [tableSchema, setTableSchema] = useState<TableSchema>();
  const [pageState, setPageState] = useState<number>(0);
  const [pageID, setPageID] = useState<string>();
  const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const refresh = async (pageID: string) => {
    const pageResp = await get<{Code: number; Msg: string; Data:{State: number}}>("/api/coder/page/get", {"id": pageID})
    if (pageResp.Code == 500){
      await message.error(pageResp.Msg)
      return
    }
    if (pageResp.Code == 404) {
      await message.error(pageResp.Msg)
      return
    }
    setPageState(pageResp.Data.State)
    const resp = await get<{Code: number; Msg: string; Data: TableSchema}>("/api/coder/table/get", {"page_id": pageID})
    if (resp.Code == 500){
      await message.error(resp.Msg)
      return
    }
    if (resp.Code == 404) {
      await message.error("数据不存在")
      return
    }
    if (resp.Code == 200) {
      // setTableSchema(resp.Data)
      tableFormRef.current?.setFieldsValue(resp.Data); 
    }
  }
  const queryArgs = history.location.query
  useEffect(() => {
    if (queryArgs?.page_id != undefined) {
        // console.log(queryArgs.page_id)
        setPageID(queryArgs.page_id)
        refresh(queryArgs.page_id)
        // (async ()=>{await refresh(queryArgs.page_id);})()
    } else {
        message.error("page_id 参数不正确")
    }
  }, [queryArgs]);

  const TableSchemaUpdateForm = (
  <ProForm<TableSchema>
  submitter={{ searchConfig: { submitText: '修改'}, resetButtonProps:{style:{display:"none"}}}}
    onFinish={async(values) => {
      await postMsg("/api/coder/table/update", values)
    }}
    formRef={tableFormRef}
    // initialValues={tableSchema}
    >
      <ProFormText name="ID" hidden />
      <ProFormText name="PageID" hidden />
      <ProFormText name={["StructSchema", "RowKey"]} hidden />

      
      <ProForm.Group>
        <ProFormText width={120} name={["StructSchema", "ItemDataTypeName"]} label="数据结构名" rules={[{ required: true }]} />
        <ProFormText width={180} name={["StructSchema", "ItemsDataUrl"]} label="数据源" rules={[{ required: true }]} />
        <ProFormText width={180} name={["StructSchema", "ItemCreateUrl"]} label="新增数据地址" />
        <ProFormText width={180} name={["StructSchema", "ItemUpdateUrl"]} label="更新数据地址" />
        <ProFormText width={180} name={["StructSchema", "ItemDeleteUrl"]} label="删除数据地址" /> 
      </ProForm.Group>

      {/* <ProForm.Group label="数据字段"> */}
      <ProCard title="数据字段" style={{ marginBlockStart: 16 }} >
        <ProFormList 
        name={["StructSchema", "Items"]} 
        creatorButtonProps={{creatorButtonText: '添加数据字段'}}
        >
          
          <TableItemFormFields />
          
        </ProFormList>
      </ProCard>
      {/* </ProForm.Group> */}

      <ProCard title="行数据操作" style={{ marginBlockStart: 16 }}>
          <ProFormList name={["StructSchema", "ItemOptions"]} creatorButtonProps={{creatorButtonText: '添加行数据操作项'}}>
          <TableItemOptFormFields />
        </ProFormList>
      </ProCard>

      <ProCard title="批量数据操作" style={{ marginBlock: 16 }}>
        <ProFormList name={["StructSchema", "BatchOptButtons"]} creatorButtonProps={{creatorButtonText: '添加批量数据操作项'}}>
          <ProFormGroup>
              <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} />
              <ProFormText name="Url" label="API地址" rules={[{ required: true }]} />
          </ProFormGroup>
        </ProFormList>
      </ProCard>


  </ProForm>)

  const codeGen = (
    <Col span={12} style={{ marginBlockEnd: 16 }}>
    <Button type='primary' shape="default" icon={<PlayCircleOutlined />} onClick={async()=>{await postMsg("/api/coder/table/createcode", {"PageID": pageID})}}>生成代码</Button>
  </Col>
  );
  let topBtn = (<Col></Col>)
  if (pageState < 2){
    topBtn = codeGen
  }

  return (
    <PageContainer>

      <Row>
        {/* TODO  判断是否已生成 然后显示按钮  生成代码 */}{/* 代码回滚 */}
        {topBtn}
      </Row>
      
      <Row>
        <Col span={20}>{TableSchemaUpdateForm}</Col>
      </Row>
      
{/*     
    <ProCard tabs={{type: 'card'}} >
        <ProCard.TabPane key="tab1" tab="产品一">
          内容一
        </ProCard.TabPane>
        <ProCard.TabPane key="tab2" tab="产品二">
          内容二
        </ProCard.TabPane>
    </ProCard>
*/}
    </PageContainer>
  );
};