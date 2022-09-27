## 环境配置

1. 安装nodejs

前往 https://nodejs.org/zh-cn/ 下载并安装node.
Linux系统可尝试命令: `sudo apt install nodejs`

安装完毕后，运行 `node -v` 确认是否安装成功

2. 安装yarn

```
npm install -g yarn
yarn --version
```

设置yarn国内源: `yarn config set registry https://registry.npm.taobao.org -g`

3. 安装golang开发环境

下载并安装Go: https://golang.google.cn/doc/install

设置GO国内代理:

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```


## 快速开始

```
# 加载依赖包
go mod tidy

# 初始化本地sqlite3数据库, 生成 WebClient 客户端资源
# (cd webclient && yarn && yarn build && cd .. && mv webclient/dist/* resource/client/)
go run . init

# 运行
go run .
```

若执行 `go mod tidy` 命令，加载依赖包出错，可能为GO版本过低或第三方包版本冲突。
手动重置 `mod.go` 文件内容, 再重试。

```
module lemocoder

go 1.19
```


## WebServer 服务端

```
go mod tidy
go run . init
go run .
```

编译成可执行文件后运行
```
go build
./lemocoder init
./lemocoder
```


## WebClient 客户端

```
cd webclient
# 下载依赖包
yarn
# 开发模式启动
yarn start
# 编译前端资源包
yarn build
```

## 笔记

webpack5 新特性: `experiments.topLevelAwait`
