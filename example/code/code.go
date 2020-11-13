package main

import (
	"log"
	"os"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/hsmtkk/paypaydev/code"
)

func main() {
	apiKey := getEnvVar("API_KEY")
	apiSecret := getEnvVar("API_KEY_SECRET")
	merchantID := getEnvVar("MERCHANT_ID")
	cred := auth.Credential{
		APIKey:       apiKey,
		APIKeySecret: apiSecret,
		MerchantID:   merchantID,
	}

	item := code.NewOrderItem("alpha", "bravo", 2, "charlie", 123)
	param := code.NewCreateCodeParameter([]code.OrderItem{item})
	creator := code.New()
	code, url, err := creator.CreateCode(cred, param)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(param.MerchantPaymentID)
	log.Print(code)
	log.Print(url)
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s is not defined", key)
	}
	return val
}
