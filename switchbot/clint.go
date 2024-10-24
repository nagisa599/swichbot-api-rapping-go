package switchbot

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const baseURL = "https://api.switch-bot.com"

type SwitchBotClient struct {
	// アクセストークン
	token string

	// シークレットキー
	secret string

	// HTTPリクエストのタイムアウト
	timeout time.Duration
}

func NewSwitchBotClient(token string, secret string, timeout time.Duration) *SwitchBotClient {
	return &SwitchBotClient{
		token:   token,
		secret:  secret,
		timeout: timeout,
	}
}

func (c SwitchBotClient) SendRequest(method string, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		return nil, err
	}
	// トークンを署名する
	t := time.Now().UnixNano() / int64(time.Millisecond)
	nonce := uuid.New().String()
	sign := generateSign(c.token, c.secret, t, nonce)

	// 必須のリクエストヘッダーをセットする
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", c.token)
	req.Header.Set("sign", sign)
	req.Header.Set("t", fmt.Sprint(t))
	req.Header.Set("nonce", nonce)

	client := &http.Client{
		// リソース節約のためにタイムアウトを設定する
		Timeout: c.timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error: bad status %d", resp.StatusCode)
	}
	fmt.Println("status: ", resp.Status)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
