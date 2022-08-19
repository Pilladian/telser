# TelSer
Telegram Service written in GO

## Installation
```bash
# Clone the repository from Github
git clone https://github.com/Pilladian/telser.git
```

## Setup
Before you run TelSer you have to replace `YOUR_TELEGRAM_BOT_TOKEN_GOES_HERE` with your bot token in `telser.go`:
```go
// telser.go

func initialize() {
	...

	// Set Bot Token
	BOT_TOKEN = "YOUR_TELEGRAM_BOT_TOKEN_GOES_HERE"

	...
}
```
After that you need to define the **authorized users** in `telser.go`. By setting `username` and `password` you specify which credentials need to be included in the API requests.
```go
// telser.go

func initialize() {
	...

	// add authorized users
	AUTH_USERS = make(map[string]string)
	AUTH_USERS["AUTH_USER_1"] = "AUTH_USER_1_PASSWORD"
}
```

Now you can build and run TelSer with the following commands:
```bash
# Build the application
go build -o telser

# Run the application : Default listening port is 8080
./telser
```

## Usage
To use TelSer to send Telegram Messages you simply perform a HTTP request. Make sure to use `POST` requests with `valid json data`. Otherwise TelSer will refuse your request. You also need to provide `valid credentials` (base64 encoded) using the `Authorization` header. 

An example request made with `cURL` would look like this:

`curl http://sub.domain.tld:8080/api/v1/send -X POST -d '{"id": "123456789", "m": "Hello, this is an example."}' -H "Authorization: Basic QVVUSF9VU0VSXzE6QVVUSF9VU0VSXzFfUEFTU1dPUkQK"`