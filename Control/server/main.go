package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"control/pkg/db"
	scannerPB "control/proto/scanner"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func (*Server) Register(ctx context.Context, req *scannerPB.ResourceRegister) (*scannerPB.ResourceRegisterResult, error) {
	fmt.Printf("Got Data %v \n", req)

	port := req.GetPort()
	service := req.GetServiceType()

	res := &scannerPB.ResourceRegisterResult{
		Result: true,
	}

	fmt.Printf("Got Dataus %d %s  \n", port, service)

	return res, nil
}

func main() {
	router := gin.Default()
	router.GET("/RegisterService/:uid/:port/:service", rigister_service)
	router.GET("/RegisterHost/:hostname/:ip", rigister_host)
	router.GET("/Loadhost", load_host)
	router.GET("/LoadService/:uid", load_service)
	router.GET("/LoadLogConut", load_log_conut)
	router.GET("/LoadLogDiffService/:uid", load_log_diffservice)
	router.GET("/LoadLogPortWithoutExist/:uid", load_log_portwithoutexist)

	router.Run(":8001")
}

func rigister_service(context *gin.Context) {

	hostnamestr := context.Param("uid")
	hostname, _ := strconv.Atoi(hostnamestr)
	portstr := context.Param("port")
	port, _ := strconv.Atoi(portstr)
	servicestr := context.Param("service")
	err := db.RegisterService(hostname, port, servicestr)
	errstring := fmt.Sprint(err)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": "",
		})
	}
}

func rigister_host(context *gin.Context) {

	hostname := context.Param("hostname")
	ip := context.Param("ip")
	uid, err := db.RegisterHost(hostname, ip)

	errstring := fmt.Sprint(err)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {

		mes := fmt.Sprintf("The host ID is %d", uid)
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": mes,
		})
	}

}

func load_host(context *gin.Context) {

	hostlist, err := db.LoadHost()
	if err != nil {
		fmt.Printf("err : %v \n", err)
	}
	for i := 0; i < len(hostlist); i++ {
		fmt.Printf("uid : %d hostname : %s ip : %s \n\n", hostlist[i].Uid, hostlist[i].Hostname, hostlist[i].Ip)
	}
	jsonData, _ := json.Marshal(hostlist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	errstring := fmt.Sprintf("%v", err)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": "",
			"json":    jsonData,
		})
	}
}

func load_service(context *gin.Context) {

	uidstr := context.Param("uid")
	uid, _ := strconv.Atoi(uidstr)
	servicelist, err := db.LoadService(uid)
	if err != nil {
		fmt.Printf("err : %v \n", err)
	}
	for i := 0; i < len(servicelist); i++ {
		fmt.Printf("uid : %d port : %d service : %s \n\n", uid, servicelist[i].Port, servicelist[i].Servicetype)
	}
	jsonData, _ := json.Marshal(servicelist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	errstring := fmt.Sprintf("%v", err)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": "",
			"json":    jsonData,
		})
	}
}
func load_log_conut(context *gin.Context) {
	log_count, err := db.LoadLogMax()

	errstring := fmt.Sprintf("%v", err)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": strconv.Itoa(log_count),
		})
	}
}
func load_log_diffservice(context *gin.Context) {
	uidstr := context.Param("uid")
	uid, _ := strconv.Atoi(uidstr)
	diff_service_list, err := db.LoadLogDiffService(uid)

	jsonData, _ := json.Marshal(diff_service_list)

	fmt.Println(string(jsonData))
	errstring := fmt.Sprintf("%v", err)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": "",
			"json":    jsonData,
		})
	}
}

func load_log_portwithoutexist(context *gin.Context) {
	uidstr := context.Param("uid")
	uid, _ := strconv.Atoi(uidstr)
	port_without_exist_list, err := db.LoadLogPortWithoutExist(uid)

	jsonData, _ := json.Marshal(port_without_exist_list)

	fmt.Println(string(jsonData))
	errstring := fmt.Sprintf("%v", err)

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": errstring,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Sussess",
			"message": "",
			"json":    jsonData,
		})
	}
}
