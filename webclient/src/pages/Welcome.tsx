import { PageContainer } from '@ant-design/pro-components';

import { Alert, Card, Typography, Row, Timeline, Col } from 'antd';
import React from 'react';
import { SmileOutlined } from '@ant-design/icons';
// import styles from './Welcome.less';

// const CodePreview: React.FC = ({ children }) => (
//   <pre className={styles.pre}>
//     <code>
//       <Typography.Text copyable>{children}</Typography.Text>
//     </code>
//   </pre>
// );

const Welcome: React.FC = () => {

  return (
    <PageContainer>
      <Card>
        {/* <Alert message={"开发神器: 代码自动生成工具"} type="success" showIcon banner style={{margin: -12, marginBottom: 24, }} /> */}
        <Row>
          <Col span={16}>
            <Timeline>
              <Timeline.Item color="green">后端开发环境</Timeline.Item>
              <Timeline.Item color="green">前端开发环境</Timeline.Item>
              <Timeline.Item color="red">
                <p>Solve initial network problems 1</p>
                <p>Solve initial network problems 2</p>
                <p>Solve initial network problems 3 2015-09-01</p>
              </Timeline.Item>
              <Timeline.Item>
                <p>Technical testing 1</p>
                <p>Technical testing 2</p>
                <p>Technical testing 3 2015-09-01</p>
              </Timeline.Item>
              <Timeline.Item color="gray">
                <p>Technical testing 1</p>
                <p>Technical testing 2</p>
                <p>Technical testing 3 2015-09-01</p>
              </Timeline.Item>
              <Timeline.Item color="gray">
                <p>Technical testing 1</p>
                <p>Technical testing 2</p>
                <p>Technical testing 3 2015-09-01</p>
              </Timeline.Item>
              <Timeline.Item color="#00CCFF" dot={<SmileOutlined />}>
                <p>Custom color testing</p>
              </Timeline.Item>
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

export default Welcome;
