package SchoolTermModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "tc_school_term"

func Api_select(grade_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("grade_id", grade_id)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(id interface{}) gorose.Data {
	db := tuuz.Db().Table(Table)
	where := map[string]interface{}{
		"id": id,
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

func Api_select_in(ids []interface{}) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.WhereIn("id", ids)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
