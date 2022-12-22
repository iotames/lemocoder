import { PageContainer, ProList } from '@ant-design/pro-components';

import { Alert, Card, Typography, Row, Timeline, Col, Space, Tag, message } from 'antd';
// import React, {  } from 'react';
import { useState, useEffect } from 'react';
import { SmileOutlined } from '@ant-design/icons';
import {post, postMsg, getTableData, postByBtn,get} from "@/services/api"
// import styles from './Welcome.less';

// const CodePreview: React.FC = ({ children }) => (
//   <pre className={styles.pre}>
//     <code>
//       <Typography.Text copyable>{children}</Typography.Text>
//     </code>
//   </pre>
// );

type DevTool = {
  Name: string;
  Version: string;
  Url: string;
}
type Status = {
  HostName: string;
  OS: string;
  Arch: string;
  DevTools: DevTool[]
}



// const Welcome: React.FC = () => {
  export default () => {
  const [isload, setIsload] = useState<boolean>(false);
  const [status, setStatus] = useState<Status>();
  


  const refresh = async () => {
    const resp = await get<{Code: number; Data: Status; Msg: string}>("/api/os/getstatus")
    if (resp.Code!=200){
      await message.error(resp.Msg)
      return
    }
    setStatus(resp.Data)
  }
  
  useEffect(()=>{
    if (!isload){
      refresh()
      setIsload(true)
    }
  })

  return (
    <PageContainer>

        {/* <Alert message={"开发神器: 代码自动生成工具"} type="success" showIcon banner style={{margin: -12, marginBottom: 24, }} /> */}
        <Row>
          <Col span={12}>

          <ProList<DevTool> rowKey="Name" headerTitle="开发工具" dataSource={status?.DevTools} // showActions="hover"
      // onDataSourceChange={}
      metas={{
        title: {
          dataIndex: 'Name',
        },
        // avatar: {
        //   dataIndex: 'image',
        //   editable: false,
        // },
        // description: {
        //   dataIndex: 'desc',
        // },
        subTitle: {
          render: (e, d) => {
            let color = "blue"
            let version = d.Version
            if (d.Version == "") {
              color = "red"
              version = "未发现"
            }
            return (
              <Space size={0}>
                <Tag color={color}>{version}</Tag>
                {/* <Tag color="#5BD8A6">TechUI</Tag> */}
              </Space>
            );
          },
        },
        actions: {
          render: (text, row, index, action) => {
            let btn = (<></>)
            if (row.Version == "") {
              btn = (<a href={row.Url} target="__blank" >下载</a>) // <a onClick={() => {}} key="link">  安装 </a>,
            }
            return ([btn])
          },
        },
      }}
    />

            <Timeline>
              {/* <Timeline.Item color={osstatus.vgo == "" ? "red" : "green"}>go {osstatus.vgo}</Timeline.Item>
              <Timeline.Item color={osstatus.vnode == "" ? "red" : "green"}>node {osstatus.vnode}</Timeline.Item>
              <Timeline.Item color={osstatus.vyarn == "" ? "red" : "green"}>yarn {osstatus.vyarn}</Timeline.Item>
              <Timeline.Item color={osstatus.vgit == "" ? "red" : "green"}>git {osstatus.vgit}</Timeline.Item> */}
              {/* <Timeline.Item color="gray">
                <p>Technical testing 1</p>
                <p>Technical testing 2</p>
                <p>Technical testing 3 2015-09-01</p>
              </Timeline.Item>

              <Timeline.Item color="#00CCFF" dot={<SmileOutlined />}>
                <p>Custom color testing</p>
              </Timeline.Item> */}

            </Timeline>

          </Col>
          <Col span={4}></Col>
          <Col span={8}>
            <Card title="开发文档">
            <Row style={{ marginBlockEnd: 16 }}>
              <Typography.Text strong>
                  <a href="https://gin-gonic.com/zh-cn/docs/" target="__blank" >Gin开发框架</a>
              </Typography.Text>
            </Row>
            <Row style={{ marginBlockEnd: 16 }}>
              <Typography.Text strong>
                <a href="https://procomponents.ant.design/components" target="__blank" >Ant Design Pro 组件</a>
              </Typography.Text>
            </Row>
            <Row style={{ marginBlockEnd: 16 }}>
              <Typography.Text strong>
                  <a href="https://ant.design/components/overview-cn/" target="__blank" >Ant Design 组件</a>
              </Typography.Text>
            </Row>
            </Card>
          </Col>
        </Row>

        {/* 
          <Card>
          <CodePreview>go mod tidy</CodePreview>
          <CodePreview>go run . init</CodePreview>
          </Card> 
        */}

    </PageContainer>
  );
};

// export default Welcome;
