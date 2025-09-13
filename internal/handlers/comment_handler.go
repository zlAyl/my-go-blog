package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/repositories"
	"github.com/zlAyl/my-go-blog/internal/response"
)

type CommentHandler struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentHandler(commentRepo *repositories.CommentRepository) *CommentHandler {
	return &CommentHandler{commentRepo: commentRepo}
}

// Publish 发布评论
func (com *CommentHandler) Publish(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "无效的文章ID ",
		}.Error(c)
		return
	}

	var publishComment models.PublishComment
	if err := c.BindJSON(&publishComment); err != nil {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "参数错误 " + err.Error(),
		}.Error(c)
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		response.BaseResponse{
			Code: http.StatusUnauthorized,
			Msg:  "未授权 ",
		}.Error(c)
		return
	}
	comment := models.Comment{
		Content: publishComment.Content,
		UserID:  userId.(uint),
		PostID:  uint(postId),
	}
	if err := com.commentRepo.PublishComment(&comment); err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "发布评论失败: " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{}.Success(c)
}

func (com *CommentHandler) Lists(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "无效的文章ID",
		}.Error(c)
		return
	}
	comments, err := com.commentRepo.CommentLists(uint(postId))
	if err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取评论失败: " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{
		Data: comments,
	}.Success(c)
}
