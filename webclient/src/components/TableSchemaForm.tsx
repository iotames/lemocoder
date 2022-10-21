import {ProFormSwitch, ProFormDigit, ProFormList, ProFormText, ProFormSelect, ProForm, StepsForm} from '@ant-design/pro-components';
// import { Button, Typography, message, Modal, InputNumber, Input, Select } from 'antd';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import type { FormInstance, StepsProps } from 'antd';

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
  PageID: string;
  Name: string;
  Title: string;
  Remark: string;
  StructSchema: StructSchema;
}

// {onUploaded:(resp:any)=>void, name: string, action: string, label: string}
const TableSchemaForm = (props:{
  tableSchema: TableSchema, 
  formRef?: React.MutableRefObject<FormInstance<any> | undefined | null>,
  setModalVisit?: React.Dispatch<React.SetStateAction<boolean>>,
  stepsFormRender?: (from: React.ReactNode, submitter: React.ReactNode) => React.ReactNode
}) => {
  const tableSchema = props.tableSchema
  // const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const formRef = props.formRef
  const setModalVisit = props.setModalVisit
  const stepsFormRender = props.stepsFormRender
  return (
    <>
      <StepsForm
        onFinish={async (values) => {
        const resp = await postMsg("/api/coder/table/add", values)
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
        <StepsForm.StepForm name="base" title="基础设置" initialValues={tableSchema?.StructSchema} >
            <ProFormText name="PageID" hidden initialValue={tableSchema?.PageID} />
            <ProFormText name="RowKey" hidden initialValue={tableSchema?.StructSchema.RowKey}  />
            <ProFormText name="ItemDataTypeName" value={tableSchema?.StructSchema.ItemDataTypeName} label="数据结构名" placeholder="ProductItem" rules={[{ required: true }]} /> 
            <ProFormText name="ItemsDataUrl" value={tableSchema?.StructSchema.ItemsDataUrl} label="数据源" placeholder="/api/table/demodata" rules={[{ required: true }]} />
            <ProFormText name="ItemUpdateUrl" value={tableSchema?.StructSchema.ItemUpdateUrl} label="更新地址" placeholder="/api/demo/post" />
            <ProFormText name="ItemDeleteUrl" value={tableSchema?.StructSchema.ItemDeleteUrl} label="删除地址" placeholder="/api/demo/post" /> 
        </StepsForm.StepForm>

        <StepsForm.StepForm name="items" title="数据字段" initialValues={tableSchema?.StructSchema}>
          <ProFormList name="Items" creatorButtonProps={{creatorButtonText: '添加数据字段'}}>
            <ProForm.Group>
              {/* <Input name="DataName" placeholder="请输入字段名" required />
              <Input name="Title" placeholder="请输入字段标题" required /> */}
              <ProFormText name="DataName" placeholder="请输入字段名" label="字段名" rules={[{ required: true }]} /> 
              <ProFormText name="Title" placeholder="请输入字段标题" label="字段标题" rules={[{ required: true }]} />
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
              <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" placeholder="请输入像素宽度" />
              <ProFormDigit label="权重" tooltip="查询表单中的权重，权重大排序靠前" min={0} max={300} name="Order" placeholder={"请输入权重"} />
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
    </>
  );
}

export default TableSchemaForm;

