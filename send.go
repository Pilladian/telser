package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Response Codes:
//
//	0 : OK
//	1 : Host does not exist
//	200-500 : HTTP Codes
func sendTelegramMessage(chatID string, message string) (int, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=%s", BOT_TOKEN, chatID, message)
	resp, resp_error := http.Get(url)
	if resp_error != nil {
		return 1, errors.New(resp_error.Error())
	}
	return resp.StatusCode, nil
}

func sendRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello There!</h1><p>This is the API page.</p>")
}
