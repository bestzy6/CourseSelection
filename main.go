package main

import (
	"ByteDanceCamp8th/cache"
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/server"
	"ByteDanceCamp8th/util"
)

func main() {
	r := server.NewRouter()
	r.Run(":8080")
}

//在训练营课程上，字节工程师并不建议使用init函数
func init() {
	model.InitMysql()
	cache.InitRedis()
	util.InitMQ()
}
