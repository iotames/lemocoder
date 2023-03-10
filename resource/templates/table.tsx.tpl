import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { PageContainer, ProTable, ModalForm, ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  ProFormDigit,
  ProFormInstance,
} from '@ant-design/pro-components';
import { Button, Space, message } from 'antd';
import { useRef, useState } from 'react';
import {postMsg, getTableData, postByBtn} from "@/services/api"
import { history } from 'umi';

type <%{.ItemDataTypeName}%> = {
    <%{range .Items}%> <%{.DataName}%>: <%{.DataType | getDataTypeForJS}%>; <%{end}%>
//  labels: {name: string; color: string;}[]; number
};

export default () => {
  const actionRef = useRef<ActionType>();
  const itemFormRef = useRef<ProFormInstance<<%{.ItemDataTypeName}%>>>();
  const [rowRecord, setRowRecord] = useState<<%{.ItemDataTypeName}%>>();
  <%{range .ItemForms}%>const [modal<%{.Key}%>Visit, setModal<%{.Key}%>Visit] = useState(false);
  <%{end}%>
  // const [editableKeys, setEditableRowKeys] = useState<React.Key[]>(() => []);

  <%{range .ToolBarForms}%>
  const <%{.Key}%> = (<ModalForm
    title="<%{.Form.Title}%>"
    trigger={<Button type="<%{.Button.Type}%>"><%{.Button.Title}%></Button>}
    onFinish={async (values) => {
      console.log(values);
      const resp = await postMsg("<%{.Form.SubmitUrl}%>", values)
      if (resp.Code == 200) {
        actionRef.current?.reload()
        return true;
      }
      return false;
    }}
  >
  <%{ .Form.FormFields | getFormFieldsHtml }%>
  </ModalForm>)
  <%{end}%>

  <%{range .ItemForms}%>
  const <%{.Key}%> = (<ModalForm
    title="<%{.Form.Title}%>"
    formRef={itemFormRef}
    open={modal<%{.Key}%>Visit}
    initialValues={rowRecord}
    onOpenChange={setModal<%{.Key}%>Visit}
    onFinish={async (values) => {
      console.log(values);
      const resp = await postMsg("<%{.Form.SubmitUrl}%>", values)
      if (resp.Code == 200) {
        actionRef.current?.reload()
        return true;
      }
      return false;
    }}
  >
  <%{ .Form.FormFields | getFormFieldsHtml }%>
  </ModalForm>)
  <%{end}%>

  // ???????????? ProTable ??????????????????????????? Form???????????????????????????
  // Form ??????????????? valueType ????????????????????????,??????????????????????????????????????????
  // valueType ??? ProComponents ????????????ProComponents ????????? valueType ?????????????????????????????????
  // https://procomponents.ant.design/components/schema#valuetype
  const columns: ProColumns<<%{.ItemDataTypeName}%>>[] = [
      <%{range .Items}%>
      {
        title: "<%{.Title}%>",
        <%{if ne .DataName ""}%>dataIndex: "<%{.DataName}%>",<%{end}%>
        <%{- if not .Editable}%>editable: <%{.Editable}%>,<%{end}%>
        <%{- if .Copyable}%>copyable: <%{.Copyable}%>,<%{end}%>
        <%{- if ne .ValueType "" }%>valueType: "<%{.ValueType}%>",<%{end}%>
        <%{- if .Ellipsis}%>ellipsis: <%{.Ellipsis}%>,<%{end}%>
        <%{- if ne .ColSize 0.0}%>colSize: <%{.ColSize}%>,<%{end}%>
        <%{- if ne .Width 0}%>width: <%{.Width}%>,<%{end}%>
        <%{- if gt .Order 0}%>order: <%{.Order}%>,// number<%{end}%>
        <%{- if .Sorter }%>sorter: <%{.Sorter}%>,// boolean<%{end}%>
        <%{- if not .Search}%>search: <%{.Search}%>,<%{end}%>
        <%{- if .HideInSearch}%>hideInSearch: <%{.HideInSearch}%>,<%{end}%>
        <%{- if .HideInTable}%>hideInTable: <%{.HideInTable}%>,<%{end}%>
        // filters,
        // search: false // undefined | false | { transform: (value: any) => any }
        // renderText, (text: any,record: T,index: number,action: UseFetchDataAction<T>) => string
      },
      <%{end}%>

  <%{if gt (len .ItemOptions) 0}%>
    {
      title: '??????',
      valueType: 'option',
      key: 'option',
      render: (text, record, _, action) => {
      return [<%{range .ItemOptions}%>
      <%{if eq .Type "edit"}%><a key="<%{.Key}%>" onClick={() => {action?.startEditable?.(record.<%{$.RowKey}%>);}}><%{.Title}%></a>,<%{end}%>
      <%{if eq .Type "action"}%><Button key="<%{.Key}%>" type='primary'  onClick={async (e)=>{
        const resp = await postByBtn(e, "<%{.Url}%>", record);
        if (resp.Code == 200) {
          actionRef.current?.reload()
          return true;
        }
        }} ><%{.Title}%></Button>,<%{end}%>
      <%{if eq .Type "form"}%><Button key="<%{.Key}%>" type="primary" onClick={() => {setModal<%{.Key}%>Visit(true);setRowRecord(record);itemFormRef.current?.setFieldsValue(record)}}><%{.Title}%></Button>,<%{end}%>
      <%{if eq .Type "redirect"}%><Button key="<%{.Key}%>" type='primary' onClick={(e)=>{ history.push("<%{.Url}%>"); }}><%{.Title}%></Button>,<%{end}%>
      <%{end}%>]},
    },
  <%{end}%>
  ];

  const pageSize = 10

  return (
    <PageContainer>
    <%{range .ItemForms}%>{<%{.Key}%>}<%{end}%>
    <ProTable<<%{.ItemDataTypeName}%>>
      headerTitle="<%{.Title}%>" 
      <%{- if not .Searchable }%> search={false} <%{end}%>
      columns={columns}
      rowSelection={{}}
      tableAlertRender={({ selectedRowKeys, selectedRows, onCleanSelected }) => (<Space><span>?????? {selectedRowKeys.length} ???<a onClick={onCleanSelected}>??????</a></span></Space>)}
      tableAlertOptionRender={({ selectedRowKeys, selectedRows, onCleanSelected }) => {
        return (<Space><%{range .BatchOptButtons}%><Button onClick={async (e)=>{
          const resp = await postByBtn(e, "<%{.Url}%>", {items:selectedRows});
          if (resp.Code == 200) {
            onCleanSelected()
            actionRef.current?.reload()
            return true;
          }
        }}><%{.Title}%></Button><%{end}%></Space>);
      }}
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
          await postMsg("<%{.ItemUpdateUrl}%>", update)
        },
        onDelete: async (k, row) => {
          await postMsg("<%{.ItemDeleteUrl}%>", row) // url must begin with /
        }
      }}
      columnsState={{
        persistenceKey: 'pro-table-singe-demos',
        persistenceType: 'localStorage',
        onChange(value) {
          console.log('value: ', value);
        },
      }}
      rowKey="<%{.RowKey}%>"
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
        // ??????????????? transform????????????????????????????????????????????????????????????
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
        pageSize: pageSize,
        onChange: (page) => console.log(page),
      }}
      dateFormatter="string"
      toolBarRender={() => [<%{range .ToolBarForms}%><%{.Key}%>,<%{end}%>]}
    />
    </PageContainer>
  );
};