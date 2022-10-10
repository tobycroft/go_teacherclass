package WechatModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "wechat"

func Api_find(project interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"project": project,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
