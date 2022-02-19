package cache

import (
	"ByteDanceCamp8th/model"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// GetCourseStateInRedis 从缓存中获取课程的绑定信息
func GetCourseStateInRedis(course *model.Course) error {
	courseid := "c_" + strconv.Itoa(course.CourseID)
	result := RedisClient.HGet(context.TODO(), courseid, "teacher_id")
	if result.Err() != nil {
		return result.Err()
	}
	teacherid, _ := result.Int()
	course.TeacherID = teacherid
	return nil
}

// GetCourseInRedis 从缓存中获取数据，返回找到的个数以及错误
func GetCourseInRedis(course *model.Course) error {
	courseid := "c_" + strconv.Itoa(course.CourseID)
	result := RedisClient.HGetAll(context.TODO(), courseid)
	if result.Err() != nil {
		return result.Err()
	}
	//坑！hgetall一个不存在key，会返回空的map{}，不会返回error
	if len(result.Val()) <= 0 {
		return redis.Nil
	}
	//获取name
	course.Name = result.Val()["name"]
	//获取teacher_id
	course.TeacherID, _ = strconv.Atoi(result.Val()["teacher_id"])
	//返回结果
	return nil
}

// UpdateCourseInRedis 绑定或解绑课程,bind为true说明绑定，bind为false说明解绑
func UpdateCourseInRedis(course *model.Course, bind bool) (err error) {
	courseid := "c_" + strconv.Itoa(course.CourseID)
	if bind {
		err = RedisClient.HSet(context.TODO(), courseid, "teacher_id", course.TeacherID).Err()
	} else {
		err = RedisClient.HSet(context.TODO(), courseid, "teacher_id", 0).Err()
	}
	return err
}

// AddCourseInRedis 添加课程至redis缓存中
func AddCourseInRedis(course *model.Course) error {
	courseid := "c_" + strconv.Itoa(course.CourseID)
	err := RedisClient.HMSet(context.TODO(), courseid, map[string]interface{}{
		"name":       course.Name,
		"teacher_id": course.TeacherID,
		"cap_total":  course.CapTotal,
		"cap_used":   course.CapUsed,
	}).Err()
	return err
}
