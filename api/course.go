package api

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTeacherCourse 获取老师绑定的所有课程
func GetTeacherCourse(c *gin.Context) {
	teacherId, ok := c.GetQuery("TeacherID")
	if !ok {
		c.JSON(http.StatusOK, model.GetTeacherCourseResponse{
			Code: model.ParamInvalid,
			Data: struct {
				CourseList []*model.TCourse
			}{CourseList: nil},
		})
	} else {
		var courseReq model.GetTeacherCourseRequest
		courseReq.TeacherID = teacherId
		resp := service.GetTeacherCourseService(courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// UnBindCourse 教师解绑课程
func UnBindCourse(c *gin.Context) {
	var courseReq model.UnbindCourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		c.JSON(http.StatusOK, model.UnbindCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		resp := service.UnBindCourseService(courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// BindCourse 教师绑定课程
func BindCourse(c *gin.Context) {
	var courseReq model.BindCourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		c.JSON(http.StatusOK, model.BindCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		resp := service.BindCourseService(courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// CreateCourse 添加课程
func CreateCourse(c *gin.Context) {
	var courseReq model.CreateCourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		c.JSON(http.StatusOK, model.CreateCourseResponse{
			Code: model.ParamInvalid,
			Data: struct {
				CourseID string
			}{CourseID: "-1"},
		})
	} else {
		resp := service.CreateCourseService(courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// GetCourse 查询课程
func GetCourse(c *gin.Context) {
	courseId, ok := c.GetQuery("CourseID")
	if !ok {
		c.JSON(http.StatusOK, model.GetCourseResponse{
			Code: model.ParamInvalid,
			Data: model.TCourse{},
		})
	} else {
		var courseReq model.GetCourseRequest
		courseReq.CourseID = courseId
		courseService := service.GetCourseService(courseReq)
		c.JSON(http.StatusOK, courseService)
	}
}

// Schedule 课程安排分配
func Schedule(c *gin.Context) {
	var sc model.ScheduleCourseRequest
	if err := c.ShouldBind(&sc); err != nil {
		c.JSON(http.StatusOK, model.ScheduleCourseResponse{
			Code: model.ParamInvalid,
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, service.ScheduleCourse(&sc))
}
