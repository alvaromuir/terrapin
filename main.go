package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alvaromuir/terrapin/strategies"
	"github.com/alvaromuir/terrapin/utils"
)

func pingSC() {
	// test SC
	method := "GET"
	URL := "https://api.omniture.com/admin/1.4/rest/?method=Company.GetVersionAccess"
	reqData := []byte("")
	resp, err := strategies.SCcallEndpoint(method, URL, reqData)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("SC Status: %v \n", resp.Status)
}

func pingBK() {
	// test bk
	callType := "services"
	method := "GET"
	endPoint := "Ping"
	data := ""
	resp, _, _ := strategies.BKpingRequest(
		utils.BKsignRequest(callType, method, endPoint, data))
	fmt.Printf("BK Status: %v \n", resp.Status)
}

func getSCSettings() {
	resp, err := strategies.SCGgetRealTimeSettings()

	if err != nil {
		resp.Body.Close()
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Println("")
	fmt.Printf("response status: %v\n", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getSCRealTimeMetrics(metrics []string) {
	resp, err := strategies.SCGgetRealTimeResults(metrics)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Println("")
	fmt.Printf("response status: %v\n", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {

	go pingBK()
	go pingSC()
	go getSCSettings()
	getSCRealTimeMetrics([]string{"pageviews"})

}
