package main

import (
	"Memorandum/cache"
	"Memorandum/config"
	"Memorandum/routes"
)

// hertz-swagger middleware
// swagger embed files

// @title Memorandum API
// @version 1.16.1
// @description This is a sample Server pets
// @name penQee
// @BasePath /api/v1
// @host localhost:8080
func main() { // http://localhost:8080/swagger/index.html
	config.Init()
	cache.Redis()
	r := routes.NewRouter()
	r.Spin()
}
