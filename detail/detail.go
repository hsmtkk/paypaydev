package detail

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hsmtkk/paypaydev/auth"
	"github.com/hsmtkk/paypaydev/constant"
	"github.com/pkg/errors"
)

type DetailGetter interface {
	GetDetail(auth.Credential, string) (string, error)
}

type detailGetterImpl struct {
	client  *http.Client
	baseURL string
}

func New() DetailGetter {
	return &detailGetterImpl{client: http.DefaultClient, baseURL: "https://stg-api.sandbox.paypay.ne.jp/v2/codes/payments"}
}

func NewForTest(client *http.Client, baseURL string) DetailGetter {
	return &detailGetterImpl{client: client, baseURL: baseURL}
}

func (g *detailGetterImpl) GetDetail(cred auth.Credential, merchantPaymentID string) (string, error) {
	url := g.baseURL + "/" + merchantPaymentID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to create new request")
	}

	authParam := auth.HMACAuthHeaderParameter{
		APIKey:        cred.APIKey,
		APIKeySecret:  cred.APIKeySecret,
		RequestURL:    "/v2/codes/payments/" + merchantPaymentID,
		RequestMethod: http.MethodGet,
	}
	authHeader, err := auth.New().HMACAuthHeader(authParam)
	if err != nil {
		return "", errors.Wrap(err, "failed to create HMAC auth header")
	}
	req.Header.Add(constant.Authorization, authHeader)
	req.Header.Add(constant.XAssumeMerchantHeader, cred.MerchantID)

	resp, err := g.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to send GET request")
	}
	defer resp.Body.Close()
	log.Print(resp.StatusCode)
	log.Print(resp.Status)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response")
	}
	return string(respBytes), nil
}
