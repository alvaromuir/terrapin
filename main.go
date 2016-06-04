package main

import (
	"fmt"

	"github.com/alvaromuir/terrapin/strategies"
	"github.com/alvaromuir/terrapin/utils"
)

func pingSC() {
	// test SC
	testURL := "https://api.omniture.com/admin/1.4/rest/?method=Bookmark.GetBookmarks"
	scAuth := utils.SCGenWSSEHeader("")
	resp, _, _ := strategies.SCpingRequest(testURL, scAuth)
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

func main() {

	go pingBK()
	pingSC()
}
