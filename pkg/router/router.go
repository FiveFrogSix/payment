package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mode": os.Getenv("ENV"),
			"msg":  http.StatusText(http.StatusOK),
		})
	})

	omise := router.Group("/omise")
	{
		omise.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "OMISE",
			})
		})
	}

}
