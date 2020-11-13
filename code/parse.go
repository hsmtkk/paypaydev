package code

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ResponseSchema struct {
	Data struct {
		CodeID string `json:"codeId"`
		URL    string `json:"url"`
	} `json:"data"`
}

func ParseResponse(bs []byte) (string, string, error) {
	sch := ResponseSchema{}
	if err := json.Unmarshal(bs, &sch); err != nil {
		return "", "", errors.Wrap(err, "failed to unmarshal JSON")
	}
	return sch.Data.CodeID, sch.Data.URL, nil
}
