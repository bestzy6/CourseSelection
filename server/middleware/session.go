package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	//使用redis存储session
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))
	//用cookie作为session基础
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   7 * 86400,
		Path:     "/",
	})
	return sessions.Sessions("gin-session", store)
}
