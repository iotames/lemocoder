import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable, ModalForm, ProFormText, PageContainer } from '@ant-design/pro-components';
import { Button, message } from 'antd';
import { useRef, useState } from 'react';
import {post, getTableData} from "@/services/api"

type GithubIssueItem = {
  url: string;
  id: number;
  number: number;
  title: string;
  labels: {name: string;color: string;}[];
  state: string;
  comments: number;
  created_at: string;
  updated_at: string;
  closed_at?: string;
};


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
    const resp = await post("/api/demo/post", values)
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
  const [modalVisit, setModalVisit] = useState(false);
  const [rowRecord, setRowRecord] = useState<GithubIssueItem>();
  const actionRef = useRef<ActionType>();
  // const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() => []);

  
const columns: ProColumns<GithubIssueItem>[] = [
  {
    title: 'ID',
    dataIndex: 'id',
    colSize:0.7,
    editable:false,
    copyable: true,
  },
  {
    title: '标题',
    dataIndex: 'title',
    copyable: true,
    ellipsis: true,
    colSize: 1,
    tip: '标题过长会自动收缩',
    formItemProps: {
      rules: [
        {
          required: true,
          message: '此项为必填项',
        },
      ],
    },
  },
  {
    title: '创建时间',
    editable:false,
    key: 'showTime',
    dataIndex: 'created_at',
    valueType: 'dateTime',
    sorter: true,
    hideInSearch: true,
  },

  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    render: (text, record, _, action) => {
    return [
      <Button key="edit" type="primary" onClick={() => {action?.startEditable?.(record.id);}}>行编辑</Button>,
      <Button key="form1" type="primary" onClick={() => {
        setRowRecord(record)
        setModalVisit(true)
      }}>编辑</Button>,
      <Button key="bttt" type='primary'  onClick={(e)=>{
        console.log(record.id, record.title)
        e.currentTarget.setAttribute("disabled", "true");
      }} >Hello</Button>,
      
      <a href={record.url} target="_blank" onClick={()=>{
        console.log("查看看看", record.id, record.title)
        message.success("hello wordd"+record.id)
        }} rel="noopener noreferrer" key="view">
        查看
      </a>,
    ]},
  },
];


  const editForm = (<ModalForm
    title="编辑表单"
    initialValues={rowRecord}
    visible={modalVisit}
    onVisibleChange={setModalVisit}
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
    <ProFormText name="id" hidden />
    <ProFormText width="md" name="title" label="标题" placeholder="请输入标题"/>
    <ProFormText width="md" name="company" label="我方公司名称" placeholder="请输入名称" />
  </ModalForm>)

  return (
    <PageContainer>
      {editForm}
    <ProTable<GithubIssueItem>
      columns={columns}
      actionRef={actionRef}
      cardBordered
      request={async (params = {}, sort, filter) => {
        console.log(sort, filter);
        params.page = params.current
        params.limit = params.pageSize
        params.sort = sort
        const resp = await getTableData<GithubIssueItem>("/api/table/demodata", params)
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
          await post("/api/demo/post", update)
        },
        onDelete: async (k, row) => {
          console.log("----delete-----", row)
          await post("/api/demo/post", row)
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