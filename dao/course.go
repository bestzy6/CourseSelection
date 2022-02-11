package dao

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/util/database"
	"context"
	"encoding/json"
	"errors"
	"strconv"
)

// GetTeacherCourse 获取老师绑定的所有课程
func GetTeacherCourse(course *model.Course) []model.Course {
	var allCourse []model.Course
	database.DB.Where("teacherid=?", course.TeacherID).Select("courseid", "name").Find(&allCourse)
	return allCourse
}

// UnBindCourse 将课程与教师解绑
func UnBindCourse(course *model.Course) error {
	var courseBind model.Course
	err := database.DB.Where("courseid=?", course.CourseID).Select("teacherid").Find(&courseBind).Error
	if err != nil {
		return err
	}
	//如果教师ID不为0说明已绑定
	if courseBind.TeacherID != 0 {
		err = database.DB.Model(course).Update("teacherid", nil).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("课程未绑定")
	}
}

// BindCourse 将课程与教师绑定
func BindCourse(course *model.Course) error {
	var courseBind model.Course
	err := database.DB.Where("courseid=?", course.CourseID).Select("teacherid").Find(&courseBind).Error
	if err != nil {
		return err
	}
	//如果教师ID为0说明未绑定
	if courseBind.TeacherID == 0 {
		err = database.DB.Model(course).Update("teacherid", course.TeacherID).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("课程已绑定")
	}
}

// CreateCourse 创建课程
func CreateCourse(course *model.Course) (id string, err error) {
	//插入时默认外键为空
	err = database.DB.Debug().Omit("teacherid").Create(course).Error
	id = strconv.Itoa(course.CourseID)
	return
}

// GetCourse 根据课程ID查找课程
func GetCourse(id string) (course model.Course, err error) {
	bys, _ := database.RedisClient.Get(context.TODO(), "course:"+id).Bytes()
	if len(bys) > 0 {
		json.Unmarshal(bys, &course)
	} else {
		result := database.DB.Where("courseid=?", id).Find(&course)
		//使用json序列化
		bytes, _ := json.Marshal(course)
		database.RedisClient.Set(context.TODO(), "course:"+id, bytes, -1)
		if result.Error != nil {
			err = result.Error
		} else if result.RowsAffected < 1 {
			err = errors.New("查询不到该课程")
		}
	}
	return
}
