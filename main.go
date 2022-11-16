package main

import (
	"balance-api/pkg/api"
	"balance-api/pkg/config"
	"balance-api/pkg/store"
	"log"
	"net/http"
)

func main() {

	api.RouteHandlers()

	store.ConnectDataBase()

	log.Fatalln(http.ListenAndServe(config.GetConfig().Host+config.GetConfig().Port, nil))
}
