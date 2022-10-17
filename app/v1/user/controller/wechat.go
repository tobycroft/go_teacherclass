package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/extend/Wechat/WechatKvModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func WechatController(route *gin.RouterGroup) {

	route.Use(cors.Default())

	route.Any("login", wechat_login)
	route.Any("phone", wechat_phone)
	route.Any("scene", wechat_scene)
	route.Any("create", wechat_create)
	route.Any("create_file", wechat_create_file)
	route.Any("scheme", wechat_scheme)
}

func wechat_scheme(c *gin.Context) {
	path, ok := Input.Post("path", c, false)
	if !ok {
		return
	}
	query, ok := Input.Post("query", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_wxa_generatescheme(app_conf.Project, path, query, true, 180)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
	} else {
		RET.Success(c, 0, ret.Openlink, nil)
	}
}

func wechat_login(c *gin.Context) {
	js_code, ok := Input.Combi("js_code", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_sns_jscode2session(app_conf.Project, js_code)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
		return
	}
	md5_pass := Calc.Md5(app_conf.Project + ret.SessionKey)
	token := Calc.GenerateToken()
	if user := UserModel.Api_find_byWxId(ret.Openid); len(user) > 0 {
		TokenModel.Api_insert(user["id"], token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   user["id"],
		}, nil)
		return
	}
	if id := UserModel.Api_insert_more("wx_"+ret.Openid, "wx_"+ret.Openid, md5_pass, ret.Openid, ret.Unionid, ""); id > 0 {
		token = Calc.GenerateToken()
		TokenModel.Api_insert(id, token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   id,
		}, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	code, ok := Input.Post("code", c, false)
	if !ok {
		return
	}
	ret, err := AossGoSdk.Wechat_wxa_getuserphonenumber(app_conf.Project, code)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
		return
	}
	if err != nil {
		RET.Fail(c, 402, nil, nil)
		return
	}
	if UserModel.Api_update_phone(uid, ret.PurePhoneNumber) {
		RET.Success(c, 0, ret.PurePhoneNumber, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_create_file(c *gin.Context) {
	data, ok := Input.Combi("data", c, false)
	if !ok {
		return
	}
	url, err := AossGoSdk.Wechat_wxa_unlimited_file(app_conf.Project, data, "pages/registerInfo/registerInfo")
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		RET.Success(c, 0, url, nil)
	}
}

func wechat_create(c *gin.Context) {
	data, ok := Input.Combi("data", c, false)
	if !ok {
		return
	}
	file_url, err := AossGoSdk.Wechat_wxa_unlimited_file(app_conf.Project, data, "pages/registerInfo/registerInfo")
	if err != nil {
		RET.Fail(c, 200, nil, err.Error())
	} else {
		c.Redirect(302, file_url)
	}
}

func wechat_scene(c *gin.Context) {
	scene, ok := Input.Combi("scene", c, false)
	if !ok {
		return
	}
	data := WechatKvModel.Api_find_val(scene)
	if data != nil {
		RET.Success(c, 0, data, nil)
	} else {
		sc, err := AossGoSdk.Wechat_wxa_scene(app_conf.Project, scene)
		if err != nil {
			RET.Fail(c, 404, nil, err.Error())
			return
		}
		RET.Success(c, 0, sc.Val, nil)
	}
}
