package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	err := CreateDatabase(this)
	if err != nil {
		return exterror.WrapExtError(err)
	}

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

func CreateDatabase(this *DBConnection) error {

	var dbName []string
	flagDB := 0
	// Connect Mysql do not connect database
	sqlconn := this.User + ":" + this.Pwd + "@tcp(" + this.Host + ":" + this.Port + ")/"
	db, err := sql.Open(this.Driver, sqlconn)
	if err != nil {
		return exterror.WrapExtError(err)
	}

	// Show list database in Mysql
	row, err := db.Query(QUERY_SHOW_DATABASE)
	if err != nil {
		return exterror.WrapExtError(err)
	}

	for row.Next() {
		var item string
		err := row.Scan(&item)
		if err != nil {
			return exterror.WrapExtError(err)
			continue
		}

		dbName = append(dbName, item)
	}

	// Check exits database
	for _,v := range dbName {
		if v == this.DBName {
			flagDB ++
		}
	}

	if flagDB == 0 {
		_, err :=db.Exec(QUERY_CREATE_DATABASE + CHAR_SPACE + this.DBName)
		if err != nil {
			return exterror.WrapExtError(err)
		}
	}

	return nil
}

const QUERY_SHOW_DATABASE    = "SHOW DATABASES"
const QUERY_CREATE_DATABASE  = "CREATE DATABASE"