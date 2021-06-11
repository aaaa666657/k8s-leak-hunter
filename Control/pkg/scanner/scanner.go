package scanner

import (
	"context"
	"fmt"
	"time"

	"control/pkg/db"

	"github.com/Ullaakut/nmap/v2"
)

type errorSameService struct {
	Port        uint16
	Servicetype []string
}

type ScannerRes struct {
	ErrorDiffService      []errorSameService
	ErrorPortWithoutExist []db.Service
}

func scanner(hostid int) ([]db.Service, error) {
	ServicesType := make([]db.Service, 0, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	ip, err := db.LoadIP(hostid)
	if err != nil {
		fmt.Printf("%v \n", err)
		return ServicesType, err
	}
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts("1-65535"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		fmt.Println("unable to create nmap scanner: %v", err)
	}
	result, warnings, err := scanner.Run()
	if err != nil {
		fmt.Println("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		fmt.Println("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			ServicesType = append(ServicesType, db.Service{port.ID, port.Service.Name})
		}
	}

	//fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
	return ServicesType, nil
}

func ScannerService(hostid int, report_id int, triggertype string, timestemp string) (int, ScannerRes) {

	var errorDiffService []errorSameService

	rigisterService, _ := db.LoadService(hostid)

	scannerService, _ := scanner(hostid)

	res := difference(rigisterService, scannerService)

	var originals = len(res)
	fmt.Printf("%v\n", res)

	fmt.Printf("%d\n", originals)
	for i := 0; i < len(res); i++ {
		fmt.Printf("scaneer port: %d ", res[i].Port)
		fmt.Println(" service: ", res[i].Servicetype)
	}

	for i := 0; i < len(res); i++ {
		nowPort := res[i].Port
		var servicelist []string
		for y := i + 1; y < len(res); y++ {
			if nowPort == res[y].Port {
				fmt.Printf("%d\n", nowPort)
				servicelist = append(servicelist, res[i].Servicetype)
				servicelist = append(servicelist, res[y].Servicetype)
				errorDiffService = append(errorDiffService, errorSameService{nowPort, servicelist})
				res = append(res[:y], res[y+1:]...)
				res = append(res[:i], res[i+1:]...)
				i--
			}
		}
	}

	fmt.Printf("--------res--------------------------------\n")
	result := ScannerRes{errorDiffService, res}
	if originals != 0 {
		fmt.Printf("########################ERR########################\n")
		for i := 0; i < len(errorDiffService); i++ {
			db.InsertLog(report_id, "DiffService", triggertype, db.LoadHostname(hostid), int(errorDiffService[i].Port), errorDiffService[i].Servicetype[0], errorDiffService[i].Servicetype[1], timestemp)
			fmt.Printf("report_id : %d type : DiffService triggertype : %s  Hostname : %s port : %d Expected_Service : %s Scanned_service : %s time : %s", report_id, triggertype, db.LoadHostname(hostid), int(errorDiffService[i].Port), errorDiffService[i].Servicetype[0], errorDiffService[i].Servicetype[1], timestemp)
		}

		for i := 0; i < len(res); i++ {
			db.InsertLog(report_id, "PortWithoutExist", triggertype, db.LoadHostname(hostid), int(res[i].Port), "", res[i].Servicetype, timestemp)
			fmt.Printf("report_id : %d type : PortWithoutExist triggertype : %s Hostname : %s port : %d Scanned_service : %s time : %s", report_id, triggertype, db.LoadHostname(hostid), int(res[i].Port), res[i].Servicetype, timestemp)
		}
		return 1, result
	} else {
		fmt.Printf("########################PASS########################\n")
		db.InsertLog(report_id, "PASS", triggertype, db.LoadHostname(hostid), 0, "", "", timestemp)

		return 0, result
	}
}

func difference(slice1 []db.Service, slice2 []db.Service) []db.Service {
	var diff []db.Service

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
