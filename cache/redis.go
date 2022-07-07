package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

//以下是配置Redis数据库
var (
	addrRedis = "124.70.23.171:6379" //redis地址
	pwd       = ""                   //redis密码
	dbnum     = 0                    //redis数据库编号
)

var RedisClient *redis.Client

// InitRedis 初始化Redis链接
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:       addrRedis,
		Password:   pwd,
		DB:         dbnum,
		MaxRetries: 1,
		PoolSize:   100,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalln("link redis failed:", err)
	}
	RedisClient = client
}

// GetNewId 从redis中获取一个自增的ID,idType为键的后缀名
func GetNewId(idType string) (int, error) {
	key := "id_" + idType
	//id_key变量作为存储的kv对的key,如果变量不存在，设置id_key值为1并返回,如果变量存在，值加1后返回
	luaId := redis.NewScript(`
		local id_key = KEYS[1]
		local current = redis.call('get',id_key)
		if current == false then
    		redis.call('set',id_key,1)
   	 		return '1'
		end
		--redis.log(redis.LOG_NOTICE,' current:'..current..':')
		local result =  tonumber(current)+1
		--redis.log(redis.LOG_NOTICE,' result:'..result..':')
		redis.call('set',id_key,result)
		return tostring(result)
	`)
	//执行lua脚本
	n, err := luaId.Run(context.TODO(), RedisClient, []string{key}, 2).Result()
	if err != nil {
		return -1, err
	}
	//
	retint, err := strconv.Atoi(n.(string))
	if err != nil {
		return -1, err
	}
	return retint, nil
}
