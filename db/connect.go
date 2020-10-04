package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
var Conn *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	// user := os.Getenv("MYSQL_USER")
	// // パスワード
	// password := os.Getenv("DB_USER")
	// // 接続先ホスト
	// host := os.Getenv("DB_HOST")
	// // 接続先ポート
	// port := os.Getenv("MYSQL_PORT")
	// // 接続先データベース
	// database := os.Getenv("MYSQL_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	// Conn, err = sql.Open(driverName,
	// 	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	connectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	dbName := os.Getenv("DB_NAME")
	// localConnection := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	cloudSQLConnection := user + ":" + pass + "@unix(/cloudsql/" + connectionName + ")/" + dbName + "?parseTime=true"
	Conn, err = sql.Open("mysql", cloudSQLConnection)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
