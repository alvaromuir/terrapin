package utils

import (
	"bytes"
	"log"
	"net/http"
)

// CoreRequest is a bare-bones JSON request, essential for all other API calls
func CoreRequest(method string, URL string, data []byte, addHeaders map[string]string) (*http.Response, error) {
	// println(string(data[:len(data)]))
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return resp, err
	}
	return resp, nil
}
