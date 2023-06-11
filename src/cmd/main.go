package main

import (
	"github.com/saeedKarami/golang-web-api/api"
	"github.com/saeedKarami/golang-web-api/config"
	"github.com/saeedKarami/golang-web-api/data/cache"
	"github.com/saeedKarami/golang-web-api/data/db"
	"github.com/saeedKarami/golang-web-api/data/db/migrations"
	"github.com/saeedKarami/golang-web-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
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
	migrations.Up_1()

	api.InitServer(cfg)
}
