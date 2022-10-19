package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/attach/model/AttachModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func IndexController(route *gin.RouterGroup) {

	route.Any("upload", index_upload)
	route.Any("list", index_list)
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
	term_id, ok := Input.PostInt64("term_id", c)
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
	if id := AttachModel.Api_insert(Type, category, term_id, title, content, url); id > 0 {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func index_list(c *gin.Context) {
	Type, ok := Input.PostIn("type", c, []string{"视频", "风采", "照片墙"})
	if !ok {
		return
	}
	term_id, ok := Input.PostInt64("term_id", c)
	if !ok {
		return
	}
	category, _ := Input.SPostString("category", c, false)
	limit, page, err := Input.PostLimitPage(c)
	if err != nil {
		return
	}
	datas := AttachModel.Api_select(Type, category, term_id, limit, page)
	RET.Success(c, 0, datas, nil)
}
