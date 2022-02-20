package service

import (
	"ByteDanceCamp8th/cache"
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/util"
	"strconv"
)

// GetStudentCourseService 获取学生已选课列表的服务
func GetStudentCourseService(req *model.GetStudentCourseRequest) *model.GetStudentCourseResponse {
	var resp model.GetStudentCourseResponse
	//
	return &resp
}

// ChooseCourseService 学生抢课服务
func ChooseCourseService(req *model.BookCourseRequest) *model.BookCourseResponse {
	var resp model.BookCourseResponse
	//检验参数
	courseid, err := strconv.Atoi(req.CourseID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	studentid, err := strconv.Atoi(req.StudentID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	//绑定模型
	sc := &model.StudentCourse{
		MemberId: studentid,
		CourseId: courseid,
	}
	//检验学生和课程是否存在，以及课程是否已满
	if errNo := checkStuCou(studentid, courseid); errNo != model.OK {
		resp.Code = errNo
		return &resp
	}
	//抢课
	if errNo := killCourse(sc); errNo != model.OK {
		//如果课程已经满则加入map
		if errNo == model.CourseNotAvailable {
			model.CourseFull[sc.CourseId] = true
		}
		resp.Code = errNo
		return &resp
	}
	//加入消息队列
	resp.Code = model.OK
	util.ChooseCourseMQ <- sc
	return &resp
}

//抢课操作
func killCourse(sc *model.StudentCourse) model.ErrNo {
	//判断学生是否已经选择该课程
	ok := cache.StudentHasCourse(sc)
	if ok {
		return model.StudentHasCourse
	}
	//获取课程选课人数余量以及错误、抢课操作
	err := cache.ChooseCourseInRedis(sc)
	switch err.(type) {
	case nil:
		return model.OK
	case cache.ZeroLeftError:
		return model.CourseNotAvailable
	default:
		return model.CourseNotAvailable
	}
}

//检验学生和课程是否存在，以及课程是否已满
func checkStuCou(sid, cid int) model.ErrNo {
	//检验学生是否存在，在本地内存中查询
	if !model.StudentList[sid] {
		return model.StudentNotExisted
	}
	//检验课程是否存在,在缓存中查询
	if cache.GetCourseExist(cid) < 1 {
		return model.CourseNotExisted
	}
	//检验课程是否已满，本地内存中查询
	if model.CourseFull[cid] {
		return model.CourseNotAvailable
	}
	return model.OK
}
