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
type Host struct {
	Uid      int    `json:"Uid"`
	Hostname string `json:"Hostname"`
	Ip       string `json:"Ip"`
}

type DiffService struct {
	TriggerType     string
	Hostname        string
	Port            int
	ExpectedService string
	ScannedService  string
	ScannedAt       string
}

type PortWithoutExist struct {
	TriggerType    string
	Hostname       string
	Port           int
	ScannedService string
	ScannedAt      string
}

func InitDB() error {
	DB.SetConnMaxLifetime(100)
	// 設定 database 最大連接數
	DB.SetMaxIdleConns(1000)
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

func RegisterHost(hostname string, ip string) (int, error) {
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
		return -1, err
	}
	for rows.Next() {
		var ip string
		err = rows.Scan(&ip)
		IPlist = append(IPlist, ip)
	}

	if findElementSTRING(IPlist, ip) {
		ipexisted := errors.New("IP exists in the database")
		return -1, ipexisted
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
	return uid, nil

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

func LoadHostname(hostid int) string {
	hostnamesql, err := DB.Query("SELECT hostname FROM Host WHERE uid=?", hostid)
	if err != nil {
		fmt.Println("find ip fail:", err)
		return ""
	}
	var hostname string
	for hostnamesql.Next() {
		_ = hostnamesql.Scan(&hostname)
	}
	return hostname
}

func LoadHost() ([]Host, error) {
	Hostlist := make([]Host, 0, 2)
	hostsql, err := DB.Query("SELECT * FROM Host ")
	if err != nil {
		fmt.Println("find ip fail:", err)
		return Hostlist, err
	}
	for hostsql.Next() {
		var uid int
		var hostname string
		var ip string
		err = hostsql.Scan(&uid, &hostname, &ip)
		Hostlist = append(Hostlist, Host{uid, hostname, ip})
	}
	return Hostlist, nil
}

func InsertLogID() int {
	uidsql, err := DB.Query("SELECT MAX(id) FROM ScannerLog")
	var uid int
	if err != nil {
		fmt.Println("Prepare fail:", err)
		return 0
	} else {
		for uidsql.Next() {
			_ = uidsql.Scan(&uid)
		}
		uid++
		fmt.Printf("uid : %d\n", uid)
		return uid
	}
}

func InsertLog(report_id int, reportType string, triggerType string, hostname string, port int, expected_service string, scanned_service string, scanned_at string) error {

	//----------insert
	stmt, err := DB.Prepare("INSERT ScannerLog SET id=?,report_type=?,trigger_type=?,hostname=?,port=?,expected_service=?,scanned_service=?,scanned_at=?")
	if err != nil {
		fmt.Println("Prepare fail:", err)
	}
	res, _ := stmt.Exec(report_id, reportType, triggerType, hostname, port, expected_service, scanned_service, scanned_at)
	if err != nil {
		fmt.Println("Exec fail:", err)
	}
	_ = res
	return nil
}

func LoadLogMax() (int, error) {
	idsql, err := DB.Query("SELECT MAX(id) FROM ScannerLog")
	var id int
	if err != nil {
		fmt.Println("Prepare fail:", err)
		return -1, err
	} else {
		for idsql.Next() {
			_ = idsql.Scan(&id)
		}
		fmt.Printf("uid : %d\n", id)
		return id, err
	}
}

func LoadLogDiffService(uid int) ([]DiffService, error) {
	diffServicelist := make([]DiffService, 0, 5)
	diffServicesql, err := DB.Query("SELECT trigger_type,hostname,port,expected_service,scanned_service,scanned_at FROM ScannerLog WHERE id=? AND report_type=? ", uid, "DiffService")
	if err != nil {
		fmt.Println("Find diffService fail:", err)
		return diffServicelist, err
	}
	for diffServicesql.Next() {
		var trigger_type string
		var hostname string
		var port int
		var expected_service string
		var scanned_service string
		var scanned_at string
		err = diffServicesql.Scan(&trigger_type, &hostname, &port, &expected_service, &scanned_service, &scanned_at)
		diffServicelist = append(diffServicelist, DiffService{trigger_type, hostname, port, expected_service, scanned_service, scanned_at})
	}
	return diffServicelist, nil
}

func LoadLogPortWithoutExist(uid int) ([]PortWithoutExist, error) {
	port_without_existlist := make([]PortWithoutExist, 0, 4)
	port_without_existsql, err := DB.Query("SELECT trigger_type,hostname,port,scanned_service,scanned_at FROM ScannerLog WHERE id=? AND report_type=? ", uid, "PortWithoutExist")
	if err != nil {
		fmt.Println("Find PortWithoutExist fail:", err)
		return port_without_existlist, err
	}
	for port_without_existsql.Next() {
		var trigger_type string
		var hostname string
		var port int
		var scanned_service string
		var scanned_at string
		err = port_without_existsql.Scan(&trigger_type, &hostname, &port, &scanned_service, &scanned_at)
		port_without_existlist = append(port_without_existlist, PortWithoutExist{trigger_type, hostname, port, scanned_service, scanned_at})
	}
	return port_without_existlist, nil
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
