package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/model/SchoolTermModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func TermController(route *gin.RouterGroup) {

	route.Any("list", class_list)
	route.Any("get", class_get)

}

func class_list(c *gin.Context) {
	datas := SchoolTermModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func class_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := SchoolTermModel.Api_find(id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
