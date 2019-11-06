package service

import (
	"Cloudreve/serializer"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	//TODO 细致调整验证规则
	UserName    string `form:"userName" json:"userName" binding:"required,min=5,max=30"`
	Password    string `form:"Password" json:"Password" binding:"required,min=8,max=40"`
	CaptchaCode string `form:"captchaCode" json:"captchaCode"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	return serializer.Response{
		Code: 0,
		Msg:  "OK",
	}
}
