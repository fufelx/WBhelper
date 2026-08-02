package wb

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) ChatInit() {
	next := "0"
next:
	req, err := http.NewRequest("GET", "https://buyer-chat-api.wildberries.ru/api/v1/seller/events?next="+next, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var chatevent ChatEvent
	err = json.Unmarshal(body, &chatevent)
	if err != nil {
		return
	}

	// выгрузка чатов в тг
	log.Println(chatevent)

	if chatevent.Result.TotalEvents != 0 {
		next = strconv.FormatInt(chatevent.Result.Next, 10)
		time.Sleep(time.Second * 3) // Добавляем задержку перед следующим запросом
		goto next
	}
}
