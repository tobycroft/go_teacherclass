package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/model/SchoolAreaModel"
	"main.go/app/v1/school/model/SchoolGradeModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func GradeController(route *gin.RouterGroup) {

	route.Any("list", grade_list)
	route.Any("ids", grade_ids)
	route.Any("get", grade_get)

}

func grade_ids(c *gin.Context) {
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

func grade_list(c *gin.Context) {
	datas := SchoolGradeModel.Api_select()
	RET.Success(c, 0, datas, nil)
}

func grade_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := SchoolGradeModel.Api_find(id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
