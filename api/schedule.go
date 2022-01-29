package api

import (
	"ByteDanceCamp8th/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Schedule(c *gin.Context) {
	var sc model.ScheduleCourseRequest
	if err := c.ShouldBind(&sc); err != nil {
		c.JSON(http.StatusOK, model.ScheduleCourseResponse{
			Code: model.ParamInvalid,
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, scheduleCourse(&sc))
}

//求解器(不对外暴露)，暂时使用的是基于dfs的匈牙利算法
func scheduleCourse(scr *model.ScheduleCourseRequest) *model.ScheduleCourseResponse {
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
