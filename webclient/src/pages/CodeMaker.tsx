import type { ActionType, ProFormInstance, ProColumns } from '@ant-design/pro-components';
import { ProCard, ProFormItem, ProFormSwitch, ProFormDigit, ProFormList, ProTable, ModalForm, ProFormText, ProFormSelect, TableDropdown, ProForm, StepsForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Button, Typography, message, Modal, InputNumber } from 'antd';
import { useRef, useState } from 'react';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
// import { history } from 'umi';

type PageItem = {
  ID: string;
  path: string;
  component: string;
  name: string;
  PageType: number;
  title: string;
  remark: string;
  created_at: string;
  updated_at: string;
};

type TableItemOptionSchema = {
  Key: string;
  Title: string;
  Type: string;
  Url: string;
}

type BatchOptButtonSchema = {
	Url: string;
  Title: string;
}


type FormSchema = {
	Title: string;
  SubmitUrl: string;
	FormFields: FormFieldSchema[]
}
type FormFieldSchema = {
	Group:     FormFieldSchema[]
  Component: string;
  Name: string;
  Label: string;
  Placeholder: string;
  Width: string;
}
type ButtonSchema = {
  Title: string;
  Type: string;
  Size: string;
}
type ModalFormSchema = {
	Key:    string;
	Button: ButtonSchema;
	Form:   FormSchema;
}

type TableItemSchema = {
  DataName: string;
  DataType: string; // number, string, 
  Title: string;
  ValueType: string; // 值的类型,会生成不同的渲染器. default(text). option, select, dateTime, dateRange
  Width: number;
  Order: number; // 查询表单中的权重，权重大排序靠前
  Editable: boolean;
  Copyable: boolean;
  Sorter: boolean;
  Search: boolean;
  HideInSearch: boolean;
  HideInTable: boolean;
}

type StructSchema = {
  ItemDataTypeName: string;
  ItemsDataUrl: string;
  ItemUpdateUrl: string;
  ItemDeleteUrl: string;
  ItemCreateUrl: string;
  RowKey: string;
  Items: TableItemSchema[];
  ItemOptions: TableItemOptionSchema[];
  ItemForms: ModalFormSchema[];
  ToolBarForms: ModalFormSchema[];
  BatchOptButtons: BatchOptButtonSchema[];
}
type TableSchema = {
  PageID: string;
  Name: string;
  Title: string;
  Remark: string;
  StructSchema: StructSchema;
}

