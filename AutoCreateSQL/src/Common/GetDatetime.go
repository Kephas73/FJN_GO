package Common

import "time"

const FORMAT_DATE_TIME_CURRENT_  =  "20060102"

func GetDateTimeCurrent() string {
	return time.Now().Format(FORMAT_DATE_TIME_CURRENT_)
}
