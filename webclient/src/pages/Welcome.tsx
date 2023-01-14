import { PageContainer } from '@ant-design/pro-components';

import { Alert, Card, List, Typography, Row, Timeline, Col, Space, Tag, Button, message } from 'antd';
// import React, {  } from 'react';
import { useState, useEffect } from 'react';
// import { Line } from '@ant-design/charts';
import { Gauge, Liquid, Progress } from '@ant-design/plots';
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
  UsedPercent: number;
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

type Percent = {
  memNum: number; cpuNum: number; memText: string; cpuText: string
}

// const Welcome: React.FC = () => {
  export default () => {
  const [isload, setIsload] = useState<boolean>(false);
  const [status, setStatus] = useState<Status>();
  const [percent, setPercent] =useState<Percent>({memNum:0, memText:"0%", cpuNum:0, cpuText:"0%"});
  const [memInfo, setMemInfo] =useState<{used: string; total: string;}>({used:"0GB", total:"0GB"});
  
  const refresh = async () => {
    const resp = await get<{Code: number; Data: Status; Msg: string}>("/api/os/getstatus")
    if (resp.Code!=200){
      await message.error(resp.Msg)
      return
    }
    const dt = resp.Data
    const cpuNum = (dt.CpuUsedPercent/100)
    const memNum = (dt.MemoryUsedPercent/100)
    setPercent({cpuNum:cpuNum, memNum:memNum, cpuText: dt.CpuUsedPercent.toFixed(0)+"%", memText:dt.MemoryUsedPercent + "%"})
    const memUsed = ((dt.MemoryTotal - dt.MemoryFree)/(1024*1024*1024)).toFixed(2) + "GB"
    const memTotal = (dt.MemoryTotal/(1024*1024*1024)).toFixed(2) + "GB"
    setMemInfo({used: memUsed, total: memTotal})
    setStatus(dt)
  }
  
  useEffect(()=>{
    if (!isload){
      refresh()
      setIsload(true)
    }
  })

  const configMem = {
    type: 'meter',
    // innerRadius: 0.75,
    range: {
      ticks: [0, 1 / 3, 2 / 3, 1],
      color: ['#30BF78', '#FAAD14', '#F4664A'],
    },

    indicator: {
      pointer: {
        style: {
          stroke: '#D0D0D0',
        },
      },
      pin: {
        style: {
          stroke: '#D0D0D0',
        },
      },
    },
    statistic: {
      content: {
        style: {
          fontSize: '16px',
          lineHeight: '16px',
          color: '#4B535E',
        },
        formatter: () => percent.cpuText,
      },
    },
  };
  // const memStatisticTitle = {style:{fontSize: "16px", color: '#4B535E'}, offsetY: -26, formatter: () => '内存'}
  const memStatistic = {content:{formatter: () => percent.memText}}

  return (
    <PageContainer>

        {/* <Alert message={"开发神器: 代码自动生成工具"} type="success" showIcon banner style={{margin: -12, marginBottom: 24, }} /> */}
        <Row style={{ marginBlockEnd: 12 }}>

          <Col span={8}>

            <List header={<Title level={5}>开发环境</Title>} bordered dataSource={status?.DevTools} style={{background:"#fff", marginRight: 8, marginBlockEnd: 8}}
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

            <List header={<Title level={5}>开发文档</Title>} bordered dataSource={[
              {title:"Gin 开发框架", url:"https://gin-gonic.com/zh-cn/docs/"},
              {title: "Ant Design Pro 前端框架", url: "https://pro.ant.design/zh-CN/docs/overview"},
              {title:"Ant Design Pro 高级组件", url:"https://procomponents.ant.design/components"},
              {title:"Ant Design 基础组件", url:"https://ant.design/components/overview-cn/"},
              {title:"Ant Design Charts 图表组件", url:"https://charts.ant.design/"},
              ]} style={{background:"#fff", marginRight: 8 }}
              renderItem={(d)=>{
                return (
                  <List.Item>
                    <a href={d.url} target="__blank" >{d.title}</a>
                  </List.Item>
                )
              }}
            />

          </Col>

          <Col span={10} >

            <Card title={"主机: "+ status?.HostName} style={{ marginRight: 8, marginBlockEnd: 8}} >
              <List bordered dataSource={[
                {title:"OS", info: status?.OS + " " + status?.Arch},
                {title: "CPU", info: status?.CpuName},
                {title: "RAM", info: memInfo.total},
                ]} style={{background:"#fff"}}
                renderItem={(d)=>{
                  return (
                    <List.Item>
                      <Space>{d.title} {d.info}</Space>
                    </List.Item>
                  )
                }}
              />

            </Card>

            <List header={<Title level={5}>磁盘</Title>} bordered dataSource={status?.DiskInfo} style={{background:"#fff", marginRight: 12}}
                renderItem={(d)=>{
                  let tagName = d.Device
                  if (d.Device != d.MountPoint) {
                    tagName = d.Device + "(" + d.MountPoint + ")"
                  }
                  const percent = d.UsedPercent / 100
                  let total = (d.Total/(1024*1024*1024)).toFixed(2) + "GB"
                  let used = ((d.Total-d.Free)/(1024*1024*1024)).toFixed(2) + "GB"
                  return (
                    <List.Item>
                      <Tag color="blue">{tagName}</Tag>{used}/{total}
                      {/* <Space> */}
                      <Progress width={300} height={50} autoFit={false} barWidthRatio={0.2} percent={percent} color={['#F4664A', '#93dfb8']} />
                      {/* </Space> */}
                    </List.Item>
                  )
                }}
              />

          </Col>

          <Col span={6}>
            <Card title={"CPU: " + status?.CpuNum + " cores"} style={{ marginBlockEnd: 16, marginRight: 8 }}>
                <Gauge width={100} height={150} percent={percent.cpuNum} {...configMem} />
            </Card>
            <Card title={"内存: "+ memInfo.used + "/" + memInfo.total}>
              <Liquid width={100} height={150} percent={percent.memNum} statistic={memStatistic} wave={{length: 180}} />
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
