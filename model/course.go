package model

type Course struct {
	CourseID  int    `gorm:"primaryKey;column:courseid"` //课程id
	Name      string `gorm:"column:name"`                //课程名称
	TeacherID int    `gorm:"column:teacherid"`           //授课教师id
	CapTotal  int    `gorm:"column:cap_total"`           //课程容量
	CapUsed   int    `gorm:"column:cap_used"`            //课程已选人数
}

func (Course) TableName() string {
	return "course"
}
