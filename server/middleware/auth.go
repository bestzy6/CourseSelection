package middleware

import (
	"ByteDanceCamp8th/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthUser 判断用户权限,input表示为应该对应的权限
func AuthUser(input model.UserType) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("user_name")
		if username == nil {
			c.JSON(http.StatusOK, &model.ResponseMeta{Code: model.PermDenied})
			c.Abort()
			return
		}
		usertype, rows, err := model.GetTypeByName(username.(string))
		if err != nil || rows <= 0 || usertype != input {
			c.JSON(http.StatusOK, &model.ResponseMeta{Code: model.PermDenied})
			c.Abort()
			return
		}
		c.Next()
	}
}
