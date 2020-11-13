package auth_test

import (
	"testing"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/stretchr/testify/assert"
)

func TestEpoch(t *testing.T) {
	got := auth.NewEpochGetter().Epoch()
	assert.Greater(t, got, int64(0))
}
