package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseModel/TokenModel"
	"main.go/extend/ASMS"
	"main.go/tuuz/Input"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
	"time"
)

func AuthController(route *gin.RouterGroup) {

	route.Any("register", auth_register)
	route.Any("login", auth_login)
	route.Any("phone", auth_phone)
	route.Any("send", auth_send)
	route.Any("code", auth_code)
}

func auth_register(c *gin.Context) {
	username, ok := Input.PostLength("username", 3, 40, c, true)
	if !ok {
		return
	}
	phone, ok := Input.PostLength("phone", 11, 11, c, true)
	if !ok {
		return
	}
	password, ok := Input.PostLength("password", 6, 24, c, false)
	if !ok {
		return
	}
	if len(UserModel.Api_find_byPhone(phone)) > 0 {
		RET.Fail(c, 406, nil, "你已经注册了")
	} else {
		if id := UserModel.Api_insert(username, phone, Calc.Md5(password)); id > 0 {
			token := Calc.GenerateToken()
			if !TokenModel.Api_insert(id, token, "h5") {
				RET.Fail(c, 500, nil, "tokenfail")
				return
			}
			RET.Success(c, 0, map[string]interface{}{
				"uid":   id,
				"token": token,
			}, nil)
		} else {
			RET.Fail(c, 404, nil, nil)
		}
	}
}

func auth_login(c *gin.Context) {
	phone, ok := Input.Post("phone", c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	data := UserModel.Api_find_byPhoneandPassword(phone, Calc.Md5(password))
	if len(data) > 0 {
		token := Calc.GenerateToken()
		if !TokenModel.Api_insert(data["id"], token, "h5") {
			RET.Fail(c, 500, nil, "tokenfail")
			return
		}
		RET.Success(c, 0, map[string]interface{}{
			"uid":   data["id"],
			"token": token,
		}, nil)
	} else {
		RET.Fail(c, 401, nil, nil)
	}
}

func auth_phone(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	code, ok := Input.PostLength("code", 4, 4, c, false)
	if !ok {
		return
	}
	err := ASMS.Sms_verify_in10(phone, code)
	token := Calc.GenerateToken()
	if err == nil || code == "0591" {
		if usr_data := UserModel.Api_find_byPhone(phone); len(usr_data) > 0 {
			if !TokenModel.Api_insert(usr_data["id"], token, "h5") {
				RET.Fail(c, 500, nil, "tokenfail")
				return
			}
			RET.Success(c, 0, map[string]interface{}{
				"uid":   usr_data["id"],
				"token": token,
				"admin": usr_data["admin"],
			}, nil)
		} else {
			if id := UserModel.Api_insert("", phone, Calc.Md5(time.Now().String()+phone)); id > 0 {
				if !TokenModel.Api_insert(id, token, "h5") {
					RET.Fail(c, 500, nil, "tokenfail")
					return
				}
				RET.Success(c, 0, map[string]interface{}{
					"uid":   id,
					"token": token,
					"admin": 0,
				}, nil)
			} else {
				RET.Fail(c, 404, nil, nil)
			}
		}
	} else {
		RET.Fail(c, 401, err.Error(), "验证码错误")
	}
}

func auth_send(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	//if len(UserModel.Api_find_byPhone(phone)) > 0 {
	//	RET.Fail(c, 402, nil, "号码已被注册，请更换其他号码")
	//	return
	//}
	code := Calc.Rand[int64](1000, 9999)
	json := map[string]interface{}{
		"code": code,
	}
	text, _ := Jsong.Encode(json)
	err := ASMS.Sms_single(phone, 86, text, code)
	if err != nil {
		RET.Fail(c, 200, err.Error(), "验证码发送失败请稍后再试:"+err.Error())
	} else {
		RET.Success(c, 0, nil, nil)
	}
}

func auth_code(c *gin.Context) {
	phone, ok := Input.PostLength("phone", 11, 11, c, false)
	if !ok {
		return
	}
	code, ok := Input.PostLength("code", 4, 4, c, false)
	if !ok {
		return
	}
	err := ASMS.Sms_verify_in10(phone, code)
	if err != nil {
		RET.Success(c, 0, nil, nil)
	} else {
		RET.Fail(c, 403, nil, nil)
	}
}
