package main

import (
	"fmt"
	"time"

	"github.com/nagisa599/swichbot-api-rapping-go/switchbot"
)

func main() {
	c := switchbot.NewSwitchBotClient("dba03e82b2fc44a12dc3769c306bdad5cfd106685bb6830382f6215a1c1e05e98e714e015c7e86786f428f7bfb1b41ef", "b679a44409354167386b0140ff08f450", time.Second * 1000,)
	// body := []byte(`{"action":"updateWebhook","url":"https://7fjt7qzok2.execute-api.ap-northeast-1.amazonaws.com/production/hooks","deviceList":"ALL"}`)
	// body := []byte(`{"action":"updateWebhook","config":{"url":"https://7fjt7qzok2.execute-api.ap-northeast-1.amazonaws.com/production/hooks","enable":true}}`)
	// resp, err := c.SendRequest("POST", "/v1.1/webhook/updateWebhook", bytes.NewBuffer(body))

	/* ----デバイス一覧を取得する関数---- */
	resp, err := c.SendRequest("GET", "/v1.1/devices", nil)


	/* ----webhookを登録する関数---- */
	// body := []byte(`{"action":"queryUrl"}`)
	// resp, err := c.SendRequest("POST", "/v1.1/webhook/queryWebhook", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	println(string(resp))
}
