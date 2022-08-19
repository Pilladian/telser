package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Pilladian/go-helper"
	"github.com/Pilladian/logger"
)

// GLOBAL VARIABLES
var BOT_TOKEN string
var AUTH_USERS map[string]string
var PORT int = 8080
var PATH string = "./telser"

func initialize() {
	helper.CreatePath(PATH)
	helper.CreatePath(PATH + "/logs")
	logger.SetLogFilename("./telser/logs/main.log")

	// Set Bot Token
	BOT_TOKEN = "YOUR_TELEGRAM_BOT_TOKEN_GOES_HERE"

	// add authorized users
	AUTH_USERS = make(map[string]string)
	AUTH_USERS["AUTH_USER_1"] = "AUTH_USER_1_PASSWORD"
}

func main() {
	initialize()

	// http request handler
	http.HandleFunc("/", rootRequestHandler)
	http.HandleFunc("/api/v1/send", sendRequestHandler)

	// start web server
	server_err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

	// handle web server errors
	if errors.Is(server_err, http.ErrServerClosed) {
		logger.Fatal("web server closed\n")
	} else if server_err != nil {
		logger.Fatal(fmt.Sprintf("error starting web server: %s\n", server_err))
	}
}
