package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/strategies"
	"github.com/alvaromuir/terrapin/utils"
)

func main() {
	// test SC
	fmt.Println("SC Testing . . .")
	testURL := "https://api.omniture.com/admin/1.4/rest/?method=Bookmark.GetBookmarks"
	scAuth := utils.SCGenWSSEHeader()
	_, scRslt, err := strategies.SCGetReportBookmarks(testURL, scAuth)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(scRslt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	fmt.Print("\n")
	// test bk
	fmt.Println("BK Testing . . .")
	callType := "services"
	method := "GET"
	endPoint := "Ping"
	data := ""
	resp, _, _ := strategies.BKpingRequest(
		utils.BKsignRequest(callType, method, endPoint, data))
	pingResponse := models.BKpingResponse{Status: resp.StatusCode}
	bkRslt, err := json.Marshal(pingResponse)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error:%v\n", resp.StatusCode)
	}
	fmt.Println(string(bkRslt))
}
