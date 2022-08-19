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
var PORT int = 8080
var PATH string = "./telser"

func initialize(token string) {
	helper.CreatePath(PATH)
	helper.CreatePath(PATH + "/logs")
	logger.SetLogFilename("./telser/logs/main.log")
	BOT_TOKEN = token
}

func main() {
	initialize("YOUR_TELEGRAM_BOT_TOKEN_GOES_HERE")

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
