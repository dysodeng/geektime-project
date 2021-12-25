#### 第四周作业-项目结构

- 目录结构
```
project
|-- api
|   |-- proto proto定义
|-- internal  内部代码包
|   |-- entity 数据实体
|   |-- provider 服务容器注入
|   |-- service 业务服务
|       |-- contracts 服务契约接口
|       |-- service.go 服务容器实例
|-- pkg  第三方代码包
|-- server http与grpc服务启动相关
|   |-- grpc gRPC服务
|   |   |-- service proto服务实现
|   |   |-- server.go grpc服务启动
|   |-- http http服务
|       |-- router.go http路由
|       |-- server.go http服务启动
|-- var 日志等临时文件目录
|-- main.go 主入口
|-- ...
```

- 由于时间原因，直接用的我以前项目的目录结构做了精简
- github.com/dysodeng/aux-rpc包是我自己根据gRPC简单封装的一个包
- 没有做配置相关的处理，直接使用的os.Getenv()从环境变量获取的配置
- wire包实在用不来，没有搞明白怎么用，所以这里没有使用wire包，而是使用的另一个di包做的服务容器
- 关于wire包的问题：如果依赖的服务需要额外传不同的参数时怎么处理呢？目前我研究的结果是：如果需要传额外的参数，
  如果2个服务的参数是同一个数据类型，那么只能使用同一个参数，当传多个同类型参数时，编译错误
![img.png](https://github.com/dysodeng/geektime-project/blob/master/image/img.png)
![img_1.png](https://github.com/dysodeng/geektime-project/blob/master/image/img_1.png)
![img_3.png](https://github.com/dysodeng/geektime-project/blob/master/image/img_3.png)
- 只传同一种类型只传一个，就能编译通过，但是又不符合业务逻辑
![img_2.png](https://github.com/dysodeng/geektime-project/blob/master/image/img_2.png)