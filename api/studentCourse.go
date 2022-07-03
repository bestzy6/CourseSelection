package api

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStudentCourse 获取学生课表api
func GetStudentCourse(c *gin.Context) {
	var req model.GetStudentCourseRequest
	// 处理无效参数情况
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, model.GetStudentCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		c.JSON(http.StatusOK, service.GetStudentCourseService(&req))
	}
}

// ChooseCourse  学生抢课api
func ChooseCourse(c *gin.Context) {
	var req model.BookCourseRequest
	// 处理无效参数情况
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.BookCourseResponse{
			Code: model.ParamInvalid,
		})
	} else {
		c.JSON(http.StatusOK, service.ChooseCourseService(&req))
	}
}
