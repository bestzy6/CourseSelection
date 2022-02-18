package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	//使用redis存储session
	//store, _ := redis.NewStore(10, "tcp", "192.168.152.6:6379", "", []byte(secret))
	//用cookie存储session
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge:   1800,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	})
	return sessions.Sessions("camp-session", store)
}
