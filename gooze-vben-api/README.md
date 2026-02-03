# 前言

**如果你是第一次使用，请先查看文档 [看这里](https://soryetong.github.io/gooze-docs/)**

# 如何使用？

1. 启动项目

    在 `build/scripts` 目录下有一个 `start_*.sh` 脚本，它就是用来启动项目的。

    ```bash
    sh ./build/scripts/start_*.sh # 记得修改脚本名称
    ```

2. 代码生成

    在 `build/scripts` 目录下有一个 `gen_*.sh` 脚本，它就是用来生成代码的。

    ```bash
    sh ./build/scripts/gen_*.sh # 记得修改脚本名称
    ```

3. 要使用 `databases`、`mongo`、`redis` 等，请记得先修改 `configs/*.yaml` 文件中的地址，或者通过 `.env_*` 重写 yaml 的内容

# 目录结构说明

```
my-project/
├── api/                      # API 描述文件目录
│   ├── admin/                # 管理后台端 API 描述
│   │   └── user.api
│   └── client/               # 客户端 API 描述
│       └── user.api
│
├── build/                    # 构建相关脚本（如 Dockerfile、CI 脚本）
│   ├── scripts/              # 启动/部署等辅助脚本（如 build.sh）
│   │   └── gen.sh            # 通用代码生成脚本
│   │   └── gen_*.sh          # 代码生成脚本 *代表对应端
│   │   └── start_*.sh        # 项目启动脚本 *代表对应端
│   └── docker/               # Dockerfile 或 compose 文件
│
├── cmd/                      # 每个服务的启动入口
│   ├── admin/                # 管理后台服务入口
│   │   └── main.go
│   └── client/               # 客户端服务入口
│       └── main.go
│
├── configs/                  # 应用级配置（YAML 格式）
│   ├── rbac_model.conf       # Casbin 权限规则
│   ├── admin.yaml            # admin 端配置文件
│   └── client.yaml           # client 端配置文件
│
├── docs/                     # 文档入口
│   ├── swagger/              # Swagger 文档
│   │   ├── admin/            # admin 接口文档
│   │   │  └── user.yaml
│   │   └── client/           # client 接口文档
│   │      └── user.yaml
│
├── internal/                 # 核心业务代码（按端划分）
│   ├── admin/                # 管理后台模块
│   │   ├── handler/          # 控制器层（接收请求，返回响应）
│   │   ├── dto/              # 请求/响应的数据结构
│   │   ├── router/           # 路由定义
│   │   ├── service/          # 业务逻辑
│   │   └── bootstrap/        # 启动逻辑
│   └── client/               # 客户端模块（结构同 admin）
│
├── pkg/                      # 可复用公共组件（非业务相关）
│   ├── model/                # 通用数据库模型
│
├── static/                   # 静态资源
│   ├── storage/              # 存放临时文件、用户上传文件、缓存等
│
├── test/                     # 测试文件（单测、集成测试等）
│
├── .env.admin                # 管理后台的环境变量文件
├── .env.client               # 客户端的环境变量文件
├── .gitignore                # Git 忽略文件
├── go.mod                    # Go 模块定义
├── go.sum                    # Go 依赖校验文件
└── README.md                 # 项目说明文档

```
