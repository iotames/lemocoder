import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { PageContainer, ProTable, ModalForm, ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-components';
import { Button, message } from 'antd';
import { useRef } from 'react';
import {post, getTableData} from "@/services/api"
import { history } from 'umi';

type <%{.ItemDataTypeName}%> = {
    <%{range .Items}%> <%{.DataName}%>: <%{.DataType}%>; <%{end}%>
//  labels: {name: string; color: string;}[]; number
};

const columns: ProColumns<<%{.ItemDataTypeName}%>>[] = [
    <%{range .Items}%>
    {
      title: "<%{.Title}%>",
     <%{if ne .DataName ""}%>dataIndex: "<%{.DataName}%>",<%{end}%>
     <%{if not .Editable}%>editable: <%{.Editable}%>,<%{end}%>
     <%{if .Copyable}%>copyable: <%{.Copyable}%>,<%{end}%>
     <%{if ne .ValueType "" }%>valueType: "<%{.ValueType}%>",<%{end}%>
     <%{if .Ellipsis}%>ellipsis: <%{.Ellipsis}%>,<%{end}%>
     <%{if ne .ColSize 0.0}%>colSize: <%{.ColSize}%>,<%{end}%>
     <%{if gt .Order 0}%>order: <%{.Order}%>,// number<%{end}%>
     <%{if .Sorter }%>sorter: <%{.Sorter}%>,// boolean<%{end}%>
     <%{if not .Search}%>search: <%{.Search}%>,<%{end}%>// search: { transform: (value: any) => any }
     <%{if .HideInSearch}%>hideInSearch: <%{.HideInSearch}%>,<%{end}%>
     <%{if .HideInTable}%>hideInTable: <%{.HideInTable}%>,<%{end}%>
      // filters,
      // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
    },
    <%{end}%>

<%{if gt (len .ItemOptions) 0}%>
  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    render: (text, record, _, action) => {
    return [<%{range .ItemOptions}%>
    <%{if eq .Type "edit"}%><Button key="<%{.Key}%>" type="primary" onClick={() => {action?.startEditable?.(record.id);}}><%{.Title}%></Button>,<%{end}%>
    <%{if eq .Type "action"}%><Button key="<%{.Key}%>" type='primary'  onClick={async (e)=>{
        const btn = e.currentTarget
        btn.setAttribute("disabled", "true");
        const resp = await post("<%{.Url}%>", record)
        if (resp.Code == 200) {
          message.success(resp.Msg);
        }else{
          message.error(resp.Msg);
        }
        btn.removeAttribute("disabled")
      }} ><%{.Title}%></Button>,<%{end}%>
    <%{if eq .Type "redirect"}%><Button key="<%{.Key}%>" type='primary' onClick={(e)=>{ history.push("<%{.Url}%>"); }}><%{.Title}%></Button>,<%{end}%>
    <%{end}%>]},
  },
<%{end}%>

];

<%{range .ModalForms}%>

const <%{.Key}%> = (<ModalForm
  title="<%{.Form.Title}%>"
  trigger={<Button type="<%{.Button.Type}%>"><%{.Button.Title}%></Button>}
  submitter={{searchConfig: {submitText: '确认',resetText: '取消',},}}

  onFinish={async (values) => {
    console.log(values);
    const resp = await post("<%{.Form.SubmitUrl}%>", values)
    if (resp.Code == 200) {
      message.success(resp.Msg);
    }else{
      message.error(resp.Msg);
    }
    return true;
  }}
>

<%{range .Form.FormFields}%>
<%{if eq (len .Group) 0}%><<%{.Component}%> name="<%{.Name}%>" label="<%{.Label}%>" <%{if ne .Width ""}%>width="<%{.Width}%>"<%{end}%> <%{if eq .Component "ProFormSelect"}%>request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]}<%{end}%> placeholder="<%{.Placeholder}%>" /><%{else}%>
<ProForm.Group><%{range .Group}%>
<<%{.Component}%> name="<%{.Name}%>" label="<%{.Label}%>" <%{if ne .Width ""}%>width="<%{.Width}%>"<%{end}%> <%{if eq .Component "ProFormSelect"}%>request={async()=>[{value:"value1", label:"label1"},{value:"value2", label:"label2"}]}<%{end}%> placeholder="<%{.Placeholder}%>" /><%{end}%>
</ProForm.Group>
<%{end}%>
<%{end}%>
</ModalForm>)
<%{end}%>

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
        params.sort = sort
        const resp = await getTableData<<%{.ItemDataTypeName}%>>("<%{.ItemsDataUrl}%>", params)
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
      toolBarRender={() => [<%{range .ModalForms}%><%{.Key}%>,<%{end}%>]}
    />
    </PageContainer>
  );
};