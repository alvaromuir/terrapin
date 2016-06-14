package strategies

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// BKtaxonomyBuyerRequest returns DMP sites in JSON format
func BKtaxonomyBuyerRequest(parentID int) (*http.Response, *models.BKbuyerViewCategoryResult, error) {
	data := []byte("parentCategory.id=" + strconv.Itoa(parentID))
	fmt.Println(string(data))
	callType := "taxonomy"
	method := "GET"
	endPoint := "categories"
	URL := utils.BKsignRequest(callType, method, endPoint, data)

	resp, err := BKcallEndpoint(method, URL, data)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("ERROR fetching data. Query returned '%s'.", resp.Status)
		return nil, nil, nil
	}
	var rslt models.BKbuyerViewCategoryResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// BKsiteRequest returns DMP sites in JSON format
func BKsiteRequest() (*http.Response, *models.BKsiteResult, error) {
	callType := "services"
	method := "GET"
	endPoint := "sites"
	data := []byte("")
	URL := utils.BKsignRequest(callType, method, endPoint, data)

	resp, err := BKcallEndpoint(method, URL, data)
	if err != nil {
		log.Fatalf("ERROR decoding response: %s", err)
		return resp, nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatalf("ERROR fetching data. Query returned '%s'.", resp.Status)
		return nil, nil, nil
	}
	var rslt models.BKsiteResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// BKcallEndpoint returns a raw BK API response
func BKcallEndpoint(method string, URL string, data []byte) (*http.Response, error) {
	resp, err := utils.CoreRequest(method, URL, data, nil)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, err
	}
	if resp.StatusCode != 200 {
		log.Fatalf("ERROR fetching data. Query returned '%s'.", resp.Status)
		return nil, nil
	}
	return resp, nil
}
