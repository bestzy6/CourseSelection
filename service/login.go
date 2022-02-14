package service

import (
	"ByteDanceCamp8th/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

// LoginService 登入服务
func LoginService(req *model.LoginRequest) *model.LoginResponse {
	var resp model.LoginResponse
	member := model.Member{
		Username: req.Username,
	}
	rows, err := member.FindByUsername()
	//数据库连接出错，返回数据库错误
	if err != nil {
		resp.Code = model.UnknownError
		return &resp
	}
	//返回用户不存在错误
	if rows <= 0 {
		resp.Code = model.UserNotExisted
		return &resp
	}
	//返回用户已删除错误
	if member.State {
		resp.Code = model.UserHasDeleted
		return &resp
	}
	//返回密码错误
	if member.Password != req.Password {
		resp.Code = model.WrongPassword
		return &resp
	}
	resp.Code = model.OK
	resp.Data = struct {
		UserID string
	}{UserID: strconv.Itoa(member.UserID)}
	return &resp
}

// Auth 权限鉴定
func Auth(c *gin.Context, input model.UserType) bool {
	token, err := c.Cookie("camp-session")
	if err != nil {
		return false
	}
	session := sessions.Default(c)
	username := session.Get(token)
	usertype, rows, err := model.GetTypeByName(username.(string))
	if err != nil || rows <= 0 {
		return false
	}
	if usertype != input {
		return false
	}
	return true
}
