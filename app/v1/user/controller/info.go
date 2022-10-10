package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/extend/ASMS"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.LoginedController(), gin.Recovery())

	route.Any("edit", info_edit)
	route.Any("get", info_get)
	route.Any("my", info_my)
	route.Any("phone", info_phone)
	route.Any("change_phone", info_change_phone)

}

func info_get(c *gin.Context) {
	uid, ok := Input.PostInt64("id", c)
	if !ok {
		return
	}
	data := UserModel.Api_find_limit(uid)
	if len(data) > 0 {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_my(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := UserModel.Api_find(uid)
	if len(data) > 0 {
		delete(data, "password")
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}

func info_edit(c *gin.Context) {
	uid := c.GetHeader("uid")
	username, ok := Input.Post("username", c, true)
	if !ok {
		return
	}
	//wx_name, ok := Input.Post("wx_name", c, true)
	//if !ok {
	//	return
	//}
	wx_img, ok := Input.Post("wx_img", c, true)
	if !ok {
		return
	}
	if UserModel.Api_update_usernameAndNameAndWxImg(uid, username, username, wx_img) {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func info_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	data := UserModel.Api_find(uid)
	RET.Success(c, 0, data["phone"], nil)
}

func info_change_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	phone, ok := Input.Post("phone", c, true)
	if !ok {
		return
	}
	code, ok := Input.PostInt64("code", c)
	if !ok {
		return
	}
	if len(UserModel.Api_find_byPhone(phone)) > 0 {
		RET.Fail(c, 402, nil, "号码已被注册，请更换其他号码")
		return
	}
	data := ASMS.Api_find_in10(phone, code)
	if len(data) > 0 {
		if UserModel.Api_update_phone(uid, phone) {
			RET.Success(c, 0, nil, nil)
		} else {
			RET.Fail(c, 500, nil, nil)
		}
	} else {
		RET.Fail(c, 403, nil, "验证码超时")
	}
}
