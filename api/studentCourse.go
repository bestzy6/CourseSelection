package api

// GetStudentCourse 获取学生课表
//func GetStudentCourse(c *gin.Context) {
//	var req model.GetStudentCourseRequest
//	// 处理无效参数情况
//	if err := c.BindQuery(&req); err != nil {
//		c.JSON(http.StatusOK, model.GetStudentCourseResponse{
//			Code: model.ParamInvalid,
//		})
//		return
//	}
//	c.JSON(http.StatusOK, service.GetStudentCourseService(&req))
//}
//
//// ChooseCourse 学生抢课
//func ChooseCourse(c *gin.Context) {
//	var req model.BookCourseRequest
//	// 处理无效参数情况
//	if err := c.BindJSON(&req); err != nil {
//		c.JSON(http.StatusOK, model.BookCourseResponse{
//			Code: model.ParamInvalid,
//		})
//		return
//	}
//	c.JSON(http.StatusOK, service.ChooseCourseService(&req))
//}
