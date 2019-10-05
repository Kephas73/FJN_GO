package main

import (
	"github.com/goframework/gf/exterror"
)

func CreateTableGCE() (eRR error) {
	gLog.Println(MSG_LOG_CREATE_TABLE_BEGIN)

	// Connect database mysql
	db := NewDatabase(config)
	err := db.ConnectDB()
	if err != nil {
		return exterror.WrapExtError(err)
	}
	defer db.CloseDB()

	// 1. Create table
	listExc, err := CfgVarArray(config, CFG_CREATE_TABLE)
	if err != nil {
		return exterror.WrapExtError(err)
	}

	for _, v := range listExc {
		gLog.Printf(MSG_LOG_CREATE_TABLE, v.Table, v.Path)
		_, err := db.DB.Exec(v.Sql)
		if err != nil {
			return exterror.WrapExtError(err)
		}
	}
	gLog.Println(MSG_LOG_CREATE_TABLE_DONE)
	gLog.Printf(MSG_LOG_PG_LINE)
	return nil
}
