package main

import (
	"fmt"
	"log"

	"github.com/alvaromuir/terrapin/strategies"
	"github.com/alvaromuir/terrapin/utils"
)

func pingSC() {
	// test SC
	method := "GET"
	URL := "https://api.omniture.com/admin/1.4/rest/?method=Company.GetVersionAccess"
	data := []byte("")
	resp, err := strategies.SCcallEndpoint(method, URL, data)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("SC Status: %v \n", resp.Status)
}

func pingBK() {
	// test bk
	callType := "services"
	method := "GET"
	endPoint := "Ping"
	data := []byte("")
	URL := utils.BKsignRequest(callType, method, endPoint, data)

	resp, err := strategies.BKcallEndpoint(method, URL, data)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("BK Status: %v \n", resp.Status)
}

func main() {
	go pingBK()
	pingSC()
}
