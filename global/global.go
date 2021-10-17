package global

import (
	"net/http"
)

var GlobalTransport *http.Transport

func init() {
	GlobalTransport = &http.Transport{
		MaxIdleConns: 100,
	}
}
