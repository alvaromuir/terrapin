package strategies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// BKpingRequest generates the method request
func BKpingRequest(URL string) (*http.Response, *models.BKpingResponse, error) {
	method := "GET"
	data := ""
	resp, err := utils.CoreRequest(method, URL, data, nil)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt models.BKpingResponse
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	resp.Body.Close()
	return resp, &rslt, nil
}

// BKtaxonomyBuyerRequest returns DMP sites in JSON format
func BKtaxonomyBuyerRequest(method string, URL string, data string) (*http.Response, *models.BKbuyerViewCategoryResult, error) {
	resp, err := utils.CoreRequest(method, URL, data, nil)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt models.BKbuyerViewCategoryResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// BKsiteRequest returns DMP sites in JSON format
func BKsiteRequest(method string, URL string, data string) (*http.Response, *models.BKsiteResult, error) {
	resp, err := utils.CoreRequest(method, URL, data, nil)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt models.BKsiteResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	return resp, &rslt, nil
}
