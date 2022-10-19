import type { ActionType, ProFormInstance, ProColumns } from '@ant-design/pro-components';
import { ProCard, ProFormItem, ProFormSwitch, ProFormDigit, ProFormList, ProTable, ModalForm, ProFormText, ProFormSelect, TableDropdown, ProForm, StepsForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Button, Typography, message, Modal, InputNumber, Input, Select } from 'antd';
import { useRef, useState, useEffect } from 'react';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import TableSchemaForm from "@/components/TableSchemaForm"
import type { TableSchema } from "@/components/TableSchemaForm"

import { history } from 'umi';

export default () => {
  const tableSchemaInit = {
    PageID:"0", Name:"", Title:"", Remark:"", 
    StructSchema:{
      ItemDataTypeName:"ProductItem", 
      ItemsDataUrl:"/api/table/demodata",
      ItemUpdateUrl:"/api/demo/post",
      ItemDeleteUrl:"/api/demo/post",
      RowKey:"ID"
    }}
  const [tableSchema, setTableSchema] = useState<TableSchema>();
  const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const refresh = async (pageID: string) => {
    const resp = await get<{Code: number; Msg: string; Data: TableSchema}>("/api/user/table/get", {"page_id": pageID})
  if (resp.Code == 500){
    await message.error(resp.Msg)
    return
  }
  let initDt = tableSchemaInit
  initDt.PageID = pageID
  if (resp.Code == 200) {
    initDt = resp.Data
  }
  setTableSchema(initDt)
  tableFormRef.current?.setFieldsValue(initDt); 
  console.log(initDt)
  }
  const queryArgs = history.location.query
  useEffect(() => {
    if (queryArgs.page_id != undefined) {
        // console.log(queryArgs.page_id)
        refresh(queryArgs.page_id)
        // (async ()=>{await refresh(queryArgs.page_id);})()
    }else{
        message.error("page_id 参数不正确")
    }
  }, [queryArgs]);

  return (
    <PageContainer>
        <TableSchemaForm tableSchema={tableSchema} formRef={tableFormRef} />
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