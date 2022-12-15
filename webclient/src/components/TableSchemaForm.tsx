import {ProFormSwitch, ProFormInstance, ProFormDigit, ProFormList, ProFormText, ProFormSelect, ProForm, StepsForm, ProFormGroup, ModalForm} from '@ant-design/pro-components';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import type { FormInstance, StepsProps } from 'antd';
import TableSchema from '@/pages/TableSchema';
import { useRef, useState } from 'react';

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

export type TableSchema = {
  ID: string;
  PageID: string;
  Name: string;
  Title: string;
  Remark: string;
  StructSchema: StructSchema;
}

// const stepsFormRender = (dom: React.ReactNode, submitter: React.ReactNode) => {
//   return (
//     <Modal
//       title="构建数据表格"
//       width={900}
//       onOk={() => setTableFormVisit(false)}
//       onCancel={() => setTableFormVisit(false)}
//       visible={tableFormVisit}
//       footer={submitter}
//       destroyOnClose
//     >
//       {dom}
//     </Modal>
//   );
// }

// {onUploaded:(resp:any)=>void, name: string, action: string, label: string}
export const TableSchemaForm = (props:{
  postUrl: string,
  tableSchema: TableSchema,
  formRef?: React.MutableRefObject<FormInstance<any> | undefined | null>,
  setModalVisit?: React.Dispatch<React.SetStateAction<boolean>>,
  stepsFormRender?: (from: React.ReactNode, submitter: React.ReactNode) => React.ReactNode
}) => {
  const tableSchema = props.tableSchema
  // const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const formRef = props.formRef
  formRef?.current?.setFieldsValue(tableSchema)
  const setModalVisit = props.setModalVisit
  const stepsFormRender = props.stepsFormRender
  const postUrl = props.postUrl

  return (
    <>
      <StepsForm
        onFinish={async (values) => {
        const resp = await postMsg(postUrl, values)
        if (resp.Code == 200) {
          if (setModalVisit != undefined){
            setModalVisit(false)
          }
          return true;
        }
        return false;
      }}
      onCurrentChange={(n)=>{console.log(n)}}
      formRef={formRef}
      stepsFormRender={stepsFormRender}
      >
        <StepsForm.StepForm name="base" title="基础设置" initialValues={tableSchema}>
            <ProFormText name="PageID" hidden />
            <ProFormText name={["StructSchema", "RowKey"]} hidden />
            <ProFormText name={["StructSchema", "ItemDataTypeName"]} label="数据结构名" placeholder="ProductItem" rules={[{ required: true }]} /> 
            <ProFormText name={["StructSchema", "ItemsDataUrl"]}  label="数据源" placeholder="/api/table/demodata" rules={[{ required: true }]} />
            <ProFormText name={["StructSchema", "ItemUpdateUrl"]}  label="更新地址" placeholder="/api/demo/post" />
            <ProFormText name={["StructSchema", "ItemDeleteUrl"]} label="删除地址" placeholder="/api/demo/post" /> 
        </StepsForm.StepForm>

        <StepsForm.StepForm name="items" title="数据字段" initialValues={tableSchema}>
          <ProFormList name={["StructSchema", "Items"]} creatorButtonProps={{creatorButtonText: '添加数据字段'}}>
            <ProFormGroup>
            <ProFormSelect name="DataType" label="字段类型" initialValue="string" options={[
                {value:"string", label:"字符串(string)"},
                {value:"int", label:"整型(int)"},
                {value:"float", label:"浮点型(float)"},
                {value:"text", label:"文本(text)"},
                {value:"smallint", label:"小整型(smallint)"},
                {value:"bigint", label:"长整型(bigint)"},
                ]} rules={[{ required: true }]} />
              <ProFormSelect name="ValueType" label="值的类型" tooltip="会生成不同的渲染器" initialValue="text" options={[
                {value:"text", label:"纯文本"}
                ]} rules={[{ required: true }]} />
              <ProFormText name="DataName" placeholder="字段名" label="字段名" rules={[{ required: true }]} width={120} tooltip="英文字母" /> 
              <ProFormText name="Title" placeholder="字段标题" label="字段标题" rules={[{ required: true }]} width={120} tooltip="中英文均可" />
              <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" placeholder="像素宽度" width={90} tooltip="整数(默认0).范围:0~300" />
              <ProFormSwitch name="Editable" label="可编辑" />
{/* 
              <ProFormSwitch name="Sorter" label="允许排序" />
              
              <ProFormSwitch name="Copyable" label="可复制" />
              <ProFormSwitch name="Ellipsis" label="省略过长内容" />
              <ProFormDigit label="权重" tooltip="查询表单中的权重，权重大排序靠前" min={0} max={300} name="Order" placeholder={"请输入权重"} />
 */}

            </ProFormGroup>

            {/* <ProFormGroup></ProFormGroup> */}

          </ProFormList>
        </StepsForm.StepForm>

        <StepsForm.StepForm name="options" title="行数据操作" initialValues={tableSchema?.StructSchema}>
        <ProFormList name="ItemOptions" creatorButtonProps={{creatorButtonText: '添加行数据操作项'}}>
          <ProFormGroup>
            <ProFormSelect name="Type" label="操作类型" initialValue="action" options={[
                  {value:"action", label:"POST数据"},
                  {value:"redirect", label:"路由跳转"},
                  {value:"form", label:"表单提交"},
                  {value:"edit", label:"快捷编辑"},
                  ]} rules={[{ required: true }]} />
            <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} width={90} />
            <ProFormText name="Key" label="操作名" rules={[{ required: true }]} width={90} tooltip="英文" />
            <ProFormText name="Url" label="地址" />

          </ProFormGroup>

        </ProFormList>
        </StepsForm.StepForm>
        <StepsForm.StepForm name="batchOpts" title="批量数据操作" initialValues={tableSchema?.StructSchema}>
        <ProFormList name="BatchOptButtons" creatorButtonProps={{creatorButtonText: '添加批量数据操作项'}}>
          <ProFormGroup>
              <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} />
              <ProFormText name="Url" label="API地址" rules={[{ required: true }]} />
          </ProFormGroup>
        </ProFormList>
        </StepsForm.StepForm>

      </StepsForm>
    </>
  );
}

