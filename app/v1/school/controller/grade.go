package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/model/SchoolGradeModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"time"
)

func GradeController(route *gin.RouterGroup) {

	route.Any("list", grade_list)
	route.Any("get", grade_get)

}

func grade_list(c *gin.Context) {
	datas := SchoolGradeModel.Api_select()
	for i, data := range datas {
		data["ttt"] = time.Now().Format("2006-01-02 15:04:05")
		data["zzz"] = time.Now().Format(time.RFC3339)
		datas[i] = data
	}
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
