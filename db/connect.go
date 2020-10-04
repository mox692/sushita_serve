package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"google.golang.org/appengine"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

var Conn *sql.DB

func init() {
	var err error

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	connectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	switch appengine.IsAppEngine() {
	case true:
		cloudSQLConnection := user + ":" + pass + "@unix(/cloudsql/" + connectionName + ")/" + dbName + "?parseTime=true"
		Conn, err = sql.Open("mysql", cloudSQLConnection)
		if err != nil {
			log.Println(err)
			panic(err.Error())
		}
	case false:
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName))
		if err != nil {
			log.Println(err)
			panic(err.Error())
		}

	}
}
