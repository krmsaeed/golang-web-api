package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saeedKarami/golang-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Working")
			return
		})
	}
	// server := &http.Server{
	// 	Handler:     r,
	// 	Addr:        fmt.Sprintf(":%s", cfg.Server.ExternalPort),
	// 	ReadTimeout: time.Second * 10,
	// }
	// server.ListenAndServe()
	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}
