package model

type Member struct {
	UserID   int      `gorm:"primaryKey;column:userid"` //用户id
	Nickname string   `gorm:"column:nickname"`          //昵称
	Username string   `gorm:"column:username"`          //用户名
	Password string   `gorm:"column:password"`          //密码
	UserType UserType `gorm:"column:usertype"`          //类型（2:学生or1:管理员or3:教师）
	State    bool     `gorm:"column:state"`             //状态，已删除为true，否则为false
}

func (Member) TableName() string {
	return "member"
}

func (member *Member) FindByUsername() (err error) {
	err = db.Where("username=?", member.Username).Find(member).Error
	return
}
