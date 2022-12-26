import { ProCard, ProFormItem, ProFormSwitch, ProFormSelect, ProFormGroup, ProFormInstance, ProFormDigit, ProFormList, ProTable, ModalForm, ProFormText, ProForm, StepsForm, PageContainer, CheckCard } from '@ant-design/pro-components';
import { Row, Col, Button, Alert, Popover, Popconfirm, message, Steps, Modal, InputNumber, Input, Select, Typography } from 'antd';
import { useRef, useState, useEffect } from 'react';
import {post, postMsg, getTableData, postByBtn, get} from "@/services/api"
import type { TableSchema } from "@/components/TableSchemaForm"
import { TableItemFormFields, TableItemOptFormFields } from "@/components/Form"
import { PlayCircleOutlined, QuestionCircleOutlined, ArrowRightOutlined  } from '@ant-design/icons';
import { history } from 'umi';

const { Title, Paragraph, Text } = Typography;

export default () => {
  // const [tableSchema, setTableSchema] = useState<TableSchema>();
  const [pageState, setPageState] = useState<number>(0);
  const [pageID, setPageID] = useState<string>();
  const tableFormRef = useRef<ProFormInstance<TableSchema>>();
  const refresh = async (pageID: string) => {
    const pageResp = await get<{Code: number; Msg: string; Data:{State: number}}>("/api/coder/page/get", {"id": pageID})
    if (pageResp.Code == 500){
      await message.error(pageResp.Msg)
      return
    }
    if (pageResp.Code == 404) {
      await message.error(pageResp.Msg)
      return
    }
    setPageState(pageResp.Data.State)
    const resp = await get<{Code: number; Msg: string; Data: TableSchema}>("/api/coder/table/get", {"page_id": pageID})
    if (resp.Code == 500){
      await message.error(resp.Msg)
      return
    }
    if (resp.Code == 404) {
      await message.error("数据不存在")
      return
    }
    if (resp.Code == 200) {
      // setTableSchema(resp.Data)
      tableFormRef.current?.setFieldsValue(resp.Data); 
    }
  }
  const queryArgs = history.location.query
  useEffect(() => {
    if (queryArgs?.page_id != undefined) {
        // console.log(queryArgs.page_id)
        setPageID(queryArgs.page_id)
        refresh(queryArgs.page_id)
        // (async ()=>{await refresh(queryArgs.page_id);})()
    } else {
        message.error("page_id 参数不正确")
    }
  }, [queryArgs]);

    const TableSchemaUpdateForm = (
    <ProForm<TableSchema>
    submitter={{ searchConfig: { submitText: '修改'}, resetButtonProps:{style:{display:"none"}}}}
      onFinish={async(values) => {
        await postMsg("/api/coder/table/update", values)
      }}
      formRef={tableFormRef}
      // initialValues={tableSchema}
      >
        <ProFormText name="ID" hidden />
        <ProFormText name="PageID" hidden />
        <ProFormText name={["StructSchema", "RowKey"]} hidden />

        
        <ProForm.Group>
          <ProFormText width={180} name={["StructSchema", "ItemsDataUrl"]} label="数据源" rules={[{ required: true }]} />
          <ProFormText width={180} name={["StructSchema", "ItemCreateUrl"]} label="新增数据地址" />
          <ProFormText width={180} name={["StructSchema", "ItemUpdateUrl"]} label="更新数据地址" />
          <ProFormText width={180} name={["StructSchema", "ItemDeleteUrl"]} label="删除数据地址" /> 
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText width={120} name={["StructSchema", "ItemDataTypeName"]} label="数据结构名" rules={[{ required: true }]} />
          <ProFormText name={["StructSchema", "Title"]} label="数据表格标题" placeholder="例: 商品列表" />
          <ProFormSwitch name={["StructSchema", "Searchable"]} label="数据搜索" checkedChildren="已开启" unCheckedChildren="已禁用" initialValue={true} />
        </ProForm.Group>

        {/* <ProForm.Group label="数据字段"> */}
        <ProCard title="数据字段" style={{ marginBlockStart: 16 }} >
          <ProFormList 
          name={["StructSchema", "Items"]} 
          creatorButtonProps={{creatorButtonText: '添加数据字段'}}
          >
            
            <TableItemFormFields />
            
          </ProFormList>
        </ProCard>
        {/* </ProForm.Group> */}

        <ProCard title="行数据操作" style={{ marginBlockStart: 16 }}>
            <ProFormList name={["StructSchema", "ItemOptions"]} creatorButtonProps={{creatorButtonText: '添加行数据操作项'}}>
            <TableItemOptFormFields />
          </ProFormList>
        </ProCard>

        <ProCard title="批量数据操作" style={{ marginBlock: 16 }}>
          <ProFormList name={["StructSchema", "BatchOptButtons"]} creatorButtonProps={{creatorButtonText: '添加批量数据操作项'}}>
            <ProFormGroup>
                <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} />
                <ProFormText name="Url" label="API地址" rules={[{ required: true }]} />
            </ProFormGroup>
          </ProFormList>
        </ProCard>

        <ProCard title="表单提交" style={{ marginBlock: 16 }}>
          <ProFormList name={["StructSchema", "ItemForms"]} creatorButtonProps={{creatorButtonText: '添加表单'}}>
            <ProFormGroup>
                <ProFormText name="Key" label="表单标识(Key)" rules={[{ required: true }]} />
                <ProFormText name={["Form", "SubmitUrl"]} label="表单提交地址" rules={[{ required: true }]} />
                <ProFormText name={["Form", "Title"]} label="表单标题" />
            </ProFormGroup>

            <ProFormList name={["Form", "FormFields"]} creatorButtonProps={{creatorButtonText: '添加表单域'}}>
              <ProFormGroup>
                  <ProFormText name="Name" label="字段名" width={100} rules={[{ required: true }]} />
                  <ProFormText name="Label" label="标题" width={100} />
                  <ProFormSelect name="Component" label="组件" width={120} rules={[{ required: true }]} options={[
                    {value:"ProFormText", label:"文本框"},
                    {value:"ProFormTextArea", label:"文本(TextArea)"},
                    {value:"ProFormDigit", label:"数字输入框"},
                    {value:"ProFormSelect", label:"选择框"},
                    ]} />
                  <ProFormText name="Placeholder" label="Placeholder" width={90} />
                  <ProFormText name="Width" label="宽度" width={90} />
              </ProFormGroup>
            </ProFormList>

          </ProFormList>
        </ProCard>

    </ProForm>)

    const htlpTitle = (
    <Popconfirm placement="right" title="是否执行?" onConfirm={async()=>{ await postMsg("/api/coder/project/rebuild", {"PageID": pageID})}} okText="是" cancelText="否">
      <Button type="primary">后续操作</Button>
    </Popconfirm>);

    const htlpContent = (
      <div>
        <p><Text copyable>go run . clientinit</Text></p>
        <p><Text copyable>go run . dbsync</Text></p>
        <p><Text copyable>go build .</Text></p>
      </div>
    );
    const htlpBtn = (
      <Popover placement="right" title={htlpTitle} content={htlpContent} trigger="click">
      <Button type='ghost' shape='circle' icon={<QuestionCircleOutlined />}></Button>
    </Popover>
    );

  const codeGen = (
    <Popconfirm placement="right" title="是否生成代码?" onConfirm={
      async()=>{
      await postMsg("/api/coder/table/createcode", {"PageID": pageID})
    }} okText="是" cancelText="否">
    <Button type='primary' style={{marginRight:6}} shape="circle" icon={<PlayCircleOutlined />} ></Button>
    </Popconfirm>
  );
  let topBtn = (<Col></Col>)
  if (pageState < 2){
    topBtn = codeGen
  }

  return (
    <PageContainer>
      <Row>
        <Alert
          message="后续操作"
          description="生成代码后, 请重新编译前后端项目, 并同步数据表.  编译后的前端资源文件(如: umi.js), 如因缓存问题而没有生效, 请手动清除浏览器缓存."
          type="warning" showIcon closable
          style={{margin: -12, marginBottom: 24, }} />
      </Row>
      <Row  style={{ marginBlockEnd: 16 }}>

<Col span={18}>

      {/* <Steps size="small" current={1} items={[{title: (topBtn),},{title: 'In Progress',},]} /> */}

        <span>{topBtn} <ArrowRightOutlined  /> {htlpBtn}</span>
        </Col>
      </Row>
      <Row>
        
          {TableSchemaUpdateForm}
      </Row>
      
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
    </PageContainer>
  );
};