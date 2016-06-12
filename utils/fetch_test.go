package utils

import (
	"encoding/json"
	"testing"

	"github.com/alvaromuir/terrapin/models"
)

// TestCoreRequest ensures that the app can actually call a RESTFul API
func TestCoreRequest(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := []byte("")

	resp, err := CoreRequest(method, URL, data, nil)
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("Error: Call response was '%v'", resp.Status)
	}
}

// TestGettingResponse ensures the app gets some kind of response from an API
func TestGettingResponse(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := []byte("")

	resp, _ := CoreRequest(method, URL, data, nil)
	var rslt models.TestResponse
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		t.Errorf("An error has occured getting a response body")
	}
}

// TestGettingOKStatus ensures the app gets a "200" status from API calls
func TestGettingOKStatus(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := []byte("")

	resp, _ := CoreRequest(method, URL, data, nil)

	if resp.StatusCode != 200 {
		t.Errorf("The status code was not 200.")
	}
}
