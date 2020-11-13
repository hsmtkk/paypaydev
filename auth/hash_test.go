package auth_test

import (
	"testing"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	want := "1j0FnY4flNp5CtIKa7x9MQ=="
	got, err := auth.NewHashCalculator().Hash(exampleParameter.RequestContentType, exampleParameter.RequestBody)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
