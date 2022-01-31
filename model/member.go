package model

type Member struct {
	UserID   int      `gorm:"primaryKey;column:userid"` //用户id
	Nickname string   //昵称
	Username string   //用户名
	Password string   //密码
	UserType UserType `gorm:"column:usertype"` //类型（学生or管理员or教师）
	State    bool     //状态，已删除为true，否则为false
}

func (Member) TableName() string {
	return "member"
}
