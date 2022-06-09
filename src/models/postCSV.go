package models

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostCSV(c *gin.Context) {
	header := strings.Split(c.PostForm("header"), ",")
	csvData := c.Request.MultipartForm.File["csvData"][0]
	outputType := c.PostForm("outputType")

	file, err := csvData.Open()
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}

	var record string

	r := csv.NewReader(file)

	csvHeader, err := r.Read()
	if err != nil {
		c.JSON(200, gin.H{
			"message": err,
		})
	}

	indexList := []int{}

	for index, value := range csvHeader {
		if include(header, value) {
			indexList = append(indexList, index)
		}
	}

	fmt.Println(outputType)
	fmt.Println(csvHeader)
	fmt.Println(header)
	fmt.Println(indexList)
	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		record += strings.Join(row, ",") + "\n"
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=test.csv")
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Write([]byte(record))
}

func include(slice []string, target string) bool {
	for _, value := range slice {
		if value == target || "\ufeff"+value == target {
			return true
		}
	}
	return false
}
