package service

import (
	"ByteDanceCamp8th/cache"
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/util"
	"strconv"
	"strings"
)

// GetStudentCourseService 获取学生已选课列表的服务
func GetStudentCourseService(req *model.GetStudentCourseRequest) *model.GetStudentCourseResponse {
	var resp model.GetStudentCourseResponse
	//验证ID有效性
	sid, err := strconv.Atoi(req.StudentID)
	if err != nil {
		resp.Code = model.ParamInvalid
		return &resp
	}
	//查询学生是否存在
	if _, ok := model.StudentList[sid]; !ok {
		resp.Code = model.StudentNotExisted
		return &resp
	}
	sc := &model.StudentCourse{MemberId: sid}
	//查询学生课程
	courses := cache.StudentCourseInfo(sc)
	//赋值给ans
	ans := make([]model.TCourse, len(courses))
	var c model.Course
	for i, v := range courses {
		//去除掉课程key的前缀c_
		prefix := strings.TrimPrefix(v, "c_")
		c.CourseID, _ = strconv.Atoi(prefix)
		err = cache.GetCourseInRedis(&c)
		if err != nil {
			resp.Code = model.UnknownError
			return &resp
		}
		//
		ans[i] = model.TCourse{
			CourseID:  prefix,
			Name:      c.Name,
			TeacherID: strconv.Itoa(c.TeacherID),
		}
	}
	resp.Code = model.OK
	resp.Data = struct{ CourseList []model.TCourse }{CourseList: ans}
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
		//如果是课程已满错误码，则加入map
		model.CourseFull[sc.CourseId] = true
		return model.CourseNotAvailable
	default:
		return model.CourseNotAvailable
	}
}

//检验学生和课程是否存在，以及课程是否已满
func checkStuCou(sid, cid int) model.ErrNo {
	//检验学生是否存在，在本地内存中查询
	if _, ok := model.StudentList[sid]; !ok {
		return model.StudentNotExisted
	}
	//检验课程是否存在,在缓存中查询
	if cache.GetCourseExist(cid) < 1 {
		return model.CourseNotExisted
	}
	//检验课程是否已满，本地内存中查询，要先判断课程是否存在再判断是否已经满
	if model.CourseFull[cid] {
		return model.CourseNotAvailable
	}
	return model.OK
}
