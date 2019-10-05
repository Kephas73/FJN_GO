package main

import "Common"

func main() {

	OK := OKLogger{config.Str(PG_NAME,""),Common.GetDateTimeCurrent()}
	// A. Information program
	NAME_PROCESS := config.Str(PG_PROCESS,"")
	gLog.Println(MSG_LOG_PG_SPACE,NAME_PROCESS)
	gLog.Println(MSG_LOG_PG_START)
	defer gLog.Println(MSG_LOG_PG_END)

	if OK.CheckExits() {
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}
	//-------------------------------------------

	// B. Execute program
	err := CreateTableGCE()
	if err != nil {
		gLog.Println(err)
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}

	err = InsertTableGCE()
	if err != nil {
		gLog.Println(err)
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}

	err = OK.Set()
	if err != nil {
		gLog.Println(err)
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}
}
