import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable, ModalForm, ProFormText, PageContainer } from '@ant-design/pro-components';
import { Button, message } from 'antd';
import { useRef } from 'react';
import {get, post} from "@/services/api"

type <%{.ItemDataTypeName}%> = {
    <%{range .Items}%> <%{.DataName}%>: <%{.DataType}%>; <%{end}%>
//  labels: {name: string; color: string;}[]; number
};

const columns: ProColumns<<%{.ItemDataTypeName}%>>[] = [
    <%{range .Items}%>
    {
      title: "<%{.Title}%>",
      dataIndex: "<%{.DataName}%>",
     <%{if not .Editable}%> editable: <%{.Editable}%>, <%{end}%>
      copyable: <%{.Copyable}%>,
     <%{if ne .ValueType "" }%> valueType: "<%{.ValueType}%>", <%{end}%>
      ellipsis: <%{.Ellipsis}%>,
      colSize: <%{.ColSize}%>,
      order: <%{.Order}%>, // number
      sorter: <%{.Sorter}%>, // boolean
<%{if not .Search}%>      search: <%{.Search}%>, <%{end}%> // search: { transform: (value: any) => any } 
      hideInSearch: <%{.HideInSearch}%>,
      hideInTable: <%{.HideInTable}%>,
      // filters,
      // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
    },
    <%{end}%>

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


const createBtn = (<ModalForm
  title="新建表单"
  trigger={<Button type="primary">创建</Button>}
  submitter={{
    searchConfig: {
      submitText: '确认',
      resetText: '取消',
    },
  }}

  onFinish={async (values) => {
    console.log(values);
    const resp = await post("<%{.ItemCreateUrl}%>", values)
    if (resp.Code == 200) {
      message.success(resp.Msg);
    }else{
      message.error(resp.Msg);
    }
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
</ModalForm>)

export default () => {
  const actionRef = useRef<ActionType>();
  // const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() => []);

  return (
    <PageContainer>
    <ProTable<<%{.ItemDataTypeName}%>>
      columns={columns}
      actionRef={actionRef}
      cardBordered
      request={async (params = {}, sort, filter) => {
        console.log(sort, filter);
        params.page = params.current;
        params.limit = params.pageSize;
        const resp = await get<{data: <%{.ItemDataTypeName}%>[]}>("<%{.ItemsDataUrl}%>", params)
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
          await post("<%{.ItemUpdateUrl}%>", update)
        },
        onDelete: async (k, row) => {
          await post("<%{.ItemDeleteUrl}%>", row) // url must begin with /
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
      toolBarRender={() => [
        createBtn,
        // <Button type="primary" onClick={()=>{
        //   setTableItemForm({visible: true, action:"/", title:"新建元素"})
        // }}>新建</Button>
      ]}
    />
    </PageContainer>
  );
};