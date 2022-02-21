package model

import (
	"gorm.io/gorm"
)

type StudentCourse struct {
	MemberId int `gorm:"column:memberid"`
	CourseId int `gorm:"column:courseid"`
}

func (StudentCourse) TableName() string {
	return "student_course"
}

// SelectCourse 写入数据
func (sc *StudentCourse) SelectCourse() error {
	err := db.Transaction(func(tx *gorm.DB) error {
		//创建抢课记录
		err := tx.Create(sc).Error
		if err != nil {
			return err
		}
		//抢课余数减一。分两段，代码太长了
		t := tx.Table("course").Where("courseid=?", sc.CourseId)
		err = t.UpdateColumn("cap_left", gorm.Expr("cap_left - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
