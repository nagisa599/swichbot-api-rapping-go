package switchbot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// トークンを署名する
func generateSign(token string, secret string, t int64, nonce string) string {
	v := []byte(fmt.Sprintf("%s%d%s", token, t, nonce))
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(v)
	sign := h.Sum(nil)
	result := base64.StdEncoding.EncodeToString(sign)
	return result
}
