package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yiwu-api/database"

	"yiwu-api/handler"
)


func main() {


	db := database.Connect()

	defer db.Close()


	r := gin.Default()


	r.GET("/ping", func(c *gin.Context){

		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})

	})

		// 商品接口
	r.GET("/products", handler.GetProducts(db))


	r.Run(":8080")

}