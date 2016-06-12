package strategies

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// BKtaxonomyBuyerRequest returns DMP sites in JSON format
func BKtaxonomyBuyerRequest(data []byte) (*http.Response, *models.BKbuyerViewCategoryResult, error) {
	callType := "taxonomy"
	method := "GET"
	endPoint := "categories"
	URL := utils.BKsignRequest(callType, method, endPoint, data)

	resp, err := BKcallEndpoint(method, URL, data)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	var rslt models.BKbuyerViewCategoryResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR: %s", err)
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
		log.Fatalf("ERROR: %s", err)
		return resp, nil, err
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
	return resp, nil
}
