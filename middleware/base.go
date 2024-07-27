package middleware

import (
	g "api_project/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// 定义一些基础中间件

// CORS 跨域请求
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})
}

// WithGormDB 将 gorm.DB 注入到 gin.Context
func WithGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(g.CTX_DB, db)
		ctx.Next()
	}
}
