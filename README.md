# my-go-blog
go 开发的简单blog

# 安装
go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gin-gonic/gin

go get -u github.com/golang-jwt/jwt/v5

# 项目目录

my-go-blog/
├── cmd/                 # 应用程序入口
│   └── server/          # 主服务器入口
│       └── main.go      # 主程序入口文件
├── internal/            # 私有应用程序代码（不对外暴露）
│   ├── config/          # 配置文件处理
│   │   └── config.go
│   ├── models/          # 数据模型/实体（表模型放在这里）
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   │   └── page.go
│   ├── repositories/    # 数据访问层（DAO）
│   │   ├── user_repository.go
│   │   ├── post_repository.go
│   │   └── comment_repository.go
│   │   └── page_repository.go
│   ├── handlers/        # HTTP 处理器（控制器）
│   │   ├── user_handler.go
│   │   ├── post_handler.go
│   │   ├── comment_handler.go
│   ├── middleware/      # 自定义中间件
│   │   ├── auth_middleware.go
│   │   └── ...
├── go.mod               # Go 模块定义
├── go.sum               # Go 模块校验和
├── .env                 # 环境变量（示例）
├── .env.example         # 环境变量示例
└── README.md            # 项目说明
