package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//const (
//	username string = "root"
//	password string = ""
//	database string = "gomart"
//)

//var (
//	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
//)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:admin@tcp(192.168.0.4:3306)/gomart")

	if err != nil {
		return nil, err
	}

	return db, nil
}