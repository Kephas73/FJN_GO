package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/goframework/gf/cfg"
	"github.com/goframework/gf/exterror"
)

type DBConnection struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pwd    string
	DBName string

	DB *sql.DB
}

func NewDatabase(config *cfg.Cfg) DBConnection {
	db := DBConnection{}
	db.Driver = config.Str(DATABASE_DRIVER, "")
	db.Host = config.Str(DATABASE_HOST, "")
	db.Port = config.Str(DATABASE_PORT, "")
	db.User = config.Str(DATABASE_USER, "")
	db.Pwd = config.Str(DATABASE_PWD, "")
	db.DBName = config.Str(DATABASE_NAME, "")

	return db
}

func (this *DBConnection) ConnectDB() error {
	sqlconn := this.User + ":" + this.Pwd + "@tcp(" + this.Host + ":" + this.Port + ")/" + this.DBName
	db, err := sql.Open(this.Driver, sqlconn)
	if err != nil {
		return exterror.WrapExtError(err)
	}
	this.DB = db
	return nil
}

func (this *DBConnection) CloseDB() {
	if this.DB != nil {
		this.DB.Close()
	}
}
