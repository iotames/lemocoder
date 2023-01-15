import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { PageContainer, ProTable, ModalForm, ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
  ProFormInstance,
} from '@ant-design/pro-components';
import { Button, Space, message } from 'antd';
import { useRef, useState } from 'react';
import {postMsg, getTableData, postByBtn} from "@/services/api"
import { history } from 'umi';

type TestTableItem = {
     id: number;  title: string;  price: number;  inventory: number;  created_at: string; 
//  labels: {name: string; color: string;}[]; number
};



const create = (<ModalForm
  title="添加数据"
  trigger={<Button type="primary">创建</Button>}
  onFinish={async (values) => {
    console.log(values);
    await postMsg("/api/demo/post", values)
    return true;
  }}
>

<ProForm.Group><ProFormSelect name="useMode" label="生效方式" request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]} placeholder="" /><ProFormDateRangePicker name="contractTime" label="有效期" placeholder="" /></ProForm.Group><ProFormText name="name" label="客户名称" placeholder="" /><ProFormText name="company" label="我方公司名称" placeholder="" />

</ModalForm>)


export default () => {
  const actionRef = useRef<ActionType>();
  const itemFormRef = useRef<ProFormInstance<TestTableItem>>();
  const [rowRecord, setRowRecord] = useState<TestTableItem>();
  const [modaleditform1Visit, setModaleditform1Visit] = useState(false);
  
  // const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() => []);

  
  const editform1 = (<ModalForm
    title="编辑数据"
    formRef={itemFormRef}
    open={modaleditform1Visit}
    initialValues={rowRecord}
    onOpenChange={setModaleditform1Visit}
    onFinish={async (values) => {
      console.log(values);
      await postMsg("/api/demo/post", values)
      return true;
    }}
  >

  <ProFormText name="id" label="ID" placeholder="" /><ProFormText name="title" label="Title" placeholder="" />

  </ModalForm>)
  


  const columns: ProColumns<TestTableItem>[] = [
      
      {
        title: "ID",dataIndex: "id",editable: false,copyable: true,colSize: 0.7,
      
      
      search: false,// search: { transform: (value: any) => any }
      
      
        // filters,
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      
      {
        title: "标题",dataIndex: "title",copyable: true,colSize: 1,
      
      
      // search: { transform: (value: any) => any }
      
      
        // filters,
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      
      {
        title: "价格",dataIndex: "price",copyable: true,colSize: 1,
      
      
      // search: { transform: (value: any) => any }
      
      
        // filters,
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      
      {
        title: "库存",dataIndex: "inventory",copyable: true,colSize: 1,
      
      
      // search: { transform: (value: any) => any }
      
      
        // filters,
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      
      {
        title: "创建时间",dataIndex: "created_at",editable: false,valueType: "dateTime",
      
      sorter: true,// boolean
      search: false,// search: { transform: (value: any) => any }
      
      
        // filters,
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      

  
    {
      title: '操作',
      valueType: 'option',
      key: 'option',
      render: (text, record, _, action) => {
      return [
      <Button key="edit" type="primary" onClick={() => {action?.startEditable?.(record.id);}}>行编辑</Button>,
      
      
      
      
      
      
      <Button key="editform1" type="primary" onClick={() => {setRowRecord(record);itemFormRef.current?.setFieldsValue(record);setModaleditform1Visit(true)}}>表单编辑</Button>,
      
      
      
      <Button key="post1" type='primary'  onClick={async (e)=>{await postByBtn(e, "/api/demo/post", record);}} >标记</Button>,
      
      
      
      
      
      
      <Button key="ret" type='primary' onClick={(e)=>{ history.push("/welcome"); }}>跳转</Button>,
      ]},
    },
  

  ];


  return (
    <PageContainer>
    {editform1}
    <ProTable<TestTableItem>
      columns={columns}
      rowSelection={{}}
      tableAlertRender={({ selectedRowKeys, selectedRows, onCleanSelected }) => (<Space><span>已选 {selectedRowKeys.length} 项<a onClick={onCleanSelected}>取消</a></span></Space>)}
      tableAlertOptionRender={({ selectedRowKeys, selectedRows, onCleanSelected }) => {
        return (<Space><Button onClick={async (e)=>{await postByBtn(e, "/api/demo/post", {items:selectedRows})}}>批量操作A</Button><Button onClick={async (e)=>{await postByBtn(e, "/api/demo/post", {items:selectedRows})}}>批量操作B</Button></Space>);
      }}
      actionRef={actionRef}
      cardBordered
      request={async (params = {}, sort, filter) => {
        console.log(sort, filter);
        params.page = params.current;
        params.limit = params.pageSize;
        params.sort = sort
        const resp = await getTableData<TestTableItem>("/api/table/demodata", params)
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
        onSave: async (k, update, origin) => {
          console.log(update, origin);
          await postMsg("/api/demo/post", update)
        },
        onDelete: async (k, row) => {
          await postMsg("/api/demo/post", row) // url must begin with /
        }
      }}
      columnsState={{
        persistenceKey: 'pro-table-singe-demos',
        persistenceType: 'localStorage',
        onChange(value) {
          console.log('value: ', value);
        },
      }}
      rowKey="id"
      search={{
        labelWidth: 'auto',
        span: 6,
        defaultCollapsed: false,
      }}
      options={{
        setting: {
          listsHeight: 400,
        },
      }}
      form={{
        // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
        syncToUrl: (values, type) => {
          if (type === 'get') {
            return {
              ...values,
              created_at: [values.startTime, values.endTime],
            };
          }
          return values;
        },
      }}
      pagination={{
        showSizeChanger: true,
        pageSize: 10,
        onChange: (page) => console.log(page),
      }}
      dateFormatter="string"
      headerTitle="高级表格"
      toolBarRender={() => [create,]}
    />
    </PageContainer>
  );
};