package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/config"
	"github.com/zlAyl/my-go-blog/internal/handlers"
	"github.com/zlAyl/my-go-blog/internal/middleware"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/repositories"
	"github.com/zlAyl/my-go-blog/internal/routes"
	"gorm.io/gorm"
)

func main() {
	//初始化gin
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())
	
	//初始化数据库
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	if err := autoMigrate(db); err != nil {
		log.Fatalf("数据库表创建失败: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	postRepo := repositories.NewPostRepository(db)
	postHandler := handlers.NewPostHandler(postRepo)

	commentRepo := repositories.NewCommentRepository(db)
	commentHandler := handlers.NewCommentHandler(commentRepo)

	//设置路由
	userRoute := routes.NewUserRouter(userHandler)
	userRoute.RegisterRoutes(r)

	postRoute := routes.NewPostRouter(postHandler)
	postRoute.RegisterRoutes(r)

	commentRoute := routes.NewCommentRouter(commentHandler)
	commentRoute.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// 初始化数据库
func initDatabase() (*gorm.DB, error) {

	db, err := config.NewDataBase()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 自动创建或更新表结构
func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		return err
	}

	log.Println("数据库表创建成功")
	return nil
}
