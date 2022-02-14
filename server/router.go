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
	g.GET("/member/find")
	g.POST("/member/create")
	g.GET("/member")
	g.GET("/member/list")
	g.POST("/member/update")
	g.POST("/member/delete")
	// 登录
	g.POST("/auth/login", api.Login)
	g.POST("/auth/logout", api.Logout)
	g.GET("/auth/whoami", api.Whoami)
	// 排课
	g.POST("/course/create", api.CreateCourse)
	g.GET("/course/get", api.GetCourse)
	g.POST("/teacher/bind_course", api.BindCourse)
	g.POST("/teacher/unbind_course", api.UnBindCourse)
	g.GET("/teacher/get_course", api.GetTeacherCourse)
	g.POST("/course/schedule", api.Schedule)
	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")
	return r
}
