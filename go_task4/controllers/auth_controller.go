package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/polarrana/go_task4/config"
	"github.com/polarrana/go_task4/models"
	"github.com/polarrana/go_task4/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// AuthController 认证控制器
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController 创建新的认证控制器
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

// Register 用户注册
func (ac *AuthController) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	// 检查用户名或邮箱是否已存在
	var existingUser models.User
	if err := ac.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "用户名或邮箱已被使用")
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Verified: false,
	}

	if err := user.HashPassword(input.Password); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "密码加密失败")
		return
	}

	if err := ac.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "用户注册成功", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	var user models.User
	if err := ac.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 创建JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成令牌失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "登录成功", gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
