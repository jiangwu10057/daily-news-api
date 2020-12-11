package router

import (
	"news/api"
	"news/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	store := cookie.NewStore([]byte("golang-gin:3.1.0"))
	r.Use(sessions.Sessions("news-session", store))
	// 路由
	v1 := r.Group("/api/v1")
	{
		//新闻操作
		v1.GET("news", api.ListNews)
		v1.GET("news/:id", api.ShowNews)

		//媒体操作
		v1.GET("media", api.ListMedia)
		v1.GET("media/:id", api.ShowMedia)

		//美句操作
		v1.GET("sentence", api.ListSentence)
		v1.GET("sentence/:id", api.ShowSentence)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)
		}
	}
	return r
}