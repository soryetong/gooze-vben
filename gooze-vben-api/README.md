# 目录结构说明

```
my-project/
├── api/                      # API 描述文件（如 user.api）
│   └── user.api
│
├── build/                    # 构建相关脚本（如 Dockerfile、CI 脚本）
│   ├── scripts/              # 启动/部署等辅助脚本（如 build.sh）
│   │   └── gen.sh            # 代码生成脚本
│   │   └── start.sh          # 项目启动脚本
│   └── docker/               # Dockerfile 或 compose 文件
│
├── cmd/                      # 程序入口
│   ├── server/               # 服务入口
│   │   └── main.go
│
├── configs/                  # 应用配置文件
│   └── config.yaml           # 主配置文件（可配合 .env 使用）
│
├── docs/                     # 文档入口
│   ├── swagger/              # Swagger 接口文档
│   │   └── user.yaml
│
├── internal/                 # 核心业务代码（推荐不导出，仅项目内部可用）
│   ├── handler/              # 控制器层（接收请求，返回响应）
│   ├── dto/                  # 请求/响应的数据结构
│   ├── router/               # 路由定义
│   ├── service/              # 业务逻辑
│   └── bootstrap/            # 启动逻辑
│
├── pkg/                      # 可复用公共组件（非业务相关）
│   ├── model/                # 通用数据库模型
│
├── static/                   # 静态资源
│   ├── storage/              # 存放临时文件、用户上传文件、缓存等
│   
├── test/                     # 单元测试 / 集成测试代码
│
├── .env                      # 环境变量文件（用于区分本地/测试/生产）
├── .gitignore                # Git 忽略文件
├── go.mod                    # Go 模块定义
├── go.sum                    # Go 依赖校验文件
└── README.md                 # 项目说明文档
```