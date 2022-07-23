# week4

#### 内容
1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

#### 项目结构说明
```
├── api
├── cmd
│   └── app
├── configs
├── internal
│    ├── biz
│    ├── conf
│    ├── data
│    ├── server
│    └── service
├── pkg
```

* [api](#api): 接口定义的目录，如果我们采用的是 grpc 那这里面一般放的就是 proto 文件

* [cmd](#cmd): 目录下一般是项目的主干，负责程序的生命周期，服务所需资源的依赖注入等
    * [app](#app): 对服务进行分类，代表具体某一个app应用，建议按功能进行命名
    
* [config](#config): 框架全局配置文件目录

* [internal](#internal): internal 目录下的包，不允许被其他项目中进行导入
    * [biz](#biz): 业务逻辑的组装、处理层
    * [data](#data): 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口
    * [server](#server): 提供快捷的启动服务全局方法
    * [service](#service): 处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑

* [pkg](#pkg): 放置可以被外部程序安全导入的包
