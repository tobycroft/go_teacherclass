package Wechat

import (
	"errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/silenceper/wechat/v2/miniprogram/urlscheme"
	"github.com/tobycroft/Calc"
	"main.go/extend/Wechat/WechatKvModel"
	"main.go/extend/Wechat/WechatModel"
	"time"
)

func Wechat_login(project interface{}, js_code string) (OpenId string, UnionId string, err error) {
	wechat_data := WechatModel.Api_find(project)
	if len(wechat_data) < 1 {
		return "", "", errors.New("未找到项目")
	}
	appid := Calc.Any2String(wechat_data["appid"])
	appsecret := Calc.Any2String(wechat_data["appsecret"])
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
		return "", "", err
	}
	OpenId = ret.OpenID
	UnionId = ret.UnionID
	return
}

func Wechat_phone(project interface{}, code string) (*auth.GetPhoneNumberResponse, error) {
	wechat_data := WechatModel.Api_find(project)
	if len(wechat_data) < 1 {
		return nil, errors.New("未找到项目")
	}
	appid := Calc.Any2String(wechat_data["appid"])
	appsecret := Calc.Any2String(wechat_data["appsecret"])
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	return miniprogram.GetAuth().GetPhoneNumber(code)
}

func Wechat_qrcode(project interface{}, page, data string, width int) ([]byte, error) {
	wechat_data := WechatModel.Api_find(project)
	if len(wechat_data) < 1 {
		return nil, errors.New("未找到项目")
	}
	appid := Calc.Any2String(wechat_data["appid"])
	appsecret := Calc.Any2String(wechat_data["appsecret"])
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
	qr.Page = page
	//qr.Path = "pages/registerInfo/registerInfo"
	//qr.EnvVersion = "release"
	//qr.EnvVersion = "develop"
	qr.EnvVersion = Calc.Any2String(wechat_data["env_version"])
	md5 := Calc.Md5(data)
	scene := WechatKvModel.Api_find_val(md5)
	if scene == nil {
		if !WechatKvModel.Api_insert(md5, data) {
			return nil, errors.New("数据库插入失败")
		}
	}
	qr.Scene = Calc.Any2String(scene)
	qr.Width = width
	return miniprogram.GetQRCode().GetWXACodeUnlimit(qr)
}

func Wechat_create(project interface{}, page, data string, width int) ([]byte, error) {
	wechat_data := WechatModel.Api_find(project)
	if len(wechat_data) < 1 {
		return nil, errors.New("未找到项目")
	}
	appid := wechat_data["appid"].(string)
	appsecret := wechat_data["appsecret"].(string)
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
	qr.Page = page
	//qr.Page = "pages/registerInfo/registerInfo"
	//qr.Path = "pages/registerInfo/registerInfo"
	qr.EnvVersion = Calc.Any2String(wechat_data["env_version"])
	//qr.EnvVersion = "release"
	//qr.EnvVersion = "develop"

	md5 := Calc.Md5(data)
	qr.Scene = md5
	qr.Width = width
	if WechatKvModel.Api_find_val(md5) != nil {
		return miniprogram.GetQRCode().GetWXACodeUnlimit(qr)
	} else {
		if WechatKvModel.Api_insert(md5, data) {
			return miniprogram.GetQRCode().GetWXACodeUnlimit(qr)
		} else {
			return nil, errors.New("小程序码插入故障")
		}
	}
}

func Wechat_scheme(project interface{}, Path, Query string) (string, error) {
	wechat_data := WechatModel.Api_find(project)
	if len(wechat_data) < 1 {
		return "", errors.New("未找到项目")
	}
	appid := wechat_data["appid"].(string)
	appsecret := wechat_data["appsecret"].(string)
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     appid,
		AppSecret: appsecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	var qs urlscheme.USParams
	var jwxa urlscheme.JumpWxa
	jwxa.Path = Path
	jwxa.Query = Query
	switch wechat_data["env_version"] {
	case "develop":
		jwxa.EnvVersion = urlscheme.EnvVersionDevelop
		break
	case "release":
		jwxa.EnvVersion = urlscheme.EnvVersionRelease
		break
	case "trail":
		jwxa.EnvVersion = urlscheme.EnvVersionTrial
		break
	default:
		jwxa.EnvVersion = urlscheme.EnvVersionRelease
		break
	}
	qs.JumpWxa = &jwxa
	qs.ExpireType = 0
	qs.ExpireTime = time.Now().Add(7 * 24 * time.Hour).Unix()
	return miniprogram.GetSURLScheme().Generate(&qs)
}

func Wechat_scene(scene interface{}) interface{} {
	data := WechatKvModel.Api_find_val(scene)
	return data
}
