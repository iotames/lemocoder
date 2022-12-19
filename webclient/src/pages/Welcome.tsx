import { PageContainer } from '@ant-design/pro-components';

import { Alert, Card, Typography, Row, Timeline, Col, message } from 'antd';
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



type OsStatus = {
  vnode: string;
  vyarn: string;
  vgit: string;
  vgo: string;
}



// const Welcome: React.FC = () => {
  export default () => {
  const [isload, setIsload] = useState<boolean>(false);
  const [osstatus, setosstatus] = useState<OsStatus>({vnode:"",vyarn:"",vgit:"",vgo:""});

  const refresh = async () => {
    const resp = await get<{Code: number; Data: OsStatus; Msg: string}>("/api/os/getstatus")
    if (resp.Code!=200){
      await message.error(resp.Msg)
      return
    }
    setosstatus(resp.Data)
  }
  
  useEffect(()=>{
    if (!isload){
      refresh()
      setIsload(true)
    }
  })

  return (
    <PageContainer>
      <Card>
        {/* <Alert message={"开发神器: 代码自动生成工具"} type="success" showIcon banner style={{margin: -12, marginBottom: 24, }} /> */}
        <Row>
          <Col span={16}>
            <Timeline>
              <Timeline.Item color={osstatus.vgo == "" ? "red" : "green"}>go {osstatus.vgo}</Timeline.Item>
              <Timeline.Item color={osstatus.vnode == "" ? "red" : "green"}>node: {osstatus.vnode}</Timeline.Item>
              <Timeline.Item color={osstatus.vyarn == "" ? "red" : "green"}>yarn {osstatus.vyarn}</Timeline.Item>
              <Timeline.Item color={osstatus.vgit == "" ? "red" : "green"}>git {osstatus.vgit}</Timeline.Item>

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
          <Col span={8}>
            <Row style={{ marginBlockEnd: 16 }}>
              <Typography.Text strong>
                <a href="https://procomponents.ant.design/components" target="__blank" >Ant Design Pro 组件</a>
              </Typography.Text>
            </Row>
            <Row style={{ marginBlockEnd: 16 }}>
              <Typography.Text strong>
                  <a href="https://ant.design/components/overview-cn/" rel="noopener noreferrer" target="__blank" >Ant Design 组件</a>
              </Typography.Text>
            </Row>
          </Col>
        </Row>

        {/* 
          <Card>
          <CodePreview>go mod tidy</CodePreview>
          <CodePreview>go run . init</CodePreview>
          </Card> 
        */}

      </Card>
    </PageContainer>
  );
};

// export default Welcome;
