package main

// Const database
const (
	CONFIG_FILE         = "config/AutoCreateSQL.cfg"

	DATABASE_DRIVER     = "DB.Driver"
	DATABASE_SERVER     = "DB.Server"
	DATABASE_USER       = "DB.User"
	DATABASE_PWD        = "DB.Pwd"
	DATABASE_NAME       = "DB.Database"
)

// Const PG
const (
	PG_NAME             = "PG.Name"
	PG_PROCESS          = "PG.Process"
	MSG_LOG_PG_SPACE    = "                  "
	MSG_LOG_PG_START    = "=============Process auto create table SQL start============="
	MSG_LOG_PG_END      = "==============Process auto create table SQL end=============="
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
