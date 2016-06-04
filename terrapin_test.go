package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alvaromuir/terrapin/models"
	"github.com/alvaromuir/terrapin/utils"
)

// TestCallingApis ensures that the app can actually call a RESTFul API
func TestCallingApis(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := ""

	resp, err := utils.CoreRequest(method, URL, data, nil)
	if err != nil {
		resp.Body.Close()
		t.Errorf("An error has occured calling an API.")
	}
	resp.Body.Close()
}

// TestGettingResponse ensures the app gets some kind of response from an API
func TestGettingResponse(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := ""

	resp, _ := utils.CoreRequest(method, URL, data, nil)
	var rslt models.TestResponse
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		t.Errorf("An error has occured getting a response body")
	}
	resp.Body.Close()

}

// TestGettingOKStatus ensures the app gets a "200" status from API calls
func TestGettingOKStatus(t *testing.T) {
	method := "GET"
	URL := "http://jsonplaceholder.typicode.com/posts/1"
	data := ""

	resp, _ := utils.CoreRequest(method, URL, data, nil)
	resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("The status code was not 200.")
	}

}

func TestBKsignRequest(t *testing.T) {
	baseURL := "http://jsonplaceholder.typicode.com"
	method := "GET"
	endPoint := "posts/1"
	data := ""
	rslts := "?bkuid=e1fa6d0257144b859e5a30e0f28098f61686d886bad19d187e1f00506d" +
		"34ed81&bksig=2Oqk4OqbTTXM8K%2BHRQeNu0b8r8InOThg6zvA22w53VA%3D"
	if utils.BKsignRequest(baseURL, method, endPoint, data) != rslts {
		t.Errorf("The BKsignRequest method failed.")
	}

}

func TestSCGenWSSEHeader(t *testing.T) {
	username := "test: testing"
	rslts := strings.Split(utils.SCGenWSSEHeader(username), ",")
	if len(rslts) != 4 {
		t.Errorf("The SCGenWSSEHeader method returned too few segments")
	}

	if strings.Contains(rslts[0], username) == false {
		t.Errorf("The SCGenWSSEHeader method did not capture the correct username")
	}
}
