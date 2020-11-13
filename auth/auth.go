package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"strconv"

	"github.com/pkg/errors"
)

type HMACAuthHeaderParameter struct {
	APIKey             string
	APIKeySecret       string
	RequestURL         string
	RequestMethod      string
	RequestBody        string
	RequestContentType string
}

type HMACAuthHeaderGenerator interface {
	HMACAuthHeader(HMACAuthHeaderParameter) (string, error)
}

type hmacAuthHeaderGeneratorImpl struct {
	epochGetter    EpochGetter
	nonceGetter    NonceGetter
	hashCalculator HashCalculator
}

func New() HMACAuthHeaderGenerator {
	return &hmacAuthHeaderGeneratorImpl{epochGetter: NewEpochGetter(), nonceGetter: NewNonceGetter(), hashCalculator: NewHashCalculator()}
}

func NewForTest(epochGetter EpochGetter, nonceGetter NonceGetter, hashCalculator HashCalculator) HMACAuthHeaderGenerator {
	return &hmacAuthHeaderGeneratorImpl{epochGetter: epochGetter, nonceGetter: nonceGetter, hashCalculator: hashCalculator}
}

func (h *hmacAuthHeaderGeneratorImpl) HMACAuthHeader(param HMACAuthHeaderParameter) (string, error) {
	epoch := h.epochGetter.Epoch()
	nonce := h.nonceGetter.Nonce()
	hash, err := h.hashCalculator.Hash(param.RequestContentType, param.RequestBody)
	if err != nil {
		return "", errors.Wrap(err, "failed to compute hash")
	}
	mac, err := MAC(param, epoch, nonce, hash)
	if err != nil {
		return "", errors.Wrap(err, "failed to compute MAC")
	}
	authHeader := "hmac OPA-Auth:" + param.APIKey + ":" + mac + ":" + nonce + ":" + strconv.FormatInt(epoch, 10) + ":" + hash
	return authHeader, nil
}

func MAC(param HMACAuthHeaderParameter, epoch int64, nonce, hash string) (string, error) {
	params := [][]byte{[]byte(param.RequestURL), []byte(param.RequestMethod), []byte(nonce), []byte(strconv.FormatInt(epoch, 10)), []byte(param.RequestContentType), []byte(hash)}
	joined := bytes.Join(params, []byte("\n"))
	mac := hmac.New(sha256.New, []byte(param.APIKeySecret))
	mac.Write(joined)
	return base64encode(mac.Sum(nil)), nil
}
