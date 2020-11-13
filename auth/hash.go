package auth

import (
	"crypto/md5"
	"io"

	"github.com/pkg/errors"
)

type HashCalculator interface {
	Hash(string, string) (string, error)
}

type hashCalculatorImpl struct{}

func NewHashCalculator() HashCalculator {
	return &hashCalculatorImpl{}
}

func (h *hashCalculatorImpl) Hash(requestContentType, requestBody string) (string, error) {
	m := md5.New()
	if _, err := io.WriteString(m, requestContentType); err != nil {
		return "", errors.Wrap(err, "failed to write string")
	}
	if _, err := io.WriteString(m, requestBody); err != nil {
		return "", errors.Wrap(err, "failed to write string")
	}
	return base64encode(m.Sum(nil)), nil
}
