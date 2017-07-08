package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DbConf struct {
	Driver string
	DSN    string
}

var Db *sqlx.DB

func Connect(conf *DbConf) *sqlx.DB {
	db, err := sqlx.Connect(conf.Driver, conf.DSN)
	if err != nil {
		panic("connect to database failed " + err.Error())
	}

	if Db == nil {
		Db = db
	}

	return db
}
