import { PageContainer, ProCard } from '@ant-design/pro-components';

import { Alert, Card, List, Typography, Row, Timeline, Col, Space, Tag, message } from 'antd';
// import React, {  } from 'react';
import { useState, useEffect } from 'react';

import { SmileOutlined } from '@ant-design/icons';
import {post, postMsg, getTableData, postByBtn,get} from "@/services/api"
// import styles from './Welcome.less';
const {Title} = Typography;

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
type DiskPart = {
  Device: string;
  MountPoint: string;
  UsedPrecent: number;
  Total: number;
  Free: number; 
}
type Status = {
  HostName: string;
  OS: string;
  Arch: string;
  CpuName: string;
  CpuNum: number;
  CpuUsedPercent: number;
  MemoryUsedPercent: number;
  MemoryTotal: number;
  MemoryFree: number;
  MemoryTotalText: string;
  MemoryFreeText: string;
  DiskInfo: DiskPart[]
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
        <Row style={{ marginBlockEnd: 16 }}>
          
          <ProCard title="主机信息" style={{ maxWidth: 600 }}>
            <div>{status?.HostName}</div>
            <div>{status?.CpuName}</div>
          </ProCard>

        </Row>
        <Row>
          <Col span={8}>

          <List header={<Title level={5}>开发工具</Title>} bordered dataSource={status?.DevTools}
            renderItem={(d)=>{
              let color = "blue"
              let version = d.Version
              let btn = (<></>)
              if (d.Version == "") {
                color = "red"
                version = "未发现"
                btn = (<a href={d.Url} target="__blank" >下载</a>) // <a onClick={() => {}} key="link">  安装 </a>,
              }
              return (
                <List.Item>
                <Space>
                  {d.Name}
                <Tag color={color}>{version}</Tag>
                {btn}
                </Space>
                </List.Item>
              )
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
