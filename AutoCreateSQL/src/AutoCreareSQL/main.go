package main

import (
	"Common"
	"time"
)

var maxCountRetry = 3

func main() {

	OK := OKLogger{config.Str(PG_NAME, ""), Common.GetDateTimeCurrent()}
	// A. Information program
	NAME_PROCESS := config.Str(PG_PROCESS, "")
	gLog.Println(MSG_LOG_PG_SPACE, NAME_PROCESS)
	gLog.Println(MSG_LOG_PG_START)
	defer gLog.Println(MSG_LOG_PG_END)

	if OK.CheckExits() {
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}
	//-------------------------------------------
	// PG retry
	countRetry := 0
	for {
		err := func() error {
			err1 := CreateTableGCE()
			if err1 != nil {
				gLog.Println(err1)
				gLog.Println(MSG_LOG_PG_LINE)
				return err1
			}
			err1 = InsertTableGCE()
			if err1 != nil {
				gLog.Println(err1)
				gLog.Println(MSG_LOG_PG_LINE)
				return err1
			}
			return nil
		}()
		if err == nil {
			break
		} else {
			if countRetry < maxCountRetry {
				countRetry ++
				gLog.Printf(MSG_LOG_RETRY, countRetry)
				gLog.Println(MSG_LOG_PG_LINE)
				time.Sleep(time.Second * 60)
			} else {
				return
			}
		}
	}

	err := OK.Set()
	if err != nil {
		gLog.Println(err)
		gLog.Println(MSG_LOG_PG_LINE)
		return
	}
}
