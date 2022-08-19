package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Pilladian/go-helper"
	"github.com/Pilladian/logger"
)

type TelegramMessage struct {
	ID      string `json:"id"`
	Message string `json:"m"`
}

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

// Response Codes:
//
//	0 : OK
//	801 : POST Request did not contain the Authorization Header
//	802 : Unauthorized
//	803 : Data Validation Failed
//	804 : JSON could not be parsed
//	805 : Message could not be sent
func processRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	auth := r.Header["Authorization"]
	if len(auth) == 0 {
		return 801, errors.New("POST Request did not contain the Authorization Header")
	}

	authorized, authorized_err := authenticate(auth[0])
	if authorized_err != nil {
		return 802, authorized_err
	}

	if authorized {
		data, _ := ioutil.ReadAll(r.Body)

		if data_err := helper.ValidateSimpleJSON(string(data)); data_err != nil {
			return 803, errors.New(fmt.Sprintf("Data validation failed: %s", data_err.Error()))
		}

		var content TelegramMessage
		if json_err := json.Unmarshal([]byte(data), &content); json_err != nil {
			return 804, json_err
		}

		resp_code, resp_err := sendTelegramMessage(content.ID, content.Message)
		if resp_err != nil {
			return 805, errors.New(fmt.Sprintf("Server Response Code: %d - %s", resp_code, resp_err.Error()))
		}
	} else {
		fmt.Fprintf(w, "unauthorized\n")
		return 0, nil
	}
	return 0, nil
}

func sendRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r, err := processRequest(w, r)
		if err != nil {
			logger.Error(fmt.Sprintf("Server Response %d : %s", r, err.Error()))
			fmt.Fprintf(w, "error\n")
			return
		}
		fmt.Fprintf(w, "success\n")
	} else {
		fmt.Fprintf(w, "denied\n")
	}
}
