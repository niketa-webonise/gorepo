package database

import (_ "github.com/go-sql-driver/mysql"
        "database/sql")

func GetConnection() *sql.DB {

	dbDriver:= "mysql"
	dbUser := "niketa"
	dbPass := "jain"
	dbName := "dbusers"

	db,err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err!=nil {
		panic(err.Error())
	}
	return db
}