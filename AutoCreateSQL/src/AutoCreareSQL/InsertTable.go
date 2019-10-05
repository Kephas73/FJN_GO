package main

import (
	"fmt"
	"github.com/goframework/gf/exterror"
	"strings"
)

func InsertTableGCE() (eRR error) {
	gLog.Println(MSG_LOG_INSERT_TABLE_BEGIN)

	// Connect database mysql
	db := NewDatabase(config)
	err := db.ConnectDB()
	if err != nil {
		return exterror.WrapExtError(err)
	}
	defer db.CloseDB()

	// 1. Insert table
	listExc, err := CfgVarArray(config, CFG_INSERT_TABLE)
	if err != nil {
		return exterror.WrapExtError(err)
	}

	for _, v := range listExc {
		numRecord := fmt.Sprint(CountRecordInsert(v.Sql))
		gLog.Printf(MSG_LOG_INSERT_TABLE, v.Table, v.Path, numRecord)
		_, err := db.DB.Exec(v.Sql)
		if err != nil {
			return exterror.WrapExtError(err)
		}
	}
	gLog.Println(MSG_LOG_INSERT_TABLE_DONE)
	gLog.Printf(MSG_LOG_PG_LINE)
	return nil
}

func CountRecordInsert(sql string) int {
	v  := strings.Split(sql, "VALUES")
	v1 := strings.Split(v[1],"),")
	return len(v1)
}