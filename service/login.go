package service

import (
	"ByteDanceCamp8th/model"
	"strconv"
)

// LoginService 登入服务
func LoginService(req *model.LoginRequest) *model.LoginResponse {
	var resp model.LoginResponse
	member := model.Member{
		Username: req.Username,
	}
	err := member.FindByUsername()
	//数据库连接出错，返回数据库错误
	if err != nil {
		resp.Code = model.UnknownError
		resp.Data = struct{ UserID string }{UserID: ""}
		return &resp
	}
	//返回用户不存在错误
	if member.Nickname == "" && member.UserID == 0 {
		resp.Code = model.UserNotExisted
		resp.Data = struct{ UserID string }{UserID: ""}
		return &resp
	}
	//返回用户已删除错误
	if member.State {
		resp.Code = model.UserHasDeleted
		resp.Data = struct{ UserID string }{UserID: ""}
		return &resp
	}
	//返回密码错误
	if member.Password != req.Password {
		resp.Code = model.WrongPassword
		resp.Data = struct{ UserID string }{UserID: ""}
		return &resp
	}
	resp.Code = model.OK
	resp.Data = struct {
		UserID string
	}{UserID: strconv.Itoa(member.UserID)}
	return &resp
}
