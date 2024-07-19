package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
			"data":   nil,
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
