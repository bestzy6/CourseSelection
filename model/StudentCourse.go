package model

type StudentCourse struct {
	MemberId int
	CourseId int
}

func (StudentCourse) TableName() string {
	return "student_course"
}
