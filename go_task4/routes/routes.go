package routes

import (
	"github.com/gin-gonic/gin"
	"log"

	"github.com/polarrana/go_task4/controllers"
	"github.com/polarrana/go_task4/middlewares"
	"gorm.io/gorm"
)

// SetupRoutes 设置路由
func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 设置信任的代理
	// 只信任本地回环地址和特定代理
	trustedProxies := []string{
		"127.0.0.1",      // 本地回环
		"192.168.1.0/24", // 本地网络
		"10.0.0.0/8",     // 私有网络
	}

	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatal("设置信任代理失败:", err)
	}

	authController := controllers.NewAuthController(db)
	postController := controllers.NewPostController(db)
	commentController := controllers.NewCommentController(db)

	// 认证路由
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// 文章路由
	posts := r.Group("/posts")
	{
		posts.GET("", postController.GetPosts)
		posts.GET("/:id", postController.GetPost)

		// 需要认证的路由
		posts.Use(middlewares.AuthMiddleware(db))
		posts.POST("", postController.CreatePost)
		posts.PUT("/:id", postController.UpdatePost)
		posts.DELETE("/:id", postController.DeletePost)
	}

	// 评论路由
	comments := r.Group("/comments/:postId")
	{
		comments.GET("", commentController.GetComments)

		// 需要认证的路由
		comments.Use(middlewares.AuthMiddleware(db))
		comments.POST("", commentController.CreateComment)
	}

	return r
}
