package strategies

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// SCGgetRealTimeSettings returns setup of realtime reports
func SCGgetRealTimeSettings() (*http.Response, error) {
	// REVIEW: move this to models
	reqData := []byte(`{
			"rsid_list":[
				"verizontelecomres"
				]
			}`)

	method := "POST"
	URL := "https://api.omniture.com/admin/1.4/rest/?method=ReportSuite.GetRealTimeSettings"

	resp, err := SCcallEndpoint(method, URL, reqData)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, err
	}
	return resp, nil
}

// SCGgetRealTimeResults returns realtime results in JSON format
func SCGgetRealTimeResults(metrics []string) (*http.Response, *models.SCRealtimeMetricsResult, error) {
	var metricsArray []models.SCMetricsArray
	for _, el := range metrics {
		metricsArray = append(metricsArray, models.SCMetricsArray{ID: el})
	}
	jsonData := models.SCMetricsRequest{}

	reqString, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return nil, nil, err
	}
	reqData := []byte(string(reqString))

	method := "POST"
	URL := "https://api.omniture.com/admin/1.4/rest/?method=Report.Run"

	resp, err := SCcallEndpoint(method, URL, reqData)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, nil, err
	}
	var rslt models.SCRealtimeMetricsResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// SCcallEndpoint returns a raw SC API response
func SCcallEndpoint(method string, URL string, data []byte) (*http.Response, error) {
	headers := map[string]string{
		"Authorization": "WSSE profile=\"UsernameToken\"",
		"X-WSSE":        utils.SCGenWSSEHeader(""),
	}

	resp, err := utils.CoreRequest(method, URL, data, headers)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, err
	}
	return resp, nil
}
