package auth

import "encoding/base64"

func base64encode(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}
