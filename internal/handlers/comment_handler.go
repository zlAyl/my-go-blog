package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/repositories"
)

type CommentHandler struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentHandler(commentRepo *repositories.CommentRepository) *CommentHandler {
	return &CommentHandler{commentRepo: commentRepo}
}

func (com *CommentHandler) Publish(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var publishComment models.PublishComment
	if err := c.BindJSON(&publishComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误：" + err.Error()})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	comment := models.Comment{
		Content: publishComment.Content,
		UserID:  userId.(uint),
		PostID:  uint(postId),
	}
	if err := com.commentRepo.PublishComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发布评论失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "评论成功"})
}

func (com *CommentHandler) Lists(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	comments, err := com.commentRepo.CommentLists(uint(postId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments, "message": "成功"})
}
