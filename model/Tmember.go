package model

type TMember struct {
	UserID   string   //用户id
	Nickname string   //昵称
	Username string   //用户名
	UserType UserType //类型（学生or管理员or教师）
}
