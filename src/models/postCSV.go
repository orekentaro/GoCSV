package models

import (
	"encoding/csv"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func PostCSV(c *gin.Context) {
	header := c.PostForm("header")
	csvData := c.Request.MultipartForm.File["csvData"][0]
	outputType := c.PostForm("outputType")
	fmt.Println(header, csvData, outputType)

	file, err := csvData.Open()
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}

	r := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	for {
		records, err := r.Read()
		if err != nil {
			break
		}
		fmt.Println(records)
	}
	c.JSON(200, gin.H{
		"message": "connect",
	})
}
