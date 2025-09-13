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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(RegUser)
	StoreUser, err := u.userRepo.FindUserByUsername(RegUser.Username)
	if StoreUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户已经注册"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	RegUser.Password = string(hashedPassword)

	var user models.User
	user.Username = RegUser.Username
	user.Password = RegUser.Password
	user.Email = RegUser.Email
	// 创建用户
	if err := u.userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// Login  用户登录
func (u *UserHandler) Login(c *gin.Context) {
	var user models.LoginUser
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	StoreUser, err := u.userRepo.FindUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户名或密码错误"})
		return
	}

	if !u.ValidatePassword(StoreUser.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": signedString,
	})

}

// ValidatePassword 验证密码
func (u *UserHandler) ValidatePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
