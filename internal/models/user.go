package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}

type RegisterUser struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
	Email    string `binding:"required" json:"email"`
}

type LoginUser struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
