package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/model/SchoolAreaModel"
	"main.go/app/v1/school/model/SchoolClassModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func ClassController(route *gin.RouterGroup) {

	route.Any("list", class_list)
	route.Any("ids", class_ids)
	route.Any("get", class_get)

}

func class_ids(c *gin.Context) {
	ids, ok := Input.PostArray[interface{}]("ids", c)
	if !ok {
		return
	}
	data := SchoolAreaModel.Api_select_in(ids)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func class_list(c *gin.Context) {
	datas := SchoolClassModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func class_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := SchoolClassModel.Api_find(id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
