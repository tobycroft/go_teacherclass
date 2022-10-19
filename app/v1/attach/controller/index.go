package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/attach/model/AttachModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("upload", upload)
}

func index_upload(c *gin.Context) {
	Type, ok := Input.PostIn("type", c, []string{"视频", "风采", "照片墙"})
	if !ok {
		return
	}
	category, ok := Input.Post("category", c, true)
	if !ok {
		return
	}
	term_id := Input.SPostDefault("term_id", c, int64(0))
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
	if id := AttachModel.Api_insert(Type, category, title, content, url); id > 0 {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func index_list(c *gin.Context) {

}
