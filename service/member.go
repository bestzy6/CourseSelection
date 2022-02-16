package service

import (
	"ByteDanceCamp8th/model"
	"regexp"
	"strconv"
)

// DeleteMemberService 删除成员服务
func DeleteMemberService(req *model.DeleteMemberRequest) *model.DeleteMemberResponse {
	var resp model.DeleteMemberResponse
	userid, err := strconv.Atoi(req.UserID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	member := model.Member{
		UserID: userid,
	}
	row, err := member.FindByUserID()
	//数据库错误
	if err != nil {
		resp.Code = model.UnknownError
		return &resp
	}
	//没有找到该成员
	if row <= 0 {
		resp.Code = model.UserNotExisted
		return &resp
	}
	//成员已经删除
	if member.State {
		resp.Code = model.UserHasDeleted
		return &resp
	}
	err = member.DeleteMember()
	if err != nil {
		resp.Code = model.UnknownError
	}
	resp.Code = model.OK
	return &resp
}

// UpdateMemberService 更新成员服务
func UpdateMemberService(req *model.UpdateMemberRequest) *model.UpdateMemberResponse {
	var resp model.UpdateMemberResponse
	userid, err := strconv.Atoi(req.UserID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	member := model.Member{
		UserID:   userid,
		Nickname: req.Nickname,
	}
	err = member.UpdateMemberNickName()
	if err != nil {
		resp.Code = model.UnknownError
	}
	resp.Code = model.OK
	return &resp
}

// CreateMemberService 创建成员服务
func CreateMemberService(req *model.CreateMemberRequest) *model.CreateMemberResponse {
	var resp model.CreateMemberResponse
	member := &model.Member{
		Nickname: req.Nickname,
		Username: req.Username,
		Password: req.Password,
		UserType: req.UserType,
	}
	err := member.CreateMember()
	if err != nil {
		//如果出现错误返回用户已经存在，因为可能被软删除了
		resp.Code = model.UserHasExisted
	} else {
		resp.Code = model.OK
		resp.Data = struct {
			UserID string
		}{UserID: strconv.Itoa(member.UserID)}
	}
	return &resp
}

// CheckCreateMemberParamService 创建成员时，对参数进行校验的服务,
//用户昵称，必填，不小于 4 位，不超过 20 位（字节）
//用户名，必填，支持大小写，不小于 8 位 不超过 20 位（字节）
//密码，必填，同时包括大小写、数字，不少于 8 位 不超过 20 位（字节）
func CheckCreateMemberParamService(request *model.CreateMemberRequest) bool {
	nickname := request.Nickname
	username := request.Username
	password := request.Password
	//昵称不小于 4 位，不超过 20 位（字节）
	if len(nickname) < 4 || len(nickname) > 20 {
		return false
	}
	//用户名不小于 8 位 不超过 20 位
	if len(username) < 8 || len(username) > 20 {
		return false
	}
	//密码不少于 8 位 不超过 20 位
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	//用户名是否为大小写
	for i := range username {
		if (username[i] >= 'a' && username[i] <= 'z') || (username[i] >= 'A' && username[i] <= 'Z') {
			continue
		} else {
			return false
		}
	}
	//密码是否同时包括大、小写和数字，利用正则表达式
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`}
	for _, pattern := range patternList {
		matchOK, _ := regexp.MatchString(pattern, password)
		if !matchOK {
			return false
		}
	}
	return true
}

// GetMemberService 获取成员服务
func GetMemberService(req *model.GetMemberRequest) *model.GetMemberResponse {
	var resp model.GetMemberResponse
	userid, err := strconv.Atoi(req.UserID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	member := &model.Member{
		UserID: userid,
	}
	rows, err := member.FindByUserID()
	//数据库出现错误
	if err != nil {
		resp.Code = model.UnknownError
		return &resp
	}
	//没有找到状态
	if rows <= 0 {
		resp.Code = model.UserNotExisted
		return &resp
	}
	//成员状态为已删除
	if member.State {
		resp.Code = model.UserHasDeleted
		return &resp
	}
	resp.Code = model.OK
	resp.Data = model.TMember{
		UserID:   req.UserID,
		Nickname: member.Nickname,
		Username: member.Username,
		UserType: member.UserType,
	}
	return &resp
}

func GetMemberListService(req *model.GetMemberListRequest) *model.GetMemberListResponse {
	var resp model.GetMemberListResponse
	members, err := model.GetMembers(req.Offset, req.Limit)
	if err != nil {
		resp.Code = model.UnknownError
	} else {
		resp.Code = model.OK
		ans := make([]model.TMember, 0, len(*members))
		for _, v := range *members {
			ans = append(ans, model.TMember{
				UserID:   strconv.Itoa(v.UserID),
				Nickname: v.Nickname,
				Username: v.Username,
				UserType: v.UserType,
			})
		}
		resp.Data = struct {
			MemberList []model.TMember
		}{MemberList: ans}
	}
	return &resp
}

// checkTeacher 校验Teacher信息
func checkTeacher(teacher *model.Member) model.ErrNo {
	//对teacher进行检验
	row, err := teacher.GetMemberStateByID()
	//数据库错误
	if err != nil {
		return model.UnknownError
	}
	//没有找到数据或者找到的数据不是老师
	if row <= 0 || teacher.UserType != model.Teacher {
		return model.UserNotExisted
	}
	return model.OK
}
