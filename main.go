package main

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/util/database"
	"fmt"
)

func main() {
	//r := server.NewRouter()
	//r.Run(":8080")
	//以下为测试部分
	//c := model.Course{
	//	Name:      "math",
	//	TeacherID: 1,
	//	CapTotal:  50,
	//	CapUsed:   0,
	//}
	//err := database.DB.Create(&c).Error
	//if err != nil {
	//	log.Fatalln("err:", err)
	//}
	//
	//var m []model.Member
	//err := database.DB.Where("username=?", "2").Find(&m).Error
	//fmt.Println(m)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//m := model.Member{
	//	Nickname: "666",
	//	Username: "1234",
	//	Password: "123",
	//	UserType: model.Teacher,
	//	State:    false,
	//}
	//err := database.DB.Create(&m).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	var m model.Member
	err := database.DB.Where("username=?", "123").Find(&m).Error
	fmt.Println(m)
	if err != nil {
		fmt.Println(err)
	}
}

//在训练营课程上，字节工程师并不建议使用init函数，然而我就是这么调皮
func init() {
	database.InitMysql()
	database.InitRedis()
}
