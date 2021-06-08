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
	router.Use(Cors())

	router.GET("/RegisterService/:uid/:port/:service", register_service)
	router.GET("/RegisterHost/:hostname/:ip", register_host)
	router.GET("/Loadhost", load_host)
	router.GET("/Loadservice", load_service)
	router.GET("/LoadLogConut", load_log_conut)
	router.GET("/LoadLogDiffService/:uid", load_log_diffservice)
	router.GET("/LoadLogPortWithoutExist/:uid", load_log_portwithoutexist)

	router.Run(":8001")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //請求頭部
		if origin != "" {
			//接收客戶端傳送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//伺服器支援的所有跨域請求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允許跨域設定可以返回其他子段，可以自定義欄位
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session, content-type")
			// 允許瀏覽器（客戶端）可以解析的頭部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//設定快取時間
			c.Header("Access-Control-Max-Age", "172800")
			//允許客戶端傳遞校驗資訊比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允許型別校驗
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func register_service(context *gin.Context) {

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
			"message": "Sussess",
		})
	}
}

func register_host(context *gin.Context) {

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

/* func load_service(context *gin.Context) {

	servicelist, err := db.LoadServiceAll()
	if err != nil {
		fmt.Printf("err : %v \n", err)
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
} */

func load_service(context *gin.Context) {
	hostlist, err := db.LoadServiceAll()
	if err != nil {
		fmt.Printf("err : %v \n", err)
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
