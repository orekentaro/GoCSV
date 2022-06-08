package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostCSV(c *gin.Context) {
	header := c.PostForm("header")
	csvData := c.Request.MultipartForm.File["csvData"][0]
	outputType := c.PostForm("outputType")
	fmt.Println(header, csvData, outputType)
	c.JSON(200, gin.H{
		"message": "connect",
	})
}
