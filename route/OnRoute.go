package route

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}
		user := version1.Group("user")
		{
			v1.UserRouter(user)
		}
		school := version1.Group("school")
		{
			v1.SchoolRouter(school)
		}
		attach := version1.Group("attach")
		{
			v1.AttachRouter(attach)
		}
	}
}
