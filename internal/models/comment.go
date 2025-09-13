package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	Post    Post
	PostID  uint
	User    User
	UserID  uint
}

type PublishComment struct {
	Content string `json:"content" binding:"required"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	var post Post
	if err := tx.Where("id = ?", comment.PostID).Find(&post).Error; err != nil {
		return err
	}
	post.CommentCount += 1
	post.CommentState = 1
	return tx.Save(&post).Error
}
