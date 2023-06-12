package main

import (
	"github.com/saeedKarami/golang-web-api/api"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	api.InitServer()
}
