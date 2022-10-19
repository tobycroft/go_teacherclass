package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("upload", upload)
}

func upload(c *gin.Context) {

}
