package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDatabaseError 处理数据库错误
func HandleDatabaseError(c *gin.Context, err error) {
	if err == gorm.ErrRecordNotFound {
		ErrorResponse(c, http.StatusNotFound, "记录不存在")
		return
	}
	log.Println("Database error:", err)
	ErrorResponse(c, http.StatusInternalServerError, "数据库操作失败")
}

// HandleValidationError 处理验证错误
func HandleValidationError(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusBadRequest, err.Error())
}
