package auth_test

import (
	"testing"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/stretchr/testify/assert"
)

func TestNonce(t *testing.T) {
	got := auth.NewNonceGetter().Nonce()
	assert.Equal(t, 8, len(got))
}
