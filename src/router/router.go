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
		//授权
		v1.POST("auth", api.Auth)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)

			//新闻操作
			authed.GET("news", api.ListNews)
			authed.GET("news/:id", api.ShowNews)

			//媒体操作
			authed.GET("media", api.ListMedia)
			authed.GET("media/:id", api.ShowMedia)

			//美句操作
			authed.GET("sentence", api.ListSentence)
			authed.GET("sentence/:id", api.ShowSentence)

			authed.POST("dailyReport", api.ShowDailyReport)
		}
	}
	return r
}