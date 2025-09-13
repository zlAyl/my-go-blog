package main

import (
	"log"

	"github.com/zlAyl/my-go-blog/internal/config"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/routes"
	"gorm.io/gorm"
)

func main() {
	//初始化数据库
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	if err := autoMigrate(db); err != nil {
		log.Fatalf("数据库表创建失败: %v", err)
	}

	//注册路由
	r := routes.RegisterAllRoutes(db)

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
