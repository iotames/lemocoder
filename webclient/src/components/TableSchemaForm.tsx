import {ProFormSwitch, ProFormInstance, ProFormDigit, ProFormList, ProFormText, ProFormSelect, ProForm, StepsForm, ProFormGroup, ModalForm} from '@ant-design/pro-components';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import type { FormInstance, StepsProps } from 'antd';
// import TableSchema from '@/pages/TableSchema';
import { TableItemFormFields, TableItemOptFormFields } from "@/components/Form"

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
  Searchable: boolean;
  Title: string;
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
          <TableItemFormFields />
          </ProFormList>
        </StepsForm.StepForm>

        <StepsForm.StepForm name="options" title="行数据操作" initialValues={tableSchema?.StructSchema}>
        <ProFormList name="ItemOptions" creatorButtonProps={{creatorButtonText: '添加行数据操作项'}}>
        <TableItemOptFormFields />
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
  const postUrl = "/api/coder/table/create"
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
      
      <ProForm.Group>
        <ProFormText width={150} name={["StructSchema", "ItemDataTypeName"]} label="数据结构名" tooltip="例: Product, ProductReview" placeholder="例: Product" rules={[{ required: true }]} />
        <ProFormText name={["StructSchema", "Title"]} label="数据表格标题" placeholder="例: 商品列表" />
        <ProFormSwitch name={["StructSchema", "Searchable"]} label="数据搜索" checkedChildren="已开启" unCheckedChildren="已禁用" initialValue={true} />
      </ProForm.Group>

      <ProFormList name={["StructSchema", "Items"]} creatorButtonProps={{creatorButtonText: '添加数据字段'}}>
        <TableItemFormFields />
      </ProFormList>

    </ModalForm>
    </>
  );
}