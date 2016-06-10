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

// SCGgetRealTimeResults returns realtime results in raw format
func SCGgetRealTimeResults(metrics []string) (*http.Response, error) {
	var metricsArray []models.SCMetricsArray
	for _, el := range metrics {
		metricsArray = append(metricsArray, models.SCMetricsArray{ID: el})
	}
	jsonData := &models.SCMetricsRequest{
		ReportDescription: &models.SCMetricsReportDescription{
			Source:        "realtime",
			ReportSuiteID: "verizontelecomres",
			Metrics:       metricsArray,
		},
	}

	reqString, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return nil, nil
	}
	reqData := []byte(string(reqString))

	method := "POST"
	URL := "https://api.omniture.com/admin/1.4/rest/?method=Report.Run"

	resp, err := SCcallEndpoint(method, URL, reqData)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, err
	}
	return resp, nil
}

// SCcallEndpoint returns a raw API response
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
