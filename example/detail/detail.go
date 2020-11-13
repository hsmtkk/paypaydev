package main

import (
	"log"
	"os"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/hsmtkk/paypaydev/detail"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s merchantPaymentID", os.Args[0])
	}

	merchantPaymentID := os.Args[1]

	apiKey := getEnvVar("API_KEY")
	apiSecret := getEnvVar("API_KEY_SECRET")
	merchantID := getEnvVar("MERCHANT_ID")
	cred := auth.Credential{
		APIKey:       apiKey,
		APIKeySecret: apiSecret,
		MerchantID:   merchantID,
	}

	getter := detail.New()
	result, err := getter.GetDetail(cred, merchantPaymentID)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s is not defined", key)
	}
	return val
}
