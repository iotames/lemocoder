import { PageContainer } from '@ant-design/pro-components';

import { Alert, Card, Typography, Row, Col } from 'antd';
import React from 'react';
import styles from './Welcome.less';

const CodePreview: React.FC = ({ children }) => (
  <pre className={styles.pre}>
    <code>
      <Typography.Text copyable>{children}</Typography.Text>
    </code>
  </pre>
);

const Welcome: React.FC = () => {

  return (
    <PageContainer>
      <Card>
        <Alert message={"开发神器: 代码自动生成工具"} type="success" showIcon banner style={{margin: -12, marginBottom: 24, }} />

      <Row>
      <Typography.Text strong>
                <a href="https://procomponents.ant.design/components" target="__blank" >Ant Design Pro 组件</a>
              </Typography.Text>
      </Row>



        <Card>
        <CodePreview>go mod tidy</CodePreview>
        <CodePreview>go run . init</CodePreview>
        <CodePreview>go run </CodePreview>
        </Card>


        <Row>
        <Typography.Text strong>
                  <a href="https://ant.design/components/overview-cn/" rel="noopener noreferrer" target="__blank" >Ant Design 组件</a>
                </Typography.Text>
        </Row>

 
      </Card>
    </PageContainer>
  );
};

export default Welcome;
