package main

func main() {

	// A. Information program
	NAME_PROCESS := config.Str(PG_PROCESS,"")
	gLog.Println(MSG_LOG_PG_SPACE,NAME_PROCESS)
	gLog.Println(MSG_LOG_PG_START)
	defer gLog.Println(MSG_LOG_PG_END)
	//-------------------------------------------

	// B. Execute program


}
