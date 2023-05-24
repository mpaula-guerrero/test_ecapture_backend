package main

import (
	"test_ecapture_backend/api"
	"test_ecapture_backend/internal/env"
)

func main() {

	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