export const NewDataTableForm = (props:{
  // tableData: TableSchema,
  formRef: React.MutableRefObject<FormInstance<any> | undefined | null>,
  modalVisit: boolean,
  setModalVisit: React.Dispatch<React.SetStateAction<boolean>>,
}) => {
  const formRef = props.formRef
  // const formRef = useRef<ProFormInstance<TableSchema>>();
  // const tableData = props.tableData
  // formRef.current?.setFieldsValue(tableData)
  const postUrl = "/api/coder/table/add"
  const modalVisit = props.modalVisit
  const setModalVisit = props.setModalVisit
  return (
    <>
    <ModalForm 
      title="新建数据表格"
      width={900}
      visible={modalVisit}
      onVisibleChange={setModalVisit}
      formRef={formRef}
      onFinish={async (values) => {
        const resp = await postMsg(postUrl, values)
        if (resp.Code == 200) {
          setModalVisit(false)
          return true;
        }
        return false;
      }}
    >
      <ProFormText name="PageID" hidden />
      <ProFormText name={["StructSchema", "ItemDataTypeName"]} label="数据结构名" tooltip="例: Product, ProductReview" placeholder="例: Product, ProductReview" rules={[{ required: true }]} />


      <ProFormList name={["StructSchema", "Items"]} creatorButtonProps={{creatorButtonText: '添加数据字段'}}>
        <ProFormGroup>
          <ProFormSelect name="DataType" label="字段类型" initialValue="string" options={[
                {value:"string", label:"字符串(string)"},
                {value:"int", label:"整型(int)"},
                {value:"float", label:"浮点型(float)"},
                {value:"text", label:"文本(text)"},
                {value:"smallint", label:"小整型(smallint)"},
                {value:"bigint", label:"长整型(bigint)"},
              ]} rules={[{ required: true }]} />
            <ProFormSelect name="ValueType" label="值的类型" tooltip="会生成不同的渲染器" initialValue="text" options={[
              {value:"text", label:"纯文本"}
              ]} rules={[{ required: true }]} />
            <ProFormText name="DataName" placeholder="字段名" label="字段名" rules={[{ required: true }]} width={120} tooltip="英文字母" /> 
            <ProFormText name="Title" placeholder="字段标题" label="字段标题" rules={[{ required: true }]} width={120} tooltip="中英文均可" />
            <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" placeholder="像素宽度" width={90} tooltip="整数(默认0).范围:0~300" />
            <ProFormSwitch name="Editable" label="可编辑" initialValue={true} />
        </ProFormGroup>
      </ProFormList>

    </ModalForm>
    </>
  );
}