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
