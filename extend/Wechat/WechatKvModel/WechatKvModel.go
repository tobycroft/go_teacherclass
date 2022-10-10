package WechatKvModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "wechat_kv"

func Api_insert(key, val interface{}) bool {
	db := tuuz.Db().Table(Table)
	data := map[string]interface{}{
		"key": key,
		"val": val,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert2(key, val, bin interface{}) bool {
	db := tuuz.Db().Table(Table)
	data := map[string]interface{}{
		"key": key,
		"val": val,
		"bin": bin,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_val(key interface{}) interface{} {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("val")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(key interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"key": key,
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

func Api_update(key, val, bin interface{}) bool {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	data := map[string]interface{}{
		"key": key,
		"val": val,
		"bin": bin,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
