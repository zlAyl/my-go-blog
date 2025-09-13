# my-go-blog
go 开发的简单blog

# 安装
go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/gin-gonic/gin

go get -u github.com/golang-jwt/jwt/v5

# 项目目录

```
my-go-blog/
├── cmd/                 # 应用程序入口
│   └── server/          # 主服务器入口
│       └── main.go      # 主程序入口文件
├── internal/            # 私有应用程序代码（不对外暴露）
│   ├── config/          # 配置文件处理
│   │   └── database.go
│   ├── models/          # 数据模型/实体（表模型放在这里）
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   │   └── page.go
│   ├── repositories/    # 数据访问层
│   │   ├── user_repository.go
│   │   ├── post_repository.go
│   │   └── comment_repository.go
│   │   └── page_repository.go
│   ├── response/        # 返回结果处理
│   │   └── response.go
│   ├── handlers/        # HTTP 控制器
│   │   ├── user_handler.go
│   │   ├── post_handler.go
│   │   ├── comment_handler.go
│   ├── middleware/      # 自定义中间件
│   │   ├── auth_middleware.go
│   │   └── ...
│   ├── routes/          # 路由
│   │   ├── user_route.go
│   │   ├── post_handler.go
│   │   ├── routres.go
│   │   ├── comment_handler.go

├── go.mod               # Go 模块定义
├── go.sum               # Go 模块校验和
├── .env                 # 环境变量（示例）
├── .env.example         # 环境变量示例
└── README.md            # 项目说明

```
### 用户注册
curl --location --request POST 'http://localhost:8080/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "string",
"password": "string",
"email": "string"
}'

### 用户登录
curl --location --request POST 'http://localhost:8080/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "string",
"password": "string"
}'

### 发布文章
curl --location --request POST 'http://localhost:8080/post/publish' \
--header 'Authorization: xxxx' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "string",
"content": "string"
}'

### 文章列表
curl --location --request GET 'http://localhost:8080/post/list?page=1&page_size=10'

### 更新文章
curl --location --request PATCH 'http://localhost:8080/post/update/1' \
--header 'Authorization: XXXX' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "string",
"content": "string"
}'

### 文章详情
curl --location --request GET 'http://localhost:8080/post/detail/1'

### 删除文章
curl --location --request DELETE 'http://localhost:8080/post/del/1' \
--header 'Authorization: xxxx' \
--header 'Content-Type: application/json' \
--data-raw '{}'

### 发布评论
curl --location --request POST 'http://localhost:8080/comment/publish/1' \
--header 'Authorization: XXXXX' \
--header 'Content-Type: application/json' \
--data-raw '{
"content": "string"
}'

### 评论列表
curl --location --request GET 'http://localhost:8080/comment/list/1' \
--header 'Authorization: XXXX' \
--header 'Content-Type: application/json' \
--data-raw '{}'