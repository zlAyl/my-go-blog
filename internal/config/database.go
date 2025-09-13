package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDataBase() (*gorm.DB, error) {
	dsn := "root:12345677@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("链接数据库失败 : %v", err)
	}
	log.Printf("数据库链接成功")
	return db, nil
}
