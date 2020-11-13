package auth_test

import (
	"testing"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/stretchr/testify/assert"
)

var exampleParameter = auth.HMACAuthHeaderParameter{
	APIKey:             "APIKeyGenerated",
	APIKeySecret:       "APIKeySecretGenerated",
	RequestURL:         "/v2/codes",
	RequestMethod:      "POST",
	RequestBody:        `{"sampleRequestBodyKey1":"sampleRequestBodyValue1","sampleRequestBodyKey2":"sampleRequestBodyValue2"}`,
	RequestContentType: "application/json;charset=UTF-8;",
}

type epochGetterStub struct{}

func (e *epochGetterStub) Epoch() int64 {
	return 1579843452
}

type nonceGetterStub struct{}

func (n *nonceGetterStub) Nonce() string {
	return "acd028"
}

type hashCalculatorStub struct{}

func (h *hashCalculatorStub) Hash(string, string) (string, error) {
	return "1j0FnY4flNp5CtIKa7x9MQ==", nil
}

func TestHMACAuthHeader(t *testing.T) {
	getter := auth.NewForTest(&epochGetterStub{}, &nonceGetterStub{}, &hashCalculatorStub{})
	want := `hmac OPA-Auth:APIKeyGenerated:NW1jKIMnzR7tEhMWtcJcaef+nFVBt7jjAGcVuxHhchc=:acd028:1579843452:1j0FnY4flNp5CtIKa7x9MQ==`
	got, err := getter.HMACAuthHeader(exampleParameter)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestMAC(t *testing.T) {
	want := "NW1jKIMnzR7tEhMWtcJcaef+nFVBt7jjAGcVuxHhchc="
	got, err := auth.MAC(exampleParameter, 1579843452, "acd028", "1j0FnY4flNp5CtIKa7x9MQ==")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
