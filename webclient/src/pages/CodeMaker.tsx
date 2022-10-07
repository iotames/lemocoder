import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProCard, ProTable, ModalForm, ProFormText, ProFormSelect, ProForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Button, Typography, message } from 'antd';
import { useRef, useState } from 'react';
import {post, postMsg, getTableData, postByBtn} from "@/services/api"
import { history } from 'umi';

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
  const [pageFormVisit, setPageFormVisit] = useState(false);
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
    hideInSearch: true,
  },
  {
    title: '标题',
    dataIndex: 'title',
    copyable: true,
    ellipsis: true,
    colSize: 1,
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

const {Title} = Typography;


  return (
    <PageContainer>
      


 
        <CheckCard
          title={<Title level={5}>新增页面</Title>}
          description="创建网站后台页面. "
          onChange={()=>{setPageFormVisit(true)}}
        />
        <CheckCard
          title={<Title level={5}>新建工程</Title>}
          description="创建Web后台项目. 包含前端和后端源码"
          onChange={()=>{history.push("/welcome")}}
        />




<ModalForm
        title="新增页面"
        visible={pageFormVisit}
        // width={600}
        onFinish={async (values) => {
          const resp = await postMsg("/api/user/pages/add", values)
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
          <ProFormText name="title"  label="页面标题(选填)" />
          <ProFormText name="remark"  label="备注(选填)" />
        </ProForm.Group>
        {/* <ProFormText name="project" disabled label="项目名称" initialValue="xxxx项目" />
        <ProFormText width="xs" name="mangerName" disabled label="商务经理" initialValue="启途" /> */}

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
      pagination={{
        showSizeChanger: true,
        pageSize: 10,
        onChange: (page) => console.log(page),
      }}
      dateFormatter="string"
      headerTitle="表格管理"
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