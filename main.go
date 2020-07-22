package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	data := generateData()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": data,
			"time": time.Now().String(),
		})
	})
	_ = r.Run()
}

func generateData() []gin.H {
	var dataList []gin.H
	for i := 0; i < 5; i++ {
		data := gin.H{
			"name":    gofakeit.Name(),
			"title":   gofakeit.JobTitle(),
			"country": gofakeit.Country(),
			"car":     fmt.Sprintf("%s - %s", gofakeit.CarMaker(), gofakeit.CarModel()),
			"address": gofakeit.Address().Address,
		}
		dataList = append(dataList, data)
	}
	return dataList
}
