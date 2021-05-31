package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "scanner"
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "scanner"
)

var dbpath = strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")

var DB, _ = sql.Open("mysql", dbpath)

func InitDB() error {
	DB.SetConnMaxLifetime(100)
	// 設定 database 最大連接數
	DB.SetMaxIdleConns(10)
	//設定上 database 最大閒置連接時間
	// 驗證是否連上 db
	err := DB.Ping()
	if err != nil {
		fmt.Println("open database fail:", err)
		return err
	} else {
		fmt.Println("connnect success , ")
	}
	return nil
}
func RegisterService(hostip string, port int, service string) error {
	//stmt, err := DB.Prepare("INSERT INTO Service ('hostIP','port','service') VALUES (?, ?, ?)")
	stmt, err := DB.Prepare("INSERT Service SET hostIP=?,port=?,service=?")
	if err != nil {
		fmt.Println("Prepare fail:", err)
		return err
	}
	res, _ := stmt.Exec(hostip, port, service)
	if err != nil {
		fmt.Println("Exec fail:", err)
		return err
	}
	fmt.Println("Sql.Result:", res)
	return nil
}