export default () => {
  const [pageFormVisit, setPageFormVisit] = useState(false);
  const [tableFormVisit, setTableFormVisit] = useState(false);
  // const [rowRecord, setRowRecord] = useState<PageItem>();
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
  const actionRef = useRef<ActionType>();

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

  {
    title: '备注',
    dataIndex: 'Remark',
    ellipsis: true,
    // colSize: 1,
  },
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
    return [
      // <Button  type="primary" onClick={() => {}}></Button>,

      <Button key="editform1" type="primary" onClick={async () => {
        // setRowRecord(record);
        if(record.PageType == 0){
          console.log(record)
          console.log(record.ID)
          const resp = await get<{Code: number; Msg: string; Data: TableSchema}>("/api/user/table/get", {"page_id": record.ID})
          if (resp.Code == 500){
            await message.error(resp.Msg)
            return
          }
          let initDt = tableSchemaInit
          initDt.Name = record.name
          initDt.Title = record.title
          initDt.Remark = record.remark
          initDt.PageID = record.ID
          if (resp.Code == 200) {
            initDt = resp.Data
          }
          setTableSchema(initDt)
          tableFormRef.current?.setFieldsValue(initDt); 
          console.log(initDt)
          setTableFormVisit(true)
        }
        
      }}>构建</Button>,

      // <Button key="bttt" type='primary'  onClick={ async (e)=>{
      //   await postByBtn(e, "/api/demo/post", record)
      // }} >Hello</Button>,
      
      <TableDropdown
      key="actionGroup"
      onSelect={(akey:string) => {
        if (akey == "edit"){
          action?.startEditable?.(record.ID);
        }
        action?.reload()
      }}
      menus={[
        { key: "edit", name: '编辑' },
      ]}
    />,

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
      
      <CheckCard
        title={<Title level={5}>新建页面</Title>}
        description="创建网站后台页面. "
        onChange={()=>{setPageFormVisit(true)}}
      />

      {/* <CheckCard
        title={<Title level={5}>新建工程</Title>}
        description="创建Web后台项目. 包含前端和后端源码"
        onChange={()=>{history.push("/welcome")}}
      /> */}

      <StepsForm  
        onFinish={async (values) => {
        const resp = await postMsg("/api/user/table/add", values)
        if (resp.Code == 200) {
          return true;
        }
        return false;
      }}
      formRef={tableFormRef}
      stepsFormRender={(dom, submitter) => {
        return (
          <Modal
            title="构建数据表格"
            width={600}
            onCancel={() => setTableFormVisit(false)}
            visible={tableFormVisit}
            footer={submitter}
            destroyOnClose
          >
            {dom}
          </Modal>
        );
      }}
      >
  
        <StepsForm.StepForm name="base" title="基础设置" initialValues={tableSchema?.StructSchema} >
            <ProFormText name="PageID" hidden initialValue={tableSchema?.PageID} />
            <ProFormText name="RowKey" hidden initialValue={tableSchema?.StructSchema.RowKey}  />
            <ProFormText name="ItemDataTypeName" label="数据结构名" placeholder="ProductItem" rules={[{ required: true }]} /> 
            <ProFormText name="ItemsDataUrl" label="数据源" placeholder="/api/table/demodata" rules={[{ required: true }]} />
            <ProFormText name="ItemUpdateUrl" label="更新地址" placeholder="/api/demo/post" />
            <ProFormText name="ItemDeleteUrl" label="删除地址" placeholder="/api/demo/post" /> 
{/*           
          <ProFormText name="Name" label="名称" placeholder="请输入名称" initialValue={tableSchema?.Name} />
          <ProFormText name="Title" label="标题"  initialValue={tableSchema?.Title} />
          <ProFormText name="Remark" label="备注" initialValue={tableSchema?.Remark} />
*/}
        </StepsForm.StepForm>

        <StepsForm.StepForm name="items" title="数据字段" initialValues={tableSchema?.StructSchema}>
          <ProFormList name="Items" creatorButtonProps={{creatorButtonText: '添加数据字段'}}>
            <ProForm.Group>
              <ProFormText name="DataName" label="字段名" rules={[{ required: true }]} />
              <ProFormText name="Title" label="字段标题" rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
              <ProFormSelect name="DataType" label="字段类型" initialValue="string" options={[
                {value:"string", label:"字符串"},
                {value:"int", label:"整型"},
                {value:"float", label:"浮点型"}
                ]} rules={[{ required: true }]} />
              <ProFormSelect name="ValueType" label="值的类型" tooltip="会生成不同的渲染器" initialValue="text" options={[
                {value:"text", label:"纯文本"}
                ]} rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
              <ProFormSwitch name="Sorter" label="允许排序" />
              <ProFormSwitch name="Editable" label="可编辑" />
              <ProFormSwitch name="Copyable" label="可复制" />
              <ProFormSwitch name="Ellipsis" label="省略过长内容" />
            </ProForm.Group>
            <ProForm.Group>
              <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" />
              <ProFormDigit label="权重" tooltip="查询表单中的权重，权重大排序靠前" min={0} max={300} name="Order" />
            </ProForm.Group>
          </ProFormList>
        </StepsForm.StepForm>

        <StepsForm.StepForm name="options" title="行数据操作" initialValues={tableSchema?.StructSchema}>
        <ProFormList name="ItemOptions" creatorButtonProps={{creatorButtonText: '添加行数据操作项'}}>
          <ProForm.Group>
            <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} />
            <ProFormText name="Key" label="操作名" rules={[{ required: true }]} />
          </ProForm.Group>
          <ProForm.Group>
            <ProFormSelect name="Type" label="操作类型" initialValue="action" options={[
                {value:"edit", label:"快捷编辑"},
                {value:"action", label:"标记数据"},
                {value:"redirect", label:"路由跳转"},
                {value:"form", label:"表单提交"}
                ]} rules={[{ required: true }]} />
            <ProFormText name="Url" label="API地址" />
          </ProForm.Group>
        </ProFormList>
        </StepsForm.StepForm>
        <StepsForm.StepForm name="batchOpts" title="批量数据操作" initialValues={tableSchema?.StructSchema}>
        <ProFormList name="BatchOptButtons" creatorButtonProps={{creatorButtonText: '添加批量数据操作项'}}>
          <ProForm.Group>
              <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} />
              <ProFormText name="Url" label="API地址" rules={[{ required: true }]} />
          </ProForm.Group>
        </ProFormList>
        </StepsForm.StepForm>

      </StepsForm>

      <ModalForm
        title="新建页面"
        visible={pageFormVisit}
        // width={600}
        onFinish={async (values) => {
          const resp = await postMsg("/api/user/page/add", values)
          if (resp.Code == 200) {
            return true;
          }
          return false;
        }}
        onVisibleChange={setPageFormVisit}
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
        <ProForm.Group>
          {/* <ProFormText name="title"  label="页面标题(选填)" /> */}
          <ProFormText name="remark"  label="备注(选填)" />
        </ProForm.Group>

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
        const resp = await getTableData<PageItem>("/api/user/pages", params)
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
        onSave: async (k, update, origin) => {await post("/api/demo/post", update)},
        onDelete: async (k, row) => {await post("/api/demo/post", row)}
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