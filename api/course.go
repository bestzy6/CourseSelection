package api

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTeacherCourse 获取老师绑定的所有课程
func GetTeacherCourse(c *gin.Context) {
	var req model.GetTeacherCourseRequest
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, model.GetTeacherCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		resp := service.GetTeacherCourseService(&req)
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
		resp := service.UnBindCourseService(&courseReq)
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
		resp := service.BindCourseService(&courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// CreateCourse 添加课程
func CreateCourse(c *gin.Context) {
	var courseReq model.CreateCourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		c.JSON(http.StatusOK, model.CreateCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		resp := service.CreateCourseService(&courseReq)
		c.JSON(http.StatusOK, resp)
	}
}

// GetCourse 查询课程
func GetCourse(c *gin.Context) {
	var req model.GetCourseRequest
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, model.GetCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		courseService := service.GetCourseService(&req)
		c.JSON(http.StatusOK, courseService)
	}
}

// Schedule 课程安排分配
func Schedule(c *gin.Context) {
	var sc model.ScheduleCourseRequest
	if err := c.ShouldBind(&sc); err != nil {
		c.JSON(http.StatusOK, model.ScheduleCourseResponse{
			Code: model.ParamInvalid,
		})
		return
	}
	c.JSON(http.StatusOK, service.ScheduleCourse(&sc))
}
