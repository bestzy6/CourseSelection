package model

type StudentCourse struct {
	MemberId int `gorm:"column:memberid"`
	CourseId int `gorm:"column:courseid"`
}

func (StudentCourse) TableName() string {
	return "student_course"
}

//抢课

//根据查课
//func
