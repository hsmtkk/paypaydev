package auth

import (
	"math/rand"
	"strings"
)

type NonceGetter interface {
	Nonce() string
}

type nonceGetterImpl struct{}

func NewNonceGetter() NonceGetter {
	return &nonceGetterImpl{}
}

func (n *nonceGetterImpl) Nonce() string {
	var buf strings.Builder
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 8; i++ {
		n := rand.Intn(len(chars))
		buf.WriteByte(chars[n])
	}
	return buf.String()
}
