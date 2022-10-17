package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/school/model/SchoolAreaModel"
	"main.go/app/v1/school/model/SchoolModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Any("list", info_list)
	route.Any("get", info_get)
	route.Any("ids", info_ids)
	route.Any("get_domain", info_domain)

}

func info_ids(c *gin.Context) {
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

func info_get(c *gin.Context) {
	id, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := SchoolModel.Api_find(id)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_list(c *gin.Context) {
	schools := SchoolModel.Api_select()
	RET.Success(c, 0, schools, nil)
}

func info_domain(c *gin.Context) {
	domain, ok := Input.Post("domain", c, false)
	if !ok {
		return
	}
	data := SchoolModel.Api_find_byDomain(domain)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, "未找到学校")
	}
}
