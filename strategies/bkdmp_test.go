package strategies

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/alvaromuir/terrapin/utils"
	"github.com/joho/godotenv"
)

func TestBKtaxonomyBuyerRequest(t *testing.T) {
	_ = godotenv.Load("../.env")
	data := []byte("parentCategory.id=301170")
	resp, rslt, err := BKtaxonomyBuyerRequest(data)
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}
	body, err := json.Marshal(rslt)
	if err != nil {
		t.Errorf("An error has occured getting a response.")
	}
	testRslt := "count"
	respBody := strings.Split(string(body[:7]), "{\"")[1]
	if testRslt != respBody {
		fmt.Println(respBody)
		t.Errorf("An error has occured reading the response.")
	}
}

func TestBKsiteRequest(t *testing.T) {
	_ = godotenv.Load("../.env")
	resp, rslt, err := BKsiteRequest()
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}
	body, err := json.Marshal(rslt)
	if err != nil {
		t.Errorf("An error has occured getting the response.")
	}
	testRslt := "total_count"
	respBody := strings.Split(string(body[:13]), "{\"")[1]
	if testRslt != respBody {
		t.Errorf("An error has occured reading the response.")
	}
}

func TestBKcallEndpoint(t *testing.T) {
	_ = godotenv.Load("../.env")

	callType := "services"
	method := "GET"
	endPoint := "Ping"
	data := []byte("")
	URL := utils.BKsignRequest(callType, method, endPoint, data)

	resp, err := BKcallEndpoint(method, URL, data)
	if err != nil {
		t.Errorf("An error has occured calling an API.")
	}
	if resp.StatusCode != 200 {
		t.Errorf("The endPoint returned: %v", resp.Status)
	}

}
