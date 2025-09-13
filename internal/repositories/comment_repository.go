package repositories

import (
	"fmt"

	"github.com/zlAyl/my-go-blog/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (com *CommentRepository) PublishComment(comment *models.Comment) error {
	var post models.Post
	result := com.db.Where("id = ?", comment.PostID).First(&post)
	if result.Error != nil {
		return fmt.Errorf("获取文章失败  %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在")
	}
	return com.db.Create(&comment).Error
}

func (com *CommentRepository) CommentLists(postId uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := com.db.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
