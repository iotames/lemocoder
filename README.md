## 快速开始

```
# 编译主程序
go mod tidy
go build

# 初始化sqlite3数据库, 编译客户端, 移动编译好的客户端文件至客户端资源目录
# 以下命令相当于执行: cd webclient && yarn && yarn build && cd .. && mv webclien/dist/* resource/client/
./lemocoder init

# 运行
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
# 编译前端资源包
yarn build
```

## 笔记

webpack5 新特性: `experiments.topLevelAwait`
