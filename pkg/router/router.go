package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
			"data":   os.Getenv("API_KEY"),
			"data2":  os.Getenv("ENV"),
		})
	})

	omise := router.Group("/omise")
	{
		omise.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": "Hello Omise",
			})
		})
	}

}
