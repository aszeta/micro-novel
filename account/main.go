package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aszeta/micro-novel/account/account"
	"github.com/aszeta/micro-novel/account/config"
	"github.com/go-kit/kit/log"
)

func main() {
	config := config.ProvideConfig("app.yaml")
	fmt.Println(config.App.Name)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", config.App.Port, "caller", log.DefaultCaller)
	}

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	r := account.NewHttpServer(account.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", config.App.Port)
	logger.Log("err", http.ListenAndServe(config.App.Port, r))
}
