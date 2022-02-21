package cache

import (
	"ByteDanceCamp8th/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// ZeroLeftError 库存为0错误
type ZeroLeftError struct {
}

func (e ZeroLeftError) Error() string {
	return "ZeroLeft"
}

// HasBindCourseError 已经绑定过错误
type HasBindCourseError struct {
}

func (e HasBindCourseError) Error() string {
	return "HasBindCourseError"
}

// StudentCourseInfo 获取学生的选课信息
func StudentCourseInfo(sc *model.StudentCourse) []string {
	sid := "s_" + strconv.Itoa(sc.MemberId)
	return RedisClient.SMembers(context.TODO(), sid).Val()
}

// StudentHasCourse 判断学生是否选择了该课程
func StudentHasCourse(sc *model.StudentCourse) bool {
	sid, cid := "s_"+strconv.Itoa(sc.MemberId), "c_"+strconv.Itoa(sc.CourseId)
	return RedisClient.SIsMember(context.TODO(), sid, cid).Val()
}

// ChooseCourseInRedis 将抢课结果写入redis
func ChooseCourseInRedis(sc *model.StudentCourse) error {
	sid, cid := "s_"+strconv.Itoa(sc.MemberId), "c_"+strconv.Itoa(sc.CourseId)
	//事务方法
	txf := func(tx *redis.Tx) error {
		capLeft, err := tx.HGet(context.TODO(), cid, "cap_left").Int()
		if err != nil {
			return err
		}
		if capLeft <= 0 {
			return ZeroLeftError{}
		}
		// 乐观锁，库存减一
		_, err = tx.TxPipelined(context.TODO(), func(pipe redis.Pipeliner) error {
			pipe.HSet(context.TODO(), cid, "cap_left", capLeft-1)
			pipe.SAdd(context.TODO(), sid, cid)
			return nil
		})
		return err
	}
	//循环三次，避免留有课程没被抢
	var err error
	for i := 0; i < 3; i++ {
		err = RedisClient.Watch(context.TODO(), txf, cid)
		//没有错误，直接返回，事务错误，多试几次
		if err == nil {
			return nil
		} else if errors.Is(err, redis.TxFailedErr) {
			continue
		} else if errors.Is(err, ZeroLeftError{}) {
			fmt.Println("课程已满")
			return err
		}
		fmt.Println("抢课失败！")
		return err
	}
	return err
}
