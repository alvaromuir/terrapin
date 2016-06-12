package utils

import (
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestBKsignRequest(t *testing.T) {
	_ = godotenv.Load("../.env")

	baseURL := "http://jsonplaceholder.typicode.com"
	method := "GET"
	endPoint := "posts/1"
	data := []byte("")
	testRslt := "?bkuid"
	rslt := strings.Split(BKsignRequest(baseURL, method, endPoint, data), "=")[0]
	if rslt != testRslt {
		t.Errorf("The BKsignRequest method failed.")
	}
}

func TestSCGenWSSEHeader(t *testing.T) {
	username := "test: testing"
	rslts := strings.Split(SCGenWSSEHeader(username), ",")
	if len(rslts) != 4 {
		t.Errorf("The SCGenWSSEHeader method returned too few segments")
	}

	if strings.Contains(rslts[0], username) == false {
		t.Errorf("The SCGenWSSEHeader method did not capture the correct username")
	}
}
