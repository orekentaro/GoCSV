package models

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
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

	var record string

	r := csv.NewReader(file)
	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		record += strings.Join(row, ",") + "\n"
	}
	fmt.Println(record)

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=test.csv")
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Write([]byte(record))
}
