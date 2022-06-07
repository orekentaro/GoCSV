package models

import "github.com/gin-gonic/gin"

func PostCSV(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "connect",
	})
}
