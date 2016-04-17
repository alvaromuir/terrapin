package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var env = godotenv.Load(".env")

// Add your API Keys below
var bkuid = os.Getenv("BK_KEY")              // BlueKai API KEY
var bksecretkey = os.Getenv("BK_SECRET")     // BlueKai API Secret
var bkpartnerid = os.Getenv("BK_PARTNER_ID") // BlueKai partner ID

var scusernanme = os.Getenv("ADOBE_USERNAME") // SiteCatalst API KEY
var scsecret = os.Getenv("ADOBE_SECRET")      // SiteCatalst API Secret

var bkBaseURLS = map[string]string{
	"services": "http://services.bluekai.com/Services/WS/",
	"taxonomy": "https://taxonomy.bluekai.com/taxonomy/",
}

var omnitureBaseURLS = map[string]string{
	"admin": "https://api.omniture.com/admin/1.4/rest/",
}

// BKsignRequest returns a HMAC-SHA256 string for BlueKai JSON request
func BKsignRequest(baseURL string, method string, endPoint string, data string) string {

	var URL string
	if baseURL == "services" {
		URL = bkBaseURLS["services"] + endPoint
	}

	if baseURL == "taxonomy" {
		URL = bkBaseURLS["taxonomy"] + endPoint + "?" + data
		if strings.HasPrefix(data, "parentCategory") {
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

	if data != "" {
		stringToSign += data
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
func SCGenWSSEHeader() string {
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
		scusernanme,
		base64.StdEncoding.EncodeToString(digest.Sum(nil)),
		base64.StdEncoding.EncodeToString(nonce),
		created)
}
