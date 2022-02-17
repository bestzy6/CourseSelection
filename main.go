package main

import (
	"ByteDanceCamp8th/cache"
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/server"
)

func main() {
	r := server.NewRouter()
	r.Run(":80")
}

//在训练营课程上，字节工程师并不建议使用init函数，然而我就是这么调皮
func init() {
	model.InitMysql()
	cache.InitRedis()
}
