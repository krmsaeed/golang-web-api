package main

import (
	"github.com/krmsaeed/golang-web-api/api"
	"github.com/krmsaeed/golang-web-api/config"
	"github.com/krmsaeed/golang-web-api/data/cache"
	"github.com/krmsaeed/golang-web-api/data/db"
	"github.com/krmsaeed/golang-web-api/pkg/logging"
)

// / @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
