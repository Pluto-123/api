package middleware

import (
	g "api_project/global"
	"api_project/model"
	"api_project/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

// JWTAuth 基于 JWT 的授权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//db := c.MustGet(g.CTX_DB).(*gorm.DB)

		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "用户未登陆",
			})
			return
		}
		// token 的正确格式: `Bearer [tokenString]`
		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token格式错误",
			})
			return
		}

		claims, err := utils.ParseToken(g.Conf.JWT.Secret, parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token 解析失败",
			})
			return
		}

		// 判断 token 已过期
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token 已过期 请重新登陆",
			})
			return
		}
		//
		db := c.MustGet(g.CTX_DB).(*gorm.DB)
		user, err := model.GetUserById(db, claims.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "数据库查询失败 请重新登陆",
			})
		}
		c.Set(g.CTX_USER_AUTH, user)
	}
}
