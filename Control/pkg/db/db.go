package db

import (
	"database/sql"
	"errors"
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
	Port        uint16
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

func RegisterHost(hostname string, ip string) error {
	//--------check ip exist
	uidsql, _ := DB.Query("SELECT MAX(uid) FROM Host")
	var uid int
	for uidsql.Next() {
		_ = uidsql.Scan(&uid)
	}
	IPlist := make([]string, 0, 0)
	rows, err := DB.Query("SELECT ip FROM Host")
	if err != nil {
		fmt.Println("Query fail:", err)
		return err
	}
	for rows.Next() {
		var ip string
		err = rows.Scan(&ip)
		IPlist = append(IPlist, ip)
	}

	if findElementSTRING(IPlist, ip) {
		ipexisted := errors.New("IP exists in the database")
		return ipexisted
	}
	//----------insert
	uid++

	stmt, err := DB.Prepare("INSERT Host SET uid=?,hostname=?,ip=?")
	if err != nil {
		fmt.Println("Prepare fail:", err)
	}
	res, _ := stmt.Exec(uid, hostname, ip)
	if err != nil {
		fmt.Println("Exec fail:", err)
	}
	_ = res
	fmt.Printf("Insert Host uid : %d hostname : %s ip : %s Success ! \n", uid, hostname, ip)
	return nil

}

func RegisterService(hostid int, port int, servicetype string) error {
	//--------check port exist in the same id
	portlist := make([]int, 0, 0)
	rows, err := DB.Query("SELECT port FROM Service WHERE hostID IN(?);", hostid)
	if err != nil {
		fmt.Println("Query fail:", err)
		return err
	}
	for rows.Next() {
		var port int
		err = rows.Scan(&port)
		portlist = append(portlist, port)
	}

	if findElementINT(portlist, port) {
		portexisted := errors.New("port exists in the database")
		return portexisted
	}
	//-----insert
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

func LoadService(hostid int) ([]Service, error) {
	ServicesType := make([]Service, 0, 1)

	rows, err := DB.Query("SELECT port,service FROM Service WHERE hostID=?;", hostid)
	if err != nil {
		fmt.Println("Query fail:", err)
		return ServicesType, err
	}
	for rows.Next() {
		var port uint16
		var service string
		err = rows.Scan(&port, &service)
		ServicesType = append(ServicesType, Service{port, service})
	}
	return ServicesType, nil
}

func LoadIP(hostid int) (string, error) {
	ipsql, err := DB.Query("SELECT ip FROM Host WHERE uid=?", hostid)
	if err != nil {
		fmt.Println("find ip fail:", err)
		return "", err
	}
	var ip string
	for ipsql.Next() {
		_ = ipsql.Scan(&ip)
	}
	return ip, nil
}

func findElementINT(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}

func findElementSTRING(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
