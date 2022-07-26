## 快速开始

```
# 编译客户端
cd webclient
yarn
yarn build
# 移动编译好的客户端文件至客户端资源目录
cd ..
mv webclien/dist/* resource/client/
# 编译并运行服务端
go mod tidy
go build
./lemocoder
```


## WebServer 服务端

```
go mod tidy
go run .
```


## WebClient 客户端

```
cd webclient
yarn
# 开发模式启动
yarn start
```

## 笔记

webpack5 新特性: `experiments.topLevelAwait`
