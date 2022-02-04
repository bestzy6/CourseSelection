package main

import (
	"ByteDanceCamp8th/server"
	"ByteDanceCamp8th/util/database"
)

func main() {
	r := server.NewRouter()
	r.Run(":8080")
}

//在训练营课程上，字节工程师并不建议使用init函数，然而我就是这么调皮
func init() {
	database.InitMysql()
	database.InitRedis()
}
