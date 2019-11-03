package main

// Const database
const (
	CONFIG_FILE = "config/AutoCreateSQL.cfg"

	DATABASE_DRIVER = "DB.Driver"
	DATABASE_HOST   = "DB.Host"
	DATABASE_PORT   = "DB.Port"
	DATABASE_USER   = "DB.User"
	DATABASE_PWD    = "DB.Pwd"
	DATABASE_NAME   = "DB.DBName"
)

// Const PG
const (
	PG_NAME          = "PG.Name"
	PG_PROCESS       = "PG.Process"
	MSG_LOG_PG_SPACE = "                  "
	MSG_LOG_PG_START = "=========================START==============================="
	MSG_LOG_PG_END   = "==========================END================================="
	MSG_LOG_PG_LINE  = "---------------------------------------------------------------"

	FORMAT_TIME_NOW = "2006/01/02 15:04:05"
	SYMBOL_ARROW    = "->"
	CHAR_SPACE      = " "
	MSG_LOG_RETRY   = "Retrying PG process -> Retry: [%v]"
)

// Const LOG
const (
	CFG_OK_FOLDER_LOG    = "LOG.Name_OK"
	CFG_FOLDER_LOG       = "LOG.Name"
	FILE_EXTENSION_OKLOG = ".OK"
	FILE_EXTENSION_LOG   = ".log"
)

// Const Setting Mail
const (
	CFG_MAIL_ENABLE         = "Mail.Enable"
	CFG_MAIL_SERVER         = "Mail.Server"
	CFG_MAIL_PORT           = "Mail.Port"
	CFG_MAIL_USER           = "Mail.User"
	CFG_MAIL_PASS           = "Mail.Pass"
	CFG_MAIL_FROM_ADDRESS   = "Mail.FromAddress"
	CFG_MAIL_FROM_NAME      = "Mail.FromName"
	CFG_MAIL_TO_ADDRESS     = "Mail.ToAddress"
	CFG_MAIL_TO_NAME        = "Mail.ToName"
)


// Const Create table
const (
	CFG_CREATE_TABLE           = "TABLE"
	MSG_LOG_CREATE_TABLE_BEGIN = "Begin the process of create the GCE table..."
	MSG_LOG_CREATE_TABLE_DONE  = "Complete the process of create the GCE table"
	MSG_LOG_CREATE_TABLE       = "Create table [%v] -> [%v]"
)

// Const Insert table
const (
	CFG_INSERT_TABLE           = "DATA"
	MSG_LOG_INSERT_TABLE_BEGIN = "Begin the process of insert the GCE table..."
	MSG_LOG_INSERT_TABLE_DONE  = "Complete the process of insert the GCE table"
	MSG_LOG_INSERT_TABLE       = "Insert table [%v] -> [%v] -> RECORD[%v]"
)

// Const File Exits
const MSG_LOG_FILE_OK_EXITS       = "The process for today has ended on this device. To continue the process, please delete today's OKLog file."

