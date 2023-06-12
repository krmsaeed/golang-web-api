package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saeedKarami/golang-web-api/api/routers"
	"github.com/saeedKarami/golang-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.TestRouter(test_router)
		routers.Health(health)
	}
	v2 := r.Group("/api/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)

	}
	r.Run(fmt.Sprintf("%s", cfg.Server.Port))
}
