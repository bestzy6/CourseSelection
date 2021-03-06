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

// StudentList 存放学生的映射，用以判断学生是否存在。
var StudentList = make(map[int]struct{})

// CreateIndexList 初始时将学生信息放入映射
func CreateIndexList() {
	var members []Member
	db.Where("usertype=?", Student).Select("userid").Find(&members)
	for i := 0; i < len(members); i++ {
		StudentList[members[i].UserID] = struct{}{}
	}
}

// DeleteMember 删除成员
func (member *Member) DeleteMember() error {
	err := db.Model(member).Update("state", 1).Error
	return err
}

// UpdateMemberNickName 更新成员昵称
func (member *Member) UpdateMemberNickName() error {
	err := db.Model(member).Update("nickname", member.Nickname).Error
	return err
}

// CreateMember 创建用户
func (member *Member) CreateMember() error {
	err := db.Create(member).Error
	return err
}

// FindByUserID 通过用户ID查找用户
func (member *Member) FindByUserID() (int, error) {
	find := db.Select("nickname", "username", "usertype", "state").Find(member)
	return int(find.RowsAffected), find.Error
}

// FindByUsername 通过用户名查找用户
func (member *Member) FindByUsername() (int, error) {
	find := db.Where("username=?", member.Username).Find(member)
	return int(find.RowsAffected), find.Error
}

// GetMemberStateByID 根据用户ID查询用户状态与类型
func (member *Member) GetMemberStateByID() (int, error) {
	result := db.Select("state", "usertype").Find(member)
	return int(result.RowsAffected), result.Error
}

// GetMembers 获取用户列表
func GetMembers(offset, limit int) (*[]Member, error) {
	var members []Member
	result := db.Limit(limit).Offset(offset).Where("state = ?", 0).Select("userid", "nickname", "username", "usertype").Find(&members)
	return &members, result.Error
}

// GetTypeByName 根据username获取usertype
func GetTypeByName(username string) (UserType, int, error) {
	var member Member
	result := db.Where("username=?", username).Select("usertype").Find(&member)
	return member.UserType, int(result.RowsAffected), result.Error
}
