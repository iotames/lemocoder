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
            <List header={<Title level={5}>开发工具</Title>} bordered dataSource={status?.DevTools} style={{background:"#fff"}}
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
                    </Space>
                    {btn}
                  </List.Item>
                )
              }}
            />
          </Col>
          <Col span={8}></Col>
          <Col span={8}>
            <List header={<Title level={5}>开发文档</Title>} bordered dataSource={[
              {title:"Gin开发框架", url:"https://gin-gonic.com/zh-cn/docs/"},
              {title:"Ant Design Pro 组件", url:"https://procomponents.ant.design/components"},
              {title:"Ant Design 组件", url:"https://ant.design/components/overview-cn/"},
              ]} style={{background:"#fff"}}
              renderItem={(d)=>{
                return (
                  <List.Item>
                    <a href={d.url} target="__blank" >{d.title}</a>
                    </List.Item>
                )
              }}
            />
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
