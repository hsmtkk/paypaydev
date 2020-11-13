package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/hsmtkk/paypaydev/auth"
)

func main() {
	apiKey := getEnvVar("API_KEY")
	apiSecret := getEnvVar("API_SECRET")
	merchantID := getEnvVar("MERCHANT_ID")

	bs, err := ioutil.ReadFile("./request.json")
	if err != nil {
		log.Fatal(err)
	}
	body := string(bs)
	param := auth.HMACAuthHeaderParameter{
		APIKey:             apiKey,
		APIKeySecret:       apiSecret,
		RequestURL:         "/v2/codes",
		RequestMethod:      "POST",
		RequestBody:        body,
		RequestContentType: "application/json;charset=UTF-8;",
	}
	token, err := auth.New().HMACAuthHeader(param)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(token)
	req, err := http.NewRequest(http.MethodPost, "https://stg-api.sandbox.paypay.ne.jp/v2/codes", bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", param.RequestContentType)
	req.Header.Add("X-ASSUME-MERCHANT", merchantID)
	reqBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(reqBytes))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Print(resp.StatusCode)
	log.Print(resp.Status)
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(out))
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("environment variable %s is not defined", key)
	}
	return val
}
