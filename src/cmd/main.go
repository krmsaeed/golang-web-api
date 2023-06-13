package main

import (
	"github.com/krmsaeed/golang-web-api/api"
	"github.com/krmsaeed/golang-web-api/config"
	"github.com/krmsaeed/golang-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	cache.InitialRedis(cfg)

}
