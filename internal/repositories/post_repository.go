package repositories

import (
	"fmt"

	"github.com/zlAyl/my-go-blog/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) PublishPost(post *models.Post) error {
	return repo.db.Create(post).Error
}

func (repo *PostRepository) FindByID(ID uint) (*models.Post, error) {
	var post models.Post
	result := repo.db.Where("id = ?", ID).First(&post)
	return &post, result.Error
}

// GetPostListWithScopes 文章列表
func (repo *PostRepository) GetPostListWithScopes(p models.PostQueryParams) (*models.PageResponse, error) {
	query := repo.db.Model(&models.Post{})
	query = query.Scopes(WithTitleScope(p.Title))

	var posts []models.Post
	var total int64

	query.Count(&total)

	//分页查询
	if err := query.Scopes(PaginateScope(p.Page, p.PageSize)).Find(&posts).Error; err != nil {
		return nil, err
	}

	response := models.PageResponse{
		Total:    total,
		Page:     p.Page,
		PageSize: p.PageSize,
		List:     posts,
	}
	return &response, nil
}

// UpdatePost 更新文章
func (repo *PostRepository) UpdatePost(updatePost models.UpdatePost, postId uint, userId uint) error {
	var post models.Post
	result := repo.db.Where("id = ?", postId).Find(&post)
	if result.Error != nil {
		return fmt.Errorf("获取文章错误: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在")
	}

	if post.UserID != userId {
		return fmt.Errorf("没有更新权限")
	}

	return repo.db.Model(&post).Updates(updatePost).Error
}

func (repo *PostRepository) DeletePost(postId uint, userId uint) error {
	var post models.Post
	result := repo.db.Where("id = ?", postId).Find(&post)
	if result.Error != nil {
		return fmt.Errorf("获取文章错误: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在")
	}

	if post.UserID != userId {
		return fmt.Errorf("没有删除权限")
	}
	return repo.db.Delete(&post).Error
}

func WithTitleScope(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if title != "" {
			return db.Where("title LIKE ?", "%"+title+"%")
		}
		return db
	}
}
