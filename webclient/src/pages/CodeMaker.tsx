import type { ActionType, ProFormInstance, ProColumns } from '@ant-design/pro-components';
import { ProTable, ModalForm, ProFormText, ProFormSelect, TableDropdown, ProForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Button, Typography, message, Modal, InputNumber, Input, Select } from 'antd';
import { useRef, useState } from 'react';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import {NewDataTableForm} from "@/components/TableSchemaForm"
import type { TableSchema } from "@/components/TableSchemaForm"
import { EyeOutlined, PlusOutlined } from '@ant-design/icons';
import { history } from 'umi';

type PageItem = {
  ID: string;
  path: string;
  component: string;
  name: string;
  PageType: number;
  title: string;
  remark: string;
  State: number;
  created_at: string;
  updated_at: string;
};

export default () => {
  const [pageFormVisit, setPageFormVisit] = useState(false);
  const [tableFormVisit, setTableFormVisit] = useState(false);
  const [pageID, setPageID] = useState<string>("");
  const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const actionRef = useRef<ActionType>();

  const tableSchemaInit = {
    PageID:"0", Name:"", Title:"", Remark:"", 
    StructSchema:{
      RowKey:"ID"
    }}


const columns: ProColumns<PageItem>[] = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: 180,
    // colSize:0.7,
    editable:false,
    copyable: true,
    hideInSearch: true,
  },
  {
    title: '路径',
    dataIndex: 'Path',
    ellipsis: true,
    // colSize: 1,
    formItemProps: {rules: [{required: true,message: '此项为必填项',},],},
  },
  {
    title: '名称',
    dataIndex: 'Name',
    ellipsis: true,
    // colSize: 0.3,
    formItemProps: {rules: [{required: true,message: '此项为必填项',},],}
  },
  {
    title: '组件',
    dataIndex: 'Component',
    ellipsis: true,
    // colSize: 0.3,
    formItemProps: {rules: [{required: true,message: '此项为必填项',},],}
  },

  // {
  //   title: '备注',
  //   dataIndex: 'Remark',
  //   ellipsis: true,
  //   // colSize: 1,
  // },

  {
    title: '创建时间',
    editable:false,
    key: 'showTime',
    dataIndex: 'CreatedAt',
    valueType: 'dateTime',
    sorter: true,
    hideInSearch: true,
  },

  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    // width: 260,
    // fixed: 'right',
    render: (text, record, _, action) => {
      let btnCoin = (<PlusOutlined />)
      if (record.State > 0) {
        btnCoin = (<EyeOutlined />)
      }
      return [
        // <Button  type="primary" onClick={() => {}}></Button>,

        <Button key="editform1" type="primary" icon={btnCoin} onClick={async () => {
          // setRowRecord(record);
          if(record.PageType == 0){
            
            console.log(record)
            console.log(record.ID)
            const resp = await get<{Code: number; Msg: string; Data: TableSchema}>("/api/coder/table/get", {"page_id": record.ID})
            if (resp.Code == 500){
              await message.error(resp.Msg)
              return
            }
            let initDt = tableSchemaInit
            initDt.Name = record.name
            initDt.Title = record.title
            initDt.Remark = record.remark
            initDt.PageID = record.ID
            if (resp.Code == 404){
              // 添加数据
              setTableFormVisit(true) // 显示表单应该在设置表单数据之前，否则第一次弹层表单会无数据
              setPageID(initDt.PageID) // TODO 丢失 tableFormRef.current
              tableFormRef.current?.setFieldsValue(initDt);
              return
            }

            if (resp.Code == 200) {
              // 修改数据
              // initDt = resp.Data
              history.push("/tableschema?page_id="+record.ID)
              return
            }
            
          }
          
        }}></Button>,

        // <Button key="createcode" type="primary" onClick={async () => {await postMsg("/api/coder/table/createcode", {"PageID": record.ID})}}>创建代码</Button>,

        // <Button key="bttt" type='primary'  onClick={ async (e)=>{
        //   await postByBtn(e, "/api/demo/post", record)
        // }} >Hello</Button>,
        
      //   <TableDropdown
      //   key="actionGroup"
      //   onSelect={(akey:string) => {
      //     if (akey == "edit"){
      //       action?.startEditable?.(record.ID);
      //     }
      //     action?.reload()
      //   }}
      //   menus={[
      //     { key: "edit", name: '编辑' },
      //   ]}
      // />,
      <a onClick={()=>{action?.startEditable?.(record.ID);action?.reload()}}>编辑</a>

        // <a href={record.path} target="_blank" onClick={()=>{
        //   console.log("查看看看", record.id, record.title)
        //   message.success("hello wordd"+record.id)
        //   }} rel="noopener noreferrer" key="view">
        //   查看
        // </a>,
      ]},
  },
];

