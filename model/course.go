package model

import (
	"errors"
)

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

// GetTeacherCourse 获取老师绑定的所有课程
func (course *Course) GetTeacherCourse() []Course {
	var allCourse []Course
	db.Where("teacherid=?", course.TeacherID).Select("courseid", "name").Find(&allCourse)
	return allCourse
}

// UnBindCourse 将课程与教师解绑
func (course *Course) UnBindCourse() error {
	var courseBind Course
	err := db.Where("courseid=?", course.CourseID).Select("teacherid").Find(&courseBind).Error
	if err != nil {
		return err
	}
	//如果教师ID不为0说明已绑定
	if courseBind.TeacherID != 0 {
		err = db.Model(course).Update("teacherid", nil).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("课程未绑定")
	}
}

// BindCourse 将课程与教师绑定
func (course *Course) BindCourse() error {
	var courseBind Course
	err := db.Where("courseid=?", course.CourseID).Select("teacherid").Find(&courseBind).Error
	if err != nil {
		return err
	}
	//如果教师ID为0说明未绑定
	if courseBind.TeacherID == 0 {
		err = db.Model(course).Update("teacherid", course.TeacherID).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("课程已绑定")
	}
}

// CreateCourse 创建课程
func (course *Course) CreateCourse() (err error) {
	//插入时默认外键为空
	err = db.Debug().Omit("teacherid").Create(course).Error
	return
}

// GetCourse 根据课程ID查找课程
func (course *Course) GetCourse() (err error) {
	result := db.Where("courseid=?", course.CourseID).Find(course)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected < 1 {
		err = errors.New("查询不到该课程")
	}
	return
}
