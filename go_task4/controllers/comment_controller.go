package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/polarrana/go_task4/models"
	"github.com/polarrana/go_task4/utils"
	"gorm.io/gorm"
	"net/http"
)

// CommentController 评论控制器
type CommentController struct {
	DB *gorm.DB
}

// NewCommentController 创建新的评论控制器
func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{DB: db}
}

// CreateComment 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	// 检查文章是否存在
	var post models.Post
	if err := cc.DB.First(&post, c.Param("postId")).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  currentUser.ID,
		PostID:  post.ID,
	}

	if err := cc.DB.Create(&comment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建评论失败")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "评论创建成功", comment)
}

// GetComments 获取文章的所有评论
func (cc *CommentController) GetComments(c *gin.Context) {
	var comments []models.Comment
	if err := cc.DB.Preload("User").Where("post_id = ?", c.Param("postId")).Find(&comments).Error; err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取评论列表成功", comments)
}
