package main

import (
	"payment/pkg/config"
	"payment/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	rt := gin.Default()
	router.SetupRoutes(rt)
	rt.Run()
}