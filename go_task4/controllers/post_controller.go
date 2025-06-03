package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/polarrana/go_task4/models"
	"github.com/polarrana/go_task4/utils"
	"gorm.io/gorm"
	"net/http"
)

// PostController 文章控制器
type PostController struct {
	DB *gorm.DB
}

// NewPostController 创建新的文章控制器
func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

// CreatePost 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  currentUser.ID,
	}

	if err := pc.DB.Create(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建文章失败")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "文章创建成功", post)
}

// GetPosts 获取所有文章
func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := pc.DB.Preload("User").Find(&posts).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取文章列表成功", posts)
}

// GetPost 获取单个文章
func (pc *PostController) GetPost(c *gin.Context) {
	var post models.Post
	if err := pc.DB.Preload("User").Preload("Comments.User").First(&post, c.Param("id")).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取文章成功", post)
}

// UpdatePost 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var post models.Post
	if err := pc.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	if post.UserID != currentUser.ID {
		utils.ErrorResponse(c, http.StatusForbidden, "无权更新此文章")
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	updates := make(map[string]interface{})
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Content != "" {
		updates["content"] = input.Content
	}

	if err := pc.DB.Model(&post).Updates(updates).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新文章失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "文章更新成功", post)
}

// DeletePost 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var post models.Post
	if err := pc.DB.First(&post, c.Param("id")).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	if post.UserID != currentUser.ID {
		utils.ErrorResponse(c, http.StatusForbidden, "无权删除此文章")
		return
	}

	if err := pc.DB.Delete(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "删除文章失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "文章删除成功", nil)
}
