package handle

import (
	g "api_project/global"
	"api_project/model"
	"api_project/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserAuth struct{}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (*UserAuth) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db := c.MustGet(g.CTX_DB).(*gorm.DB)
	user, err := model.GetUserByName(db, req.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "用户不存在"})
		return
	}
	// 检查密码是否正确
	if !utils.BcryptCheck(req.Password, user.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "密码错误"})
		return
	}
	// 登陆信息正确 生成token
	conf := g.Conf.JWT
	token, err := utils.GenToken(conf.Secret, conf.Issuer, int(conf.Expire), user.Id, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
