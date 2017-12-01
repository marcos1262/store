package DAO

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"store/util"
)

var db *sql.DB

func InitializeDB() {
	var err error
	db, err = sql.Open("mysql", "root:toor@/store")
	util.CheckMortalErr(err)
}

func CloseDB() {
	db.Close()
}
