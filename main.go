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

func getBKCategories(parentID int) {
	resp, rslt, err := strategies.BKtaxonomyBuyerRequest(parentID)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	if resp.StatusCode == 200 {
		// body, err := json.Marshal(rslt)
		// if err != nil {
		// 	log.Fatalf("An error has occured getting the response.")
		// }
		// fmt.Print(string(body))
		fmt.Printf("\n%s \t%13s \t%12s\n",
			"BKID", "REACH", "NAME")
		for _, item := range rslt.Items {
			fmt.Printf("#%d \t%-10d \t%s\n",
				item.ID, item.Stats.Reach, item.Name)
		}
		// fmt.Println(rslt.Count)
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

func getBKSites() {
	resp, rslt, err := strategies.BKsiteRequest()
	if err != nil {
		log.Fatalf("An error has occured calling an API.")
	}
	if resp.StatusCode == 200 {
		for _, item := range rslt.Sites {
			fmt.Printf("#%d \t%s \t%s\n",
				item.ID, item.UpdatedAt, item.Name)
		}
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

func main() {
	// pingBK()
	// pingSC()
	getBKCategories(301170)
	// getBKSites()
}
