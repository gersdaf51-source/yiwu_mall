# yiwu_mall


WSL2 Ubuntu
   ├── Go (手动安装)
   ├── Node.js (nvm)
   ├── Docker Desktop (Windows)
   ├── PostgreSQL (Docker)
   ├── Redis (Docker)





yiwu-mall/
│
├── apps/
│   ├── web/        # Next.js 前端    只负责页面
│   ├── admin/      # 管理后台（可后做）
│   └── api/        # Go Gin 后端    只负责业务
│
├── packages/
│   └── shared/    # 类型/工具（可选）
│
├── infra/
│   └── docker/
│
├── docker-compose.yml
├── .env
└── README.md



细节结构：

yiwu-mall/
└── apps/
    └── api/
        ├── main.go
        ├── go.mod
        │
        ├── config/
        │   └── config.go          # 配置读取
        │
        ├── database/
        │   └── postgres.go        # 数据库连接
        │
        ├── router/
        │   └── router.go          # 路由
        │
        ├── handler/
        │   └── product.go         # 接收HTTP请求
        │
        ├── service/
        │   └── product.go         # 业务逻辑
        │
        ├── repository/
        │   └── product.go         # 数据库操作
        │
        └── model/
            └── product.go          # 数据结构
