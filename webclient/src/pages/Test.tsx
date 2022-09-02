import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { PageContainer, ProTable, ModalForm, ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-components';
import { Button, message } from 'antd';
import { useRef } from 'react';
import {get, post} from "@/services/api"

type TestTableItem = {
     id: number;  title: string;  created_at: string; 
//  labels: {name: string; color: string;}[]; number
};

const columns: ProColumns<TestTableItem>[] = [
    
    {
      title: "ID",
      dataIndex: "id",
      editable: false, 
      copyable: true,
     
      ellipsis: false,
      colSize: 0.7,
      order: 0, // number
      sorter: false, // boolean
      search: false,  // search: { transform: (value: any) => any } 
      hideInSearch: false,
      hideInTable: false,
      // filters,
      // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
    },
    
    {
      title: "标题",
      dataIndex: "title",
     
      copyable: true,
     
      ellipsis: false,
      colSize: 1,
      order: 0, // number
      sorter: false, // boolean
 // search: { transform: (value: any) => any } 
      hideInSearch: false,
      hideInTable: false,
      // filters,
      // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
    },
    
    {
      title: "创建时间",
      dataIndex: "created_at",
      editable: false, 
      copyable: false,
      valueType: "dateTime", 
      ellipsis: false,
      colSize: 0,
      order: 0, // number
      sorter: true, // boolean
      search: false,  // search: { transform: (value: any) => any } 
      hideInSearch: false,
      hideInTable: false,
      // filters,
      // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
    },
    

  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    render: (text, record, _, action) => {
    return [
      <Button key="edit" type="primary" onClick={() => {action?.startEditable?.(record.id);}}>编辑</Button>,

      <Button key="bttt" type='primary'  onClick={(e)=>{
        console.log(record.id, record.title)
        e.currentTarget.setAttribute("disabled", "true");
      }} >Hello</Button>,
      
      <a href="#" target="_blank" onClick={()=>{
        console.log("查看看看", record.id, record.title)
        message.success("hello wordd"+record.id)
        }} rel="noopener noreferrer" key="view">
        查看
      </a>,
    ]},
  },
];



const create = (<ModalForm
  title="添加数据"
  trigger={<Button type="primary">创建</Button>}
  submitter={{searchConfig: {submitText: '确认',resetText: '取消',},}}

  onFinish={async (values) => {
    console.log(values);
    const resp = await post("/api/demo/post", values)
    if (resp.Code == 200) {
      message.success(resp.Msg);
    }else{
      message.error(resp.Msg);
    }
    return true;
  }}
>




<ProForm.Group>

<ProFormSelect
  name="useMode"
  label="生效方式"
  
  request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]}
  placeholder=""
/>

<ProFormDateRangePicker
  name="contractTime"
  label="有效期"
  
  
  placeholder=""
/>

</ProForm.Group>





<ProFormText
  name="name"
  label="客户名称"
  
  
  placeholder=""
/>





<ProFormText
  name="company"
  label="我方公司名称"
  
  
  placeholder=""
/>



</ModalForm>)


export default () => {
  const actionRef = useRef<ActionType>();
  // const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() => []);

  return (
    <PageContainer>
    <ProTable<TestTableItem>
      columns={columns}
      actionRef={actionRef}
      cardBordered
      request={async (params = {}, sort, filter) => {
        console.log(sort, filter);
        params.page = params.current;
        params.limit = params.pageSize;
        const resp = await get<{data: TestTableItem[]}>("/api/table/demodata", params)
        return {
          data: resp.data,
          success: true,
          total: 30
        }
      }}
      editable={{
        type: 'multiple',
        // editableKeys,
        // onChange: setEditableRowKeys,
        onSave: async (k, update, origin) => {
          console.log(update, origin);
          await post("/api/demo/post", update)
        },
        onDelete: async (k, row) => {
          await post("/api/demo/post", row) // url must begin with /
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