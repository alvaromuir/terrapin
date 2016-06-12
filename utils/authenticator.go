package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var env = godotenv.Load(".env")

// BKsignRequest returns a HMAC-SHA256 string for BlueKai JSON request
func BKsignRequest(callType string, method string, endPoint string, data []byte) string {
	// Add your API Keys below
	var bkuid = os.Getenv("BK_KEY")              // BlueKai API KEY
	var bksecretkey = os.Getenv("BK_SECRET")     // BlueKai API Secret
	var bkpartnerid = os.Getenv("BK_PARTNER_ID") // BlueKai partner ID

	var bkBaseURLS = map[string]string{
		"services": "http://services.bluekai.com/Services/WS/",
		"taxonomy": "https://taxonomy.bluekai.com/taxonomy/",
		"audience": "api.bluekai.com/audience/v1/",
	}

	var reqData string
	if len(data) < 1 {
		reqData = ""
	} else {
		jsonData := new(bytes.Buffer)
		// if the data is not a JSON object
		if bytes.Contains(data, []byte("=")) {
			reqData = bytes.NewBuffer(data).String()
		} else {
			if err := json.Compact(jsonData, data); err != nil {
				log.Fatalf("JSON ENCODING ERROR: %s", err)
			}
			reqData = jsonData.String()
		}
	}

	var URL string

	if callType == "services" {
		URL = bkBaseURLS["services"] + endPoint
	}

	if callType == "taxonomy" {
		URL = bkBaseURLS["taxonomy"] + endPoint + "?" + reqData
		if strings.HasPrefix(reqData, "parentCategory") {
			URL += "&view=BUYER"
		} else {
			URL += "&view=OWNER"
		}
		URL += "&showReach=true&countryCode=US&partner.id=" + bkpartnerid
	}

	stringToSign := method

	parsedURL, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	stringToSign += parsedURL.Path

	qP := strings.Split(parsedURL.RawQuery, "&")
	if len(qP) > 0 {
		for qS := 0; qS < len(qP); qS++ {
			qP2 := strings.Split(qP[qS], "=")
			if len(qP2) > 1 {
				stringToSign += qP2[1]
			}
		}
	}

	if reqData != "" {
		stringToSign += reqData
	}

	h := hmac.New(sha256.New, []byte(bksecretkey))
	h.Write([]byte(stringToSign))
	digest := base64.StdEncoding.EncodeToString(h.Sum(nil))
	u := url.QueryEscape(digest)

	newURL := URL

	if strings.Contains(URL, "?") {
		newURL += "&"
	} else {
		newURL += "?"
	}

	newURL += "bkuid=" + bkuid + "&bksig=" + u
	return newURL
}

// SCGenWSSEHeader returns WSSE header for authenticated requests
func SCGenWSSEHeader(username string) string {

	var scusernanme = os.Getenv("ADOBE_USERNAME") // SiteCatalst API KEY
	var scsecret = os.Getenv("ADOBE_SECRET")      // SiteCatalst API Secret

	if len(username) < 1 {
		username = scusernanme
	}
	timeStamp := time.Now()
	locale, _ := time.LoadLocation("GMT")
	created := timeStamp.In(locale).Format("2006-01-02T15:04:05Z")

	nonce := make([]byte, 20)
	rand.Read(nonce)

	digest := sha1.New()
	digest.Write([]byte(nonce))
	digest.Write([]byte(created))
	digest.Write([]byte(scsecret))

	return fmt.Sprintf(`UsernameToken Username="%s", PasswordDigest="%s", Nonce="%s", Created="%s"`,
		username,
		base64.StdEncoding.EncodeToString(digest.Sum(nil)),
		base64.StdEncoding.EncodeToString(nonce),
		created)
}
