package code_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/paypaydev/code"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	bs, err := ioutil.ReadFile("./response.json")
	assert.Nil(t, err)
	code, url, err := code.ParseResponse(bs)
	assert.Nil(t, err)
	assert.Equal(t, "04-HjgKmkExUyF9vaCz", code)
	assert.Equal(t, "https://qr-stg.sandbox.paypay.ne.jp/28180104HjgKmkExUyF9vaCz", url)
}
