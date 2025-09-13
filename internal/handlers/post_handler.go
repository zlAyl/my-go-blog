package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/repositories"
	"github.com/zlAyl/my-go-blog/internal/response"
)

type PostHandler struct {
	postRepo *repositories.PostRepository
}

func NewPostHandler(postRepo *repositories.PostRepository) *PostHandler {
	return &PostHandler{postRepo: postRepo}
}

// Publish 发布文章
func (postHandler *PostHandler) Publish(c *gin.Context) {
	var publishPost models.PublishPost
	if err := c.ShouldBindJSON(&publishPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求失败 参数错误: " + err.Error()})
		return
	}
	var post models.Post
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	post.Title = publishPost.Title
	post.Content = publishPost.Content
	post.UserID = userId.(uint)
	if err := postHandler.postRepo.PublishPost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章发布失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "文章发布成功"})
}

// List 文章列表
func (postHandler *PostHandler) List(c *gin.Context) {
	var queryParams models.PostQueryParams
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	postLists, err := postHandler.postRepo.GetPostListWithScopes(queryParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postLists, "message": "success"})
}

// Update 更新文章
func (postHandler *PostHandler) Update(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "无效的文章ID",
		}.Error(c)
		return
	}

	var updatePost models.UpdatePost
	if err := c.ShouldBindJSON(&updatePost); err != nil {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "参数错误" + err.Error(),
		}.Error(c)
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		response.BaseResponse{
			Code: http.StatusUnauthorized,
			Msg:  "未授权",
		}.Error(c)
		return
	}
	if err := postHandler.postRepo.UpdatePost(updatePost, uint(postId), userId.(uint)); err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新文章失败: " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{}.Success(c)

}

// Delete 删除文章
func (postHandler *PostHandler) Delete(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "无效的文章ID ",
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
	if err := postHandler.postRepo.DeletePost(uint(postId), userId.(uint)); err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除文章失败 " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{}.Success(c)
}

// Detail 文章详情
func (postHandler *PostHandler) Detail(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil || postId <= 0 {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "无效的文章ID ",
		}.Error(c)
		return
	}

	post, err := postHandler.postRepo.FindByID(uint(postId))
	if err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取文章失败 " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{
		Data: post,
	}.Success(c)
}
