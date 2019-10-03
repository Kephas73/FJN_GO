package main

// Const database
const (
	CONFIG_FILE         = "config/AutoCreateSQL.cfg"

	DATABASE_DRIVER     = "DB.Driver"
	DATABASE_HOST       = "DB.Host"
	DATABASE_PORT       = "DB.Port"
	DATABASE_USER       = "DB.User"
	DATABASE_PWD        = "DB.Pwd"
	DATABASE_NAME       = "DB.DBName"
)

// Const PG
const (
	PG_NAME             = "PG.Name"
	PG_PROCESS          = "PG.Process"
	MSG_LOG_PG_SPACE    = "                  "
	MSG_LOG_PG_START    = "=========================START==============================="
	MSG_LOG_PG_END      = "==========================END================================="
	MSG_LOG_PG_LINE     = "-------------------------------------------------------------"

	FORMAT_TIME_NOW     = "2006/01/02 15:04:05"
)

// Const LOG
const (
	CFG_OK_FOLDER_LOG     = "LOG.Name_OK"
	CFG_FOLDER_LOG        = "LOG.Name"
	FILE_EXTENSION_OKLOG  = ".OK"
	FILE_EXTENSION_LOG    = ".log"
)

// Const Create table
const (
	MSG_LOG_CREATE_TABLE_BEGIN  = "Begin the process of creating the GCE table..."
	MSG_LOG_CREATE_TABLE_DONE  =  "Complete the process of creating the GCE table"
)
