package api

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Login 登入api
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.LoginResponse{
			Code: model.ParamInvalid,
		})
	} else {
		resp := service.LoginService(&req)
		//登入成功
		if resp.Code == model.OK {
			session := sessions.Default(c)
			//使用uuid作为token
			//token := uuid.NewV4().String()
			session.Clear()
			session.Set("user_name", req.Username)
			session.Save()
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Logout 登出api
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user_name")
	if username == nil {
		c.JSON(http.StatusOK, model.LogoutResponse{
			Code: model.LoginRequired,
		})
	} else {
		//清除session
		s := sessions.Default(c)
		s.Clear()
		s.Save()
		c.JSON(http.StatusOK, model.LogoutResponse{
			Code: model.OK,
		})
	}
}

// Whoami 是get方法
func Whoami(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user_name")
	//无法获取到token，用户未登录
	if username == nil {
		c.JSON(http.StatusOK, model.WhoAmIResponse{
			Code: model.LoginRequired,
		})
	} else {
		member := &model.Member{
			Username: username.(string),
		}
		_, err := member.FindByUsername()
		//返回数据库访问错误
		if err != nil {
			c.JSON(http.StatusOK, model.WhoAmIResponse{Code: model.UnknownError})
			return
		}
		//假如session查询不到结果，可能存在cookie造假
		if member.Nickname == "" && member.UserID == 0 {
			c.JSON(http.StatusOK, model.WhoAmIResponse{Code: model.ParamInvalid})
			return
		}
		//成功返回信息
		c.JSON(http.StatusOK, model.WhoAmIResponse{
			Code: model.OK,
			Data: struct {
				UserID   string
				Nickname string
				Username string
				UserType model.UserType
			}{
				UserID:   strconv.Itoa(member.UserID),
				Nickname: member.Nickname,
				Username: member.Username,
				UserType: member.UserType},
		})
	}
}
