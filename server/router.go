package server

import (
	"ByteDanceCamp8th/api"
	"ByteDanceCamp8th/server/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Session("secret"))
	r.Use(middleware.Cors())
	// 路由
	g := r.Group("/api/v1")
	// 成员管理
	g.POST("/member/create", api.CreateMember) //创建成员
	g.GET("/member", api.GetMember)            //获取单个成员
	g.GET("/member/list", api.GetMemberList)   //批量获取成员
	g.POST("/member/update", api.UpdateMember) //更新成员
	g.POST("/member/delete", api.DeleteMember) //删除成员
	// 登录
	g.POST("/auth/login", api.Login)   //登入
	g.POST("/auth/logout", api.Logout) //登出
	g.GET("/auth/whoami", api.Whoami)  //获取个人信息
	// 排课
	g.POST("/course/create", api.CreateCourse)         //创建课程
	g.GET("/course/get", api.GetCourse)                //获取课程
	g.POST("/teacher/bind_course", api.BindCourse)     //绑定课程
	g.POST("/teacher/unbind_course", api.UnBindCourse) //解绑课程
	g.GET("/teacher/get_course", api.GetTeacherCourse) //获取老师课程
	g.POST("/course/schedule", api.Schedule)           //排课
	// 抢课
	//g.POST("/student/book_course", api.ChooseCourse) //学生抢课
	//g.GET("/student/course", api.GetStudentCourse)   //获取学生课表
	return r
}
