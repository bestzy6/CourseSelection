package service

import (
	"ByteDanceCamp8th/model"
	"strconv"
)

// GetTeacherCourseService 获取老师所绑定课程的服务
func GetTeacherCourseService(req *model.GetTeacherCourseRequest) *model.GetTeacherCourseResponse {
	var resp model.GetTeacherCourseResponse
	teacherId, _ := strconv.Atoi(req.TeacherID)
	course := &model.Course{
		TeacherID: teacherId,
	}
	teacherCourse := course.GetTeacherCourse()
	ans := make([]*model.TCourse, 0, len(teacherCourse))
	for _, v := range teacherCourse {
		thisCourse := &model.TCourse{
			CourseID:  strconv.Itoa(v.CourseID),
			Name:      v.Name,
			TeacherID: req.TeacherID,
		}
		ans = append(ans, thisCourse)
	}
	resp.Code = model.OK
	resp.Data = struct {
		CourseList []*model.TCourse
	}{CourseList: ans}
	return &resp
}

// UnBindCourseService 解绑课程的服务
func UnBindCourseService(req *model.UnbindCourseRequest) *model.UnbindCourseResponse {
	var resp model.UnbindCourseResponse
	courseid, _ := strconv.Atoi(req.CourseID)
	teacherid, _ := strconv.Atoi(req.TeacherID)
	course := &model.Course{
		CourseID:  courseid,
		TeacherID: teacherid,
	}
	if err := course.UnBindCourse(); err != nil {
		resp.Code = model.CourseNotBind
	} else {
		resp.Code = model.OK
	}
	return &resp
}

// BindCourseService 绑定课程与教师服务
func BindCourseService(req *model.BindCourseRequest) *model.BindCourseResponse {
	var resp model.BindCourseResponse
	courseid, _ := strconv.Atoi(req.CourseID)
	teacherid, _ := strconv.Atoi(req.TeacherID)
	course := &model.Course{
		CourseID:  courseid,
		TeacherID: teacherid,
	}
	if err := course.BindCourse(); err != nil {
		resp.Code = model.CourseHasBound
	} else {
		resp.Code = model.OK
	}
	return &resp
}

// CreateCourseService 创建课程的服务
func CreateCourseService(req *model.CreateCourseRequest) *model.CreateCourseResponse {
	var resp model.CreateCourseResponse
	course := &model.Course{
		Name:     req.Name,
		CapTotal: req.Cap,
	}
	err := course.CreateCourse()
	if err != nil {
		resp.Code = model.UnknownError
		return &resp
	} else {
		resp.Code = model.OK
		resp.Data = struct {
			CourseID string
		}{strconv.Itoa(course.CourseID)}
		return &resp
	}
}

// GetCourseService 查询课程的服务
func GetCourseService(req *model.GetCourseRequest) *model.GetCourseResponse {
	var resp model.GetCourseResponse
	courseid, _ := strconv.Atoi(req.CourseID)
	course := &model.Course{
		CourseID: courseid,
	}
	err := course.GetCourse()
	if err != nil {
		resp.Code = model.CourseNotExisted
		return &resp
	} else {
		resp.Code = model.OK
		resp.Data = model.TCourse{
			CourseID:  strconv.Itoa(course.CourseID),
			Name:      course.Name,
			TeacherID: strconv.Itoa(course.TeacherID),
		}
		return &resp
	}
}

// ScheduleCourse 求解器，使用的是匈牙利算法
func ScheduleCourse(scr *model.ScheduleCourseRequest) *model.ScheduleCourseResponse {
	ans := new(model.ScheduleCourseResponse)
	ans.Code = model.OK
	//ans.Data表示老师与课程的关系
	ans.Data = make(map[string]string, len(scr.TeacherCourseRelationShip))
	//course表示该课程是否被选择
	course := make(map[string]string, len(scr.TeacherCourseRelationShip))
	//visited表示当前一轮寻找访问过的课程
	var visited map[string]bool
	//寻找是否选到了课
	var find func(teacherid string) bool
	find = func(teacherid string) bool {
		//cid表示课程id
		for _, cid := range scr.TeacherCourseRelationShip[teacherid] {
			if !visited[cid] {
				visited[cid] = true
				if _, isSelected := course[cid]; !isSelected || find(course[cid]) {
					course[cid] = teacherid
					ans.Data[teacherid] = cid
					return true
				}
			}
		}
		return false
	}
	//寻找增广路径
	for k := range scr.TeacherCourseRelationShip {
		visited = make(map[string]bool, len(scr.TeacherCourseRelationShip))
		find(k)
	}

	if len(ans.Data) < len(scr.TeacherCourseRelationShip) {
		ans.Code = model.ParamInvalid
	}
	return ans
}
