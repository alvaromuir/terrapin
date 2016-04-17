package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// CoreRequest is a bare-bones JSON request, essential for all other API calls
func CoreRequest(method string, URL string, data string, addHeaders map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, strings.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()

		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		os.Exit(1)
		return nil, nil
	}
	return resp, nil
}
