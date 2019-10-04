package main

import (
	"github.com/goframework/gf/exterror"
)

func CreateTableGCE() (eRR error) {
	gLog.Println(MSG_LOG_CREATE_TABLE_BEGIN)
	defer gLog.Println(MSG_LOG_CREATE_TABLE_DONE)

	// Connect database mysql
	db := NewDatabase(config)
	err := db.ConnectDB()
	if err != nil {
		return exterror.WrapExtError(err)
	}
	defer db.CloseDB()

	// 1. Create table
	listExc, err := CfgVarArray(config,CFG_CREATE_TABLE)
	if err != nil {
		return err
	}

	for _, v := range listExc {
		gLog.Printf(MSG_LOG_CREATE_TABLE, v.Table, v.Path)
	}


	return nil

}