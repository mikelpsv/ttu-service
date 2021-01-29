package app

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func InitDb() error {

	var err error
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", Cfg.DbHost, Cfg.DbName, Cfg.DbUser, Cfg.DbPassword)
	fmt.Println(connStr)
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = Db.Ping()
	if err != nil {
		return err
		//log.Fatal(err)
	}
	return nil
}
