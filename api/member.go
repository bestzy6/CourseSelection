package api

import (
	"ByteDanceCamp8th/model"
	"ByteDanceCamp8th/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateMember(c *gin.Context) {
	var req model.CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.CreateMemberResponse{
			Code: model.ParamInvalid,
		})
	} else {
		if service.CheckCreateMemberParamService(&req) {
			//校验通过，进行创建成员操作
			c.JSON(http.StatusOK, service.CreateMemberService(&req))
		} else {
			//校验不通过
			c.JSON(http.StatusOK, model.CreateMemberResponse{
				Code: model.ParamInvalid,
			})
		}
	}
}

func GetMember(c *gin.Context) {
	var req model.GetMemberRequest
	userid, ok := c.GetQuery("UserID")
	if !ok {
		c.JSON(http.StatusOK, model.GetMemberResponse{
			Code: model.ParamInvalid,
		})
	} else {
		req.UserID = userid
		c.JSON(http.StatusOK, service.GetMemberService(&req))
	}
}

func GetMemberList(c *gin.Context) {
	offset, ok := c.GetQuery("Offset")
	if !ok {
		c.JSON(http.StatusOK, model.GetMemberListResponse{
			Code: model.ParamInvalid,
		})
	}
	limit, ok := c.GetQuery("Limit")
	if !ok {
		c.JSON(http.StatusOK, model.GetMemberListResponse{
			Code: model.ParamInvalid,
		})
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusOK, model.GetMemberListResponse{
			Code: model.ParamInvalid,
		})
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusOK, model.GetMemberListResponse{
			Code: model.ParamInvalid,
		})
	}
	req := &model.GetMemberListRequest{
		Limit:  limitInt,
		Offset: offsetInt,
	}
	c.JSON(http.StatusOK, service.GetMemberListService(req))
}

// UpdateMember 更新成员
func UpdateMember(c *gin.Context) {
	var req model.UpdateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.UpdateMemberResponse{
			Code: model.ParamInvalid,
		})
	} else {
		c.JSON(http.StatusOK, service.UpdateMemberService(&req))
	}
}

func DeleteMember(c *gin.Context) {
	var req model.DeleteMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.DeleteMemberResponse{
			Code: model.ParamInvalid,
		})
	} else {
		c.JSON(http.StatusOK, service.DeleteMemberService(&req))
	}
}
