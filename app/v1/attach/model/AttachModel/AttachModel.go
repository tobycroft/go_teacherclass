package AttachModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "tc_attach"

func Api_insert(Type, category, term_id, title, content, url interface{}) int64 {
	db := tuuz.Db().Table(Table)
	data := map[string]any{
		"type":     Type,
		"category": category,
		"term_id":  term_id,
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

func Api_select(Type, category, term_id interface{}, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table)
	db.Fields("*", "Unix_Timestamp(date) as date_int")
	if Type != nil {
		db.Where("type", Type)
	}
	if category != "" {
		db.Where("category", category)
	}
	if term_id != nil {
		db.Where("term_id", term_id)
	}
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Paginator()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}
