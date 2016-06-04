package strategies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// SCpingRequest generates the method request
func SCpingRequest(URL string, wsseAuth string) (*http.Response, *models.SCpingResponse, error) {
	method := "GET"
	data := ""
	headers := map[string]string{
		"Authorization": "WSSE profile=\"UsernameToken\"",
		"X-WSSE":        wsseAuth,
	}

	resp, err := utils.CoreRequest(method, URL, data, headers)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt models.SCpingResponse
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	resp.Body.Close()
	return resp, &rslt, nil
}

// SCGetReportBookmarks returns a list of user bookmarked reports
func SCGetReportBookmarks(URL string, wsseAuth string) (*http.Response, *models.SCBookmarks, error) {
	method := "GET"
	data := ""
	headers := map[string]string{
		"Authorization": "WSSE profile=\"UsernameToken\"",
		"X-WSSE":        wsseAuth,
	}

	resp, err := utils.CoreRequest(method, URL, data, headers)
	if err != nil {
		defer resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}

	var rslt models.SCBookmarks
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	return resp, &rslt, nil
}
