import { PageContainer } from '@ant-design/pro-components';

import { Alert, Card, Typography } from 'antd';
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
        <Alert
          message={"开发神器: 代码自动生成工具"}
          type="success"
          showIcon
          banner
          style={{
            margin: -12,
            marginBottom: 24,
          }}
        />

{/* 
        <Typography.Text strong>
          <a
            href="#"
            rel="noopener noreferrer"
            target="__blank"
          >
            Welcome
          </a>
        </Typography.Text>
        <Card>
        编译客户端
        <CodePreview>cd webclient</CodePreview>
        <CodePreview>yarn</CodePreview>
        <CodePreview>yarn build</CodePreview>
        移动编译好的客户端文件至客户端资源目录
        <CodePreview>cd ..</CodePreview>
        <CodePreview>mv webclien/dist/* resource/client/</CodePreview>
        编译并运行服务端
        <CodePreview>go mod tidy</CodePreview>
        <CodePreview>go build</CodePreview>
        <CodePreview>./lemocoder</CodePreview>
        </Card>
*/}

      </Card>
    </PageContainer>
  );
};

export default Welcome;
