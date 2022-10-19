package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Input"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("upload", upload)
}

func upload(c *gin.Context) {
	Type, ok := Input.PostIn("type", c, []string{"视频", "风采", "照片墙"})
	if !ok {
		return
	}
	category, ok := Input.Post("category", c, true)
	if !ok {
		return
	}
	title, ok := Input.SPostString("title", c, true)
	if !ok {
		return
	}
	content, ok := Input.SPostString("content", c, true)
	if !ok {
		return
	}
	url, ok := Input.SPostString("url", c, true)
	if !ok {
		return
	}
	
}
