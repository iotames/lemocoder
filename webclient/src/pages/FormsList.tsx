import { InfoCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, message, Space, Tag, Tooltip, Modal } from 'antd';
import type { UploadProps, UploadFile } from 'antd';
import { Search, Table, useTable, withTable } from 'table-render';
import FormRender, { useForm } from 'form-render';
import request from 'umi-request';
import { useState } from 'react';
import FormModal from "@/components/FormModal";

const searchSchema = {
  type: 'object',
  properties: {
    state: {
      title: '状态',
      type: 'string',
      enum: ['open', 'closed'],
      enumNames: ['营业中', '已打烊'],
      width: '25%',
      widget: 'select',
    },
    labels: {
      title: '等级',
      type: 'string',
      width: '25%',
    },
    created_at: {
      title: '创建时间',
      type: 'string',
      format: 'date',
      width: '25%',
    },
  },
  labelWidth: 80,
};




const formSchema = {
  type: 'object',
  properties: {
    title: {
      title: '标题',
      type: 'string',
      required: true,
      default: ""
    },
    upfile: {
      title: "上传文件",
      type: "string",
      format: "upload",
      props: {
        action: "http://127.0.0.1:8888/api/public/upload",
        onChange: (info) => {console.log("onChange----upfile", info.file.status)}
      },
      displayType: "row",
    },
    upfilepath: {
      hidden: true,
      type: "string",
    },
    url: {
      required: true,
      title: '提交地址',
      type: 'string',
      default: ""
    },
    formSchema: {
      required: true,
      title: "表单结构",
      type: "string",
      format: "textarea",
      default: ""
    },
    // status: {
    //   title: '状态',
    //   default: "b",
    //   type: 'string',
    //   enum: ['a', 'b', 'c'],
    //   enumNames: ['早', '中', '晚'],
    // },
  },
};


const Demo = () => {

  // const { refresh, tableState, setTable } = useTable();

  const searchApi = (params, sorter) => {
    console.group(sorter);
    return request
      .get(
        'https://www.fastmock.site/mock/62ab96ff94bc013592db1f67667e9c76/getTableList/api/basic',
        { params }
      )
      .then(res => {
        if (res && res.data) {
          return {
            rows: [...res.data, { money: null }],
            total: res.data.length,
          };
        }
      })
      .catch(e => {
        console.log('Oops, error', e);
        // 注意一定要返回 rows 和 total
        return {
          rows: [],
          total: 0,
        };
      });
  };

  // 配置完全透传antd table
  const columns = [
    {
      title: '酒店名称',
      dataIndex: 'title',
      valueType: 'text',
      width: '20%',
    },
    {
      title: '酒店地址',
      dataIndex: 'address',
      ellipsis: true,
      copyable: true,
      valueType: 'text',
      width: '25%',
    },
    {
      title: "酒店状态",
      enum: {
        open: '营业中',
        closed: '已打烊',
      },
      dataIndex: 'state',
    },

    {
      title: '酒店GMV',
      key: 'money',
      sorter: true,
      dataIndex: 'money',
      valueType: 'money',
    },
    {
      title: '成立时间',
      key: 'created_at',
      dataIndex: 'created_at',
      valueType: 'date',
    },
    {
      title: '操作',
      render: () => (
        <Space>
          <a target="_blank" key="1">
            <div
              onClick={() => {
                message.success('预订成功');
              }}
            >
              预订
            </div>
          </a>
        </Space>
      ),
    },
  ];

  const [visible, setVisible] = useState(false);

  return (
    <div>
      <FormModal title='添加表单' formSchema={formSchema} postUrl="/api/forms/add" visible={visible} setVisible={setVisible} />

      <Search schema={searchSchema} displayType="row" api={searchApi} />
      <Table
        pagination={{ pageSize: 4 }}
        columns={columns}
        rowKey="id"
        toolbarRender={() => [
          <Button
            key="primary"
            type="primary"
            onClick={()=>setVisible(true)}
          >
            <PlusOutlined />
            创建
          </Button>,
        ]}
        toolbarAction
      />
    </div>
  );
};

export default withTable(Demo);