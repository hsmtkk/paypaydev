package code

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hsmtkk/paypaydev/auth"
	"github.com/hsmtkk/paypaydev/constant"
	"github.com/pkg/errors"
)

type CodeCreator interface {
	CreateCode(auth.Credential, CreateCodeParameter) (string, string, error)
}

type codeCreatorImpl struct {
	client *http.Client
	url    string
}

type CreateCodeParameter struct {
	MerchantPaymentID string      `json:"merchantPaymentId"`
	Amount            Amount      `json:"amount"`
	OrderItems        []OrderItem `json:"orderItems"`
	CodeType          string      `json:"codeType"`
}

func NewCreateCodeParameter(items []OrderItem) CreateCodeParameter {
	amount := 0
	for _, item := range items {
		amount += item.UnitPrice.Amount * item.Quantity
	}
	return CreateCodeParameter{
		MerchantPaymentID: UUID(),
		Amount: Amount{
			Amount:   amount,
			Currency: constant.JPY,
		},
		OrderItems: items,
		CodeType:   constant.OrderQR,
	}
}

type OrderItem struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Quantity  int    `json:"quantity"`
	ProductID string `json:"productId"`
	UnitPrice Amount `json:"unitPrice"`
}

func NewOrderItem(name, category string, quantity int, productID string, unitPrice int) OrderItem {
	return OrderItem{
		Name:      name,
		Category:  category,
		Quantity:  quantity,
		ProductID: productID,
		UnitPrice: Amount{
			Amount:   unitPrice,
			Currency: constant.JPY,
		},
	}
}

type Amount struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func New() CodeCreator {
	return &codeCreatorImpl{client: http.DefaultClient, url: "https://stg-api.sandbox.paypay.ne.jp/v2/codes"}
}

func NewForTest(client *http.Client, url string) CodeCreator {
	return &codeCreatorImpl{client: client, url: url}
}

func (c *codeCreatorImpl) CreateCode(cred auth.Credential, param CreateCodeParameter) (string, string, error) {
	bodyBytes, err := json.Marshal(&param)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to marshal JSON")
	}
	authParam := auth.HMACAuthHeaderParameter{
		APIKey:             cred.APIKey,
		APIKeySecret:       cred.APIKeySecret,
		RequestURL:         "/v2/codes",
		RequestMethod:      http.MethodPost,
		RequestBody:        string(bodyBytes),
		RequestContentType: constant.ApplicationJSON,
	}
	authHeader, err := auth.New().HMACAuthHeader(authParam)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to create HMAC auth header")
	}
	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", "", errors.Wrap(err, "failed to make new request")
	}
	req.Header.Add(constant.Authorization, authHeader)
	req.Header.Add(constant.ContentType, constant.ApplicationJSON)
	req.Header.Add(constant.XAssumeMerchantHeader, cred.MerchantID)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to post")
	}
	defer resp.Body.Close()
	log.Print(resp.StatusCode)
	log.Print(resp.Status)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to read request")
	}
	code, url, err := ParseResponse(respBytes)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to parse response")
	}
	return code, url, nil
}

func UUID() string {
	return uuid.New().String()
}
