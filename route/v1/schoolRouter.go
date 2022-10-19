package v1

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/controller"
)

func SchoolRouter(route *gin.RouterGroup) {
	route.Any("/", func(context *gin.Context) {
		context.String(0, route.BasePath())
	})

	controller.GradeController(route.Group("grade"))
	controller.TermController(route.Group("term"))

}
