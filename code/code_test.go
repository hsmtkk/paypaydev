package code_test

import (
	"log"
	"os"
	"testing"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/hsmtkk/paypaydev/code"
	"github.com/stretchr/testify/assert"
)

func TestCreateCode(t *testing.T) {
	cred := auth.Credential{
		APIKey:       os.Getenv("API_KEY"),
		APIKeySecret: os.Getenv("API_SECRET"),
		MerchantID:   os.Getenv("MERCHANT_ID"),
	}
	log.Print(cred)
	item := code.NewOrderItem("alpha", "bravo", 1, "charlie", 100)
	param := code.NewCreateCodeParameter([]code.OrderItem{item})
	creator := code.New()
	code, url, err := creator.CreateCode(cred, param)
	assert.Nil(t, err)
	assert.NotEmpty(t, code)
	assert.NotEmpty(t, url)
	log.Print(code)
}

func TestUUID(t *testing.T) {
	u := code.UUID()
	assert.Equal(t, 36, len(u))
}
