package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var userName, _ = os.LookupEnv("DBUSER")
var password, _ = os.LookupEnv("DBPW")
var host, _ = os.LookupEnv("DBIP")
var port, _ = os.LookupEnv("DBPORT")
var dbName = "scanner"

var dbpath = strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")
var DB, _ = sql.Open("mysql", dbpath)

//DB

type Service struct {
	Port        int
	Servicetype string
}

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

func RegisterHost(uid int, hostname string, ip string) error {
	//stmt, err := DB.Prepare("INSERT INTO Service ('hostIP','port','service') VALUES (?, ?, ?)")

	stmt, err := DB.Prepare("INSERT Host SET uid=?,hostname=?,ip=?")
	if err != nil {
		fmt.Println("Prepare fail:", err)
	}
	res, _ := stmt.Exec(uid, hostname, ip)
	if err != nil {
		fmt.Println("Exec fail:", err)
	}
	_ = res
	fmt.Printf("Insert Service uid : %d hostname : %s ip : %d Success ! \n", uid, hostname, ip)
	return nil

}

func RegisterService(hostid int, port int, servicetype string) error {
	/* //stmt, err := DB.Prepare("INSERT INTO Service ('hostIP','port','service') VALUES (?, ?, ?)")

	err := DB.QueryRow("SELECT port FROM Service WHERE hostID IN (?)", hostid).Scan(&port)

	if err != nil {
		if err != sql.ErrNoRows {
		}
		stmt, err := DB.Prepare("INSERT Service SET hostID=?,port=?,service=?")
		if err != nil {
			fmt.Println("Prepare fail:", err)
		}
		res, _ := stmt.Exec(hostid, port, servicetype)
		if err != nil {
			fmt.Println("Exec fail:", err)
		}
		_ = res
		fmt.Printf("RegisterService HostID: %d Port: %d Service: %s Success! \n", hostid, port, servicetype)
		return nil
	}
	portexisted := errors.New("port exists in the database")
	fmt.Print(portexisted)
	return portexisted */

	stmt, err := DB.Prepare("INSERT Service SET hostID=?,port=?,service=?")
	if err != nil {
		fmt.Println("Prepare fail:", err)
	}
	res, _ := stmt.Exec(hostid, port, servicetype)
	if err != nil {
		fmt.Println("Exec fail:", err)
	}
	_ = res
	fmt.Printf("Insert Service HostID : %d Port : %d Service : %s Success ! \n", hostid, port, servicetype)
	return nil
}

func LoadService(hostip string) ([]Service, error) {
	ServicesType := make([]Service, 0, 1)

	rows, err := DB.Query("SELECT * FROM Service WHERE hostIP IN(?);", hostip)
	if err != nil {
		fmt.Println("Query fail:", err)
		return ServicesType, err
	}
	for rows.Next() {
		var host string
		var port int
		var service string
		err = rows.Scan(&host, &port, &service)
		ServicesType = append(ServicesType, Service{port, service})
	}
	return ServicesType, nil
}
