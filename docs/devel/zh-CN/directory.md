## miniblog 项目目录说明

```bash
├── api/ # Swagger / OpenAPI 文档存放目录
│   └── openapi/
│       └── miniblog/
│           └── v1/
│               ├── openapi.yaml # penAPI 3.0 API 接口文档
├── build/ # 构建物料目录
│   └── docker/
│       └── mb-apiserver/
│           └── Dockerfile # mb-apiserver 组件的 Dockerfile 文件
├── cmd/ # main 文件存放目录
│   ├── gen-gorm-model/ # 根据数据库表结构生成 GORM Model
│   │   └── gen_gorm_model.go
│   └── mb-apiserver/ # mb-apiserver 应用层代码入口
│       ├── app/
│       │   ├── config.go # viper 配置文件
│       │   ├── options/ # 定义运行时配置项
│       │   │   └── options.go
│       │   └── server.go # 应用框架实现文件
│       └── main.go # main 文件
├── configs/ # # 配置文件存放目录
│   ├── mb-apiserver.yaml # mb-apiserver 组件配置文件
│   ├── miniblog.sql # 数据库初始化 SQL
│   ├── nginx.loadbalance.conf # nginx 负载均衡配置
│   └── nginx.reverse.conf # nginx 反向代理配置
├── deployments/ # Kubernetes 部署物料
│   ├── mb-apiserver-configmap.yaml # ConfigMap YAML 定义
│   ├── mb-apiserver-deployment.yaml # Deloyment YAML 定义
│   └── mb-apiserver-service.yaml # Service YAML 定义
├── docs/ # 项目文档
│   ├── book/ # 《从零开发企业级Go应用》书中知识补充相关文档
│   ├── devel/ # 开发文档
│   │   │── en-US # 英文文档
│   │   └── zh-CN/ # 中文文档
│   │       ├── architecture.md # miniblog 架构介绍
│   │       ├── conversions/ # 规范文档存放目录
│   │       │   ├── api.md # 接口规范
│   │       │   ├── commit.md # Commit 规范
│   │       │   ├── directory.md # 目录结构规范
│   │       │   ├── error_code.md # 错误码规范
│   │       │   ├── go_code.md # 代码规范
│   │       │   ├── logging.md # 日志规范
│   │       │   ├── README.md 
│   │       │   └── version.md # 版本规范
│   │       ├── directory.md
│   │       ├── mysql.md
│   │       ├── README.md
│   │       └── useful_commands.md
│   │── guide/ # 用户文档
│   │   │── en-US # 英文文档
│   │   └── zh-CN/ # 中文文档
│   │       ├── announcements.md # 动态与公
│   │       ├── best-practice/ # 最佳实践
│   │       ├── faq/ # 常见问题
│   │       ├── installation/ # 安装指南
│   │       ├── introduction/ # 产品介绍
│   │       ├── operation-guide/ # 操作指南
│   │       ├── quickstart/ # 快速入门
│   └── images # 项目图片存放目录
├── examples/ # 示例源码
│   ├── client/ # 客户端调用示例
│   │   ├── health/ # 健康检查示例
│   │   ├── interceptor/ # gRPC 拦截器示例
│   │   ├── post/ # 博客操作示例
│   │   └── user/ # 用户操作示例
│   ├── errorsx/ # errorsx 包使用示例
│   ├── gin/ # Gin 框架使用示例
│   │   └── middleware/ # Gin 框架中间件使用示例
│   ├── helper/ # 示例源码共享函数
│   ├── logpattern/ # 日志包开发模式示例
│   ├── performance/ # 性能测试示例
│   │   ├── benchmark/
│   │   │   └── benchmark_test.go
│   │   ├── nethttppprof/
│   │   │   └── main.go
│   │   └── runtimepprof/
│   │       └── main.go
│   └── simple.mk # 最简单的 Makefile 文件示例
├── go.mod # Go Modules 文件
├── go.sum # Go 模块 checksum 文件
├── go.work # Go 工作区文件
├── go.work.sum
├── init/ # Systemd Unit 文件保存目录
│   ├── mb-apiserver.service # mb-apiserver systemd unit
│   └── README.md
├── internal/ # 内部代码保存目录，这里面的代码不能被外部程序引用
│   ├── miniblog/ # miniblog 代码实现目录
│   │   ├── biz/ # Biz 层代码
│   │   │   ├── biz.go
│   │   │   ├── README.md
│   │   │   ├── v1/ # Biz 层 v1 版本实现
│   │   │   │   ├── post/ # 博客实现
│   │   │   │   └── user/ # 用户实现
│   │   │   └── v2/ # Biz 层 v2 版本实现
│   │   ├── grpcserver.go # gRPC 服务器实现代码
│   │   ├── handler/ # Handler 层代码
│   │   │   ├── grpc/ # gRPC Handler 层代码
│   │   │   │   ├── handler.go # Handler 实现定义
│   │   │   │   ├── healthz.go # 健康检查接口
│   │   │   │   ├── post.go # 博客接口
│   │   │   │   ├── post_test.go
│   │   │   │   └── user.go # 用户接口
│   │   │   └── http/ # HTTP Handler 层代码
│   │   │       ├── handler.go
│   │   │       ├── healthz.go
│   │   │       ├── post.go
│   │   │       └── user.go
│   │   ├── httpserver.go # HTTP  服务器实现代码
│   │   ├── miniblog.go # # miniblog 主业务逻辑实现代码
│   │   ├── model/ # GORM 模型定义
│   │   │   ├── casbin_rule.gen.go # casbin_rule 表对应的 GORM Model 文件
│   │   │   ├── hook.go # GORM Model 钩子
│   │   │   ├── post.gen.go # post 表对应的 GORM Model 文件
│   │   │   └── user.gen.go # user 表对应的 GORM Model 文件
│   │   ├── pkg/ # mb-apsierver 组件级别的共享包
│   │   │   ├── conversion/ # GORM Model 和 Proto 结构体互转代码实现
│   │   │   │   └── conversion.go
│   │   │   └── validation/ # API 接口参数校验
│   │   │       ├── post.go
│   │   │       ├── user.go
│   │   │       ├── validation.go
│   │   │       └── validation_test.go
│   │   ├── store/ # Store 层代码
│   │   │   ├── concrete_post.go
│   │   │   ├── doc.go
│   │   │   ├── logger.go
│   │   │   ├── mock_store.go
│   │   │   ├── post.go
│   │   │   ├── README.md
│   │   │   ├── store.go
│   │   │   └── user.go
│   │   ├── wire_gen.go
│   │   └── wire.go
│   └── pkg/ # miniblog 项目级别的共享包
│       ├── contextx/ # 自定义上下文
│       ├── errno/ # 自定义错误
│       │   ├── code.go # 通用错误类型
│       │   ├── doc.go
│       │   ├── post.go # 博客相关错误类型
│       │   └── user.go # 用户相关错误类型
│       ├── known/ # 常量包
│       │   ├── doc.go
│       │   ├── known.go # 共用常量
│       │   └── role.go # 角色相关常量
│       ├── log/ # 日志包
│       ├── mid/ # 资源唯一 ID 生成包
│       ├── middleware/ # Web 中间件
│       │   ├── gin/ # Gin 框架中间件
│       │   │   ├── authn.go # 认证中间件
│       │   │   ├── authz.go # 授权中间件
│       │   │   ├── bypass.go # 所有请求认证通过中间件
│       │   │   ├── doc.go
│       │   │   ├── header.go # 指定 HTTP Response Header
│       │   │   └── requestid.go # 请求 / 返回头中添加 x-request-id
│       │   └── grpc/ # gRPC 拦截器
│       │       ├── authn.go # 认证中间件
│       │       ├── authz.go # 授权中间件
│       │       ├── bypass.go # 所有请求认证通过中间件
│       │       ├── defaulter.go # 默认值设置拦截器
│       │       ├── doc.go
│       │       ├── requestid.go # 请求 / 返回头中添加 x-request-id
│       │       └── validator.go # 请求参数校验拦截器
│       └── server/ # 服务器实现代码
│           ├── doc.go
│           ├── grpc_server.go # gRPC 服务器
│           ├── http_server.go # HTTP 服务器
│           ├── reverse_proxy_server.go # HTTP 反向代码服务器
│           └── server.go
├── LICENSE # 声明代码所遵循的开源协议
├── Makefile # Makefile 文件，一般大型软件系统都是采用 make 来作为编译工具
├── _output/ # 临时文件存放目录
│   ├── cert/ # CA 文件
│   └── platforms/
│       └── linux/
│           └── amd64/ # 编译后生成的二进制文件
├── pkg/ # 可供外部程序直接使用的 Go 包存放目录
│   └── api/ # Protobuf 文件定义目录
│       └── miniblog/
│           └── v1/
│               ├── example.proto # 示例 Protobuf 文件
│               ├── healthz.proto # 包含了健康检查请求参数和返回参数的定义
│               ├── miniblog.proto # MiniBlog 服务定义
│               ├── post.proto # 包含了博客请求参数和返回参数的定义
│               └── user.proto # 包含了用户请求参数和返回参数的定义
├── README.md # 中文 README
├── scripts/ # 脚本文件
│   ├── boilerplate.txt # 指定版权头信息
│   ├── coverage.awk # awk 脚本，用来计算覆盖率
│   ├── gen_token.sh* # 生成 Token 脚本
│   ├── make-rules/ # 子 Makefile 保存目录
│   │   ├── all.mk
│   │   ├── common.mk # 存放通用的 Makefile 变量
│   │   ├── generate.mk # 用来生成相关代码
│   │   ├── golang.mk # 用来编译源码
│   │   ├── swagger.mk # Swagger 文档规则
│   │   └── tools.mk # 用来完成工具的安装
│   ├── test_nginx.sh* # Nginx 部署测试脚本
│   ├── test_smoke.sh* # 冒烟测试脚本
│   ├── test_tls.sh* # TLS 测试脚本
│   └── wrktest.sh* # API 接口压测工具
├── third_party/  # 第三方 Go 包存放目录、Protobuf 文件
│   └── protobuf/ # 第三方 Protobuf 文件
```