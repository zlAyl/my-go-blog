package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Content      string `gorm:"not null"`
	UserID       uint
	User         User      `json:"-"`
	CommentCount uint      `gorm:"default:0"`
	CommentState uint      `gorm:"type:tinyint(1);default:0;comment:评论状态:1已评论0无评论"`
	Comments     []Comment `gorm:"foreignKey:PostID" json:"-"`
}

// PublishPost 发布文章
type PublishPost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PageRequest struct {
	Page     int `form:"page,default=1"`       // 页码，默认为1
	PageSize int `form:"page_size,default=10"` // 每页数量，默认为10
}

type PostQueryParams struct {
	PageRequest
	Title   string `form:"title"`
	Content string `form:"content"`
}