const {Title} = Typography;

  return (
    <PageContainer>
      <NewDataTableForm pageID={pageID} formRef={tableFormRef} setModalVisit={setTableFormVisit} modalVisit={tableFormVisit}  />
      <ModalForm
        title="新建页面"
        open={pageFormVisit}
        // width={600}
        onFinish={async (values) => {
          const resp = await postMsg("/api/coder/page/create", values)
          if (resp.Code == 200) {
            actionRef.current?.reload()
            return true;
          }
          return false;
        }}
        onOpenChange={setPageFormVisit}
      >
        <ProForm.Group>
        <ProFormSelect initialValue={0} width="sm" name="project_id" label="工程项目" required
        rules={[{ required: true, message: '此项必填' }]}
            options={[{value: 0, label: '本项目'}]}
          />
        <ProFormSelect initialValue={0} width="sm" name="page_type" label="页面类型" required
        rules={[{ required: true, message: '页面类型为必填项' }]}
            options={[{value: 0, label: '数据表格'}]}
          />

        </ProForm.Group>

        <ProForm.Group>
        <ProFormText width="sm" name="name" label="名称(菜单名)" rules={[{ required: true, message: '此项必填' }]} tooltip="中英文均可" placeholder="产品列表" />
        <ProFormText width="sm" name="path" label="路径" rules={[{ required: true, message: '此项必填' }]} placeholder="/products" />
        <ProFormText width="sm" name="component"  label="前端组件名" tooltip="前端组件文件名。英文, 大写字母开头. 例: Product 生成文件Product.tsx" rules={[{ required: true, message: '此项必填' }]} placeholder="Product" />
        </ProForm.Group>
{/* 
        <ProForm.Group>
          <ProFormText name="title"  label="页面标题(选填)" />
          <ProFormText name="remark"  label="备注(选填)" />
        </ProForm.Group>
 */}
      </ModalForm>


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

    <ProTable<PageItem>
      search={false}
      headerTitle="页面管理"
      columns={columns}
      actionRef={actionRef}
      // scroll={{ x: 1500 }}
      cardBordered
      request={async (params = {}, sort, filter) => {
        console.log(sort, filter);
        params.page = params.current
        params.limit = params.pageSize
        params.sort = sort
        const resp = await getTableData<PageItem>("/api/coder/pages", params)
        if (resp.Code != 200) {
          message.error(resp.Msg)
          return {success: false}
        }
        return {
          data: resp.Data.Items,
          success: true,
          total: resp.Data.Total
        }
      }}
      editable={{
        type: 'multiple',
        // editableKeys,
        // onChange: setEditableRowKeys,
        onSave: async (k, update, origin) => {await postMsg("/api/coder/page/update", update); actionRef.current?.reload},
        onDelete: async (k, row) => {await postMsg("/api/coder/page/delete", row); actionRef.current?.reload}
      }}
      columnsState={{
        persistenceKey: 'pro-table-singe-demos',
        persistenceType: 'localStorage',
        onChange(value) {
          console.log('value: ', value);
        },
      }}
      rowKey="ID"
      pagination={{
        showSizeChanger: true,
        pageSize: 10,
        onChange: (page) => console.log(page),
      }}
      dateFormatter="string"
      toolBarRender={() => [
        <Button type="primary" onClick={()=>{
          setPageFormVisit(true)
        }}>新建</Button>
      ]}
    />

    </PageContainer>
  );
};