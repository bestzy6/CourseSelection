package main

import (
	"ByteDanceCamp8th/server"
)

func main() {
	r := server.NewRouter()
	r.Run(":8080")
}

//在训练营上，字节工程师并不建议使用init函数
func init() {

}
