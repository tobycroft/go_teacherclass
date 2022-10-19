package model

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "tc_attach"

func Api_insert(Type, category, title, content, url interface{}) int64 {
	db := tuuz.Db().Table(Table)
	data := map[string]any{
		"type":     Type,
		"category": category,
		"title":    title,
		"content":  content,
		"url":      url,
	}
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select(Type, category interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("type", Type)
	db.Where("category", category)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
