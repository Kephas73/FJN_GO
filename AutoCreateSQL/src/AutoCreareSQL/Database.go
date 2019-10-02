package main

//type SqlDBStruct struct {
//	Driver      string
//	Server      string
//	User        string
//	Pwd         string
//	Database    string
//}
//
//func NewConnectDB() *SqlDBStruct {
//	paramater := SqlDBStruct{
//		Driver   :config.Str(DATABASE_DRIVER,""),
//		Server   :config.Str(DATABASE_SERVER,""),
//		User     :config.Str(DATABASE_USER,""),
//		Pwd      :config.Str(DATABASE_PWD,""),
//		Database :config.Str(DATABASE_NAME,""),
//	}
//
//	return &paramater
//}
//
//func (this *SqlDBStruct) ConnectDB() (*sql.DB, error) {
//	db, err := sql.Open(this.Driver,
//		fmt.Sprintf(this.User, this.Pwd, this.Server, this.Database))
//	if err != nil {
//		logger.Println(exterror.WrapExtError(err))
//		logger.Println("hhh")
//		return nil, err
//	}
//	return db,nil
//}