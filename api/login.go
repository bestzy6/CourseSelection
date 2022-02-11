package api

//// Login 登入api
//func Login(c *gin.Context) {
//	var req model.LoginRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusOK, model.LoginResponse{
//			Code: model.ParamInvalid,
//			Data: struct {
//				UserID string
//			}{UserID: ""},
//		})
//	} else {
//		c.JSON(http.StatusOK, service.LoginService(&req))
//	}
//}
//
//// Logout 登出api
//func Logout(c *gin.Context) {
//	var req model.LogoutRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusOK, model.LogoutResponse{
//			Code: model.ParamInvalid,
//		})
//	} else {
//		c.JSON(http.StatusOK, service.LogoutService(&req))
//	}
//}
//
//// Whoami 是get方法
//func Whoami(c *gin.Context) {
//	c.JSON(http.StatusOK, nil)
//}
