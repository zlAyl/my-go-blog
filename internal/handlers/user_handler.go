package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zlAyl/my-go-blog/internal/models"
	"github.com/zlAyl/my-go-blog/internal/repositories"
	"github.com/zlAyl/my-go-blog/internal/response"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userPro *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userPro}
}

// Register 用户注册
func (u *UserHandler) Register(c *gin.Context) {
	var RegUser models.RegisterUser
	if err := c.ShouldBindWith(&RegUser, binding.JSON); err != nil {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: " + err.Error(),
		}.Error(c)
		return
	}
	log.Println(RegUser)
	StoreUser, err := u.userRepo.FindUserByUsername(RegUser.Username)
	if StoreUser != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "用户已经注册",
		}.Error(c)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegUser.Password), bcrypt.DefaultCost)
	if err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "密码加密失败",
		}.Error(c)
		return
	}
	RegUser.Password = string(hashedPassword)

	var user models.User
	user.Username = RegUser.Username
	user.Password = RegUser.Password
	user.Email = RegUser.Email
	// 创建用户
	if err := u.userRepo.CreateUser(&user); err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "创建用户失败: " + err.Error(),
		}.Error(c)
		return
	}
	response.BaseResponse{}.Success(c)
}

// Login  用户登录
func (u *UserHandler) Login(c *gin.Context) {
	var user models.LoginUser
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		response.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  "参数错误:" + err.Error(),
		}.Error(c)
		return
	}

	StoreUser, err := u.userRepo.FindUserByUsername(user.Username)
	if err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "用户名或密码错误:" + err.Error(),
		}.Error(c)
		return
	}

	if !u.ValidatePassword(StoreUser.Password, user.Password) {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "用户名或密码错误",
		}.Error(c)
		return
	}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       StoreUser.ID,
		"username": StoreUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	signedString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response.BaseResponse{
			Code: http.StatusInternalServerError,
			Msg:  "生成token失败",
		}.Error(c)
		return
	}
	response.BaseResponse{
		Data: map[string]string{"token": signedString},
	}.Success(c)

}

// ValidatePassword 验证密码
func (u *UserHandler) ValidatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
