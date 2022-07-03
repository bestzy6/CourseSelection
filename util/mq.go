package util

import (
	"ByteDanceCamp8th/model"
	"log"
)

var (
	CreateCourseMQ chan *model.Course
	UnBindCourseMQ chan *model.Course
	BindCourseMQ   chan *model.Course
	ChooseCourseMQ chan *model.StudentCourse
)

//通道的缓存最大值
const maxMessageNum = 2000

// InitMQ 初始化消息队列并启动监听
func InitMQ() {
	//初始化消息队列
	CreateCourseMQ = make(chan *model.Course, maxMessageNum)
	BindCourseMQ = make(chan *model.Course, maxMessageNum)
	UnBindCourseMQ = make(chan *model.Course, maxMessageNum)
	ChooseCourseMQ = make(chan *model.StudentCourse, maxMessageNum)
	//启动监听线程
	go listenCreateCourseMQ()
	go listenBindCourseMQ()
	go listenUnBindCourseMQ()
	go listenChooseCourseMQ()
}

//监听抢课的消息队列
func listenChooseCourseMQ() {
	for {
		sc := <-ChooseCourseMQ
		err := sc.SelectCourse()
		if err != nil {
			log.Println("数据库中添加抢课信息失败！", err)
		} else {
			log.Println("数据库中添加抢课信息成功！Info:", sc.MemberId, "-", sc.CourseId)
		}
	}
}

//监听创建课程的消息队列
func listenCreateCourseMQ() {
	for {
		course := <-CreateCourseMQ
		err := course.CreateCourse()
		if err != nil {
			log.Println("数据库中添加课程失败！", err)
		} else {
			log.Println("数据库中添加课程成功！CourseID:", course.CourseID)
		}
	}
}

//监听解绑课程的消息队列
func listenUnBindCourseMQ() {
	for {
		course := <-UnBindCourseMQ
		err := course.UnBindCourse()
		if err != nil {
			log.Println("数据库中解绑课程失败！", err)
		} else {
			log.Println("数据库中解绑课程成功！CourseID:", course.CourseID)
		}
	}
}

//监听绑定课程的消息队列
func listenBindCourseMQ() {
	for {
		course := <-BindCourseMQ
		err := course.BindCourse()
		if err != nil {
			log.Println("数据库中绑定课程失败！", err)
		} else {
			log.Println("数据库中绑定课程成功！CourseID:", course.CourseID)
		}
	}
}
