package controller

import (
	"github.com/Unknwon/goconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseModel/SystemParamModel"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/extend/Wechat"
	"main.go/extend/Wechat/WechatKvModel"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func WechatController(route *gin.RouterGroup) {

	route.Use(cors.Default())

	route.Any("login", wechat_login)
	route.Any("phone", wechat_phone)
	route.Any("qrcode", wechat_qrcode)
	route.Any("scene", wechat_scene)
	route.Any("create", wechat_create)
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
	ret, err := Wechat.Wechat_scheme(app_conf.Project, path, query)
	if err != nil {
		RET.Fail(c, 200, ret, err.Error())
	} else {
		RET.Success(c, 0, ret, nil)
	}
}

func wechat_login(c *gin.Context) {
	js_code, ok := Input.Post("js_code", c, false)
	if !ok {
		return
	}
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	}
	value, err := conf.GetSection("wechat")
	if err != nil {
	}
	appid := value["appid"]
	appsecret := value["appsecret"]
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	//TODO 调用对应接口
	ret, err := miniprogram.GetAuth().Code2Session(js_code)
	if err != nil {
		RET.Fail(c, 401, nil, nil)
		return
	}
	md5_pass := Calc.Md5(app_conf.Project + ret.SessionKey)
	token := Calc.GenerateToken()
	if user := UserModel.Api_find_byWxId(ret.OpenID); len(user) > 0 {
		TokenModel.Api_insert(user["id"], token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   user["id"],
		}, nil)
		return
	}
	if id := UserModel.Api_insert_more("wx_"+ret.OpenID, "wx_"+ret.OpenID, md5_pass, ret.OpenID, ret.UnionID, ""); id > 0 {
		token = Calc.GenerateToken()
		TokenModel.Api_insert(id, token, "wx")
		RET.Success(c, 0, map[string]interface{}{
			"token": token,
			"uid":   id,
		}, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}

	//wpcfg := &config2.Config{
	//	AppID:     appid,
	//	MchID:     appsecret,
	//	Key:       appsecret,
	//	NotifyURL: "http://127.0.0.1",

	//}
	//pay := wc.GetPay(wpcfg)
	//pay.GetOrder().PrePayOrder()
}

func wechat_phone(c *gin.Context) {
	uid := c.GetHeader("uid")
	code, ok := Input.Post("code", c, false)
	if !ok {
		return
	}
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	}
	value, err := conf.GetSection("wechat")
	if err != nil {
	}
	appid := value["appid"]
	appsecret := value["appsecret"]
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	ret, err := miniprogram.GetAuth().GetPhoneNumber(code)
	if err != nil {
		RET.Fail(c, 402, nil, nil)
		return
	}
	if UserModel.Api_update_phone(uid, ret.PhoneInfo.PurePhoneNumber) {
		RET.Success(c, 0, ret.PhoneInfo.PurePhoneNumber, nil)
	} else {
		RET.Fail(c, 500, nil, nil)
	}
}

func wechat_qrcode(c *gin.Context) {
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	}
	value, err := conf.GetSection("wechat")
	if err != nil {
	}
	appid := value["appid"]
	appsecret := value["appsecret"]
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	//var qs urlscheme.USParams
	//var jwxa urlscheme.JumpWxa
	//jwxa.Path = "pages/registerInfo/registerInfo"
	//jwxa.Query = "type=register&school_id=1&grade=1&class=19"
	//jwxa.EnvVersion = "develop"
	//qs.JumpWxa = &jwxa
	//qs.ExpireType = 0
	//qs.ExpireTime = time.Now().Add(7 * 24 * time.Hour).Unix()
	//ret, err := miniprogram.GetSURLScheme().Generate(&qs)
	var qr qrcode.QRCoder
	qr.Page = "pages/registerInfo/registerInfo"
	//qr.Path = "pages/registerInfo/registerInfo"
	//qr.EnvVersion = "release"
	//qr.EnvVersion = "develop"
	qr.Scene = "reg_code"
	qr.Width = 600
	bt, err := miniprogram.GetQRCode().GetWXACodeUnlimit(qr)
	if err != nil {
		RET.Fail(c, 200, bt, err.Error())
	} else {
		c.Writer.Write(bt)
	}
}

func wechat_create(c *gin.Context) {
	data, ok := Input.Combi("data", c, false)
	if !ok {
		return
	}
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	}
	value, err := conf.GetSection("wechat")
	if err != nil {
	}

	appid := value["appid"]
	appsecret := value["appsecret"]
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	//var qs urlscheme.USParams
	//var jwxa urlscheme.JumpWxa
	//jwxa.Path = "pages/registerInfo/registerInfo"
	//jwxa.Query = "type=register&school_id=1&grade=1&class=19"
	//jwxa.EnvVersion = "develop"
	//qs.JumpWxa = &jwxa
	//qs.ExpireType = 0
	//qs.ExpireTime = time.Now().Add(7 * 24 * time.Hour).Unix()
	//ret, err := miniprogram.GetSURLScheme().Generate(&qs)
	var qr qrcode.QRCoder
	qr.Page = "pages/registerInfo/registerInfo"
	//qr.Path = "pages/registerInfo/registerInfo"
	version := SystemParamModel.Api_find_val("wx_qr_envversion")
	qr.EnvVersion = Calc.Any2String(version)
	//qr.EnvVersion = "release"
	//qr.EnvVersion = "develop"

	md5 := Calc.Md5(data)
	qr.Scene = md5
	qr.Width = 400
	if val := WechatKvModel.Api_find(md5); val["bin"] != nil {
		c.Writer.Write([]byte(val["bin"].(string)))
	} else {
		bt, err := miniprogram.GetQRCode().GetWXACodeUnlimit(qr)
		if err != nil {
			RET.Fail(c, 200, bt, err.Error())
			return
		}
		if val != nil {
			if WechatKvModel.Api_update(md5, data, bt) {
			} else {
				RET.Fail(c, 500, nil, "程序码存储故障")
				return
			}
		} else {
			if WechatKvModel.Api_insert(md5, data) {

			} else {
				RET.Fail(c, 500, nil, "小程序故障")
				return
			}
		}

		c.Writer.Write(bt)
	}
}

func wechat_scene(c *gin.Context) {
	scene, ok := Input.Post("scene", c, false)
	if !ok {
		return
	}
	data := WechatKvModel.Api_find_val(scene)
	if data != nil {
		RET.Success(c, 0, data, nil)
	} else {
		RET.Fail(c, 404, nil, nil)
	}
}
