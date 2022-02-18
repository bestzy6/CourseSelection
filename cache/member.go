package cache

import (
	"ByteDanceCamp8th/model"
	"context"
	"math/rand"
	"strconv"
	"time"
)

// DelMemberByIdinRedis 删除用户的redis缓存,返回row和err
func DelMemberByIdinRedis(member *model.Member) (int, error) {
	userid := "m_" + strconv.Itoa(member.UserID)
	result := RedisClient.Del(context.TODO(), userid)
	return int(result.Val()), result.Err()
}

// AddMemberInRedis 将member添加入redis缓存
func AddMemberInRedis(member *model.Member) error {
	userid := "m_" + strconv.Itoa(member.UserID)
	pipe := RedisClient.Pipeline()
	defer pipe.Close()
	//添加用户缓存至队列中
	pipe.HMSet(context.TODO(), userid, map[string]interface{}{
		"nickname": member.Nickname,
		"username": member.Username,
		"usertype": int(member.UserType),
	})
	//设置过期时间，添加至队列中
	rand.Seed(time.Now().Unix())
	ttl := rand.Intn(5) + 30 //过期时间在30~35分钟之间,防止雪崩。
	pipe.Expire(context.TODO(), userid, time.Duration(ttl)*time.Minute)
	//执行队列
	_, err := pipe.Exec(context.TODO())
	//返回结果
	return err
}

// GetMemberByIDinRedis 在redis缓存中寻找Member，返回的结果是row和err
func GetMemberByIDinRedis(member *model.Member) (row int, err error) {
	key := "m_" + strconv.Itoa(member.UserID)
	result := RedisClient.HGetAll(context.TODO(), key)
	err = result.Err()
	if err != nil {
		return
	}
	//
	val := result.Val()
	row = len(val)
	if row <= 0 {
		return
	}
	//给member赋值
	member.Nickname = val["nickname"]
	member.Username = val["username"]
	usertype, _ := strconv.Atoi(val["usertype"])
	member.UserType = model.UserType(usertype)
	return
}
