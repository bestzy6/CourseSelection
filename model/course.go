package model

type Course struct {
	CourseID  int    `gorm:"primaryKey"` //课程id
	Name      string //课程名称
	TeacherID int    `gorm:"column:teacherid"` //授课教师id
	CapTotal  int    //课程容量
	CapUsed   int    //课程已选人数
}

func (Course) TableName() string {
	return "course"
}
