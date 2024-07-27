package routes

import (
	"api_project/handle"
	"api_project/middleware"
	"github.com/gin-gonic/gin"
)

var (
	// 后台接口
	userAPI handle.UserAuth
	fileAPI handle.Down
)

func RegisterHandlers(r *gin.Engine) {
	r.POST("/login", userAPI.Login)
	r.Use(middleware.JWTAuth())
	{
		r.GET("/download", fileAPI.DownloadFile)
	}
}
