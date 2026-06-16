package handler


import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"yiwu-api/repository"
)



func GetProducts(db *sql.DB) gin.HandlerFunc {


	return func(c *gin.Context){


		products,err:=repository.GetProducts(db)


		if err!=nil{

			c.JSON(500,gin.H{
				"error":err.Error(),
			})

			return
		}


		c.JSON(http.StatusOK,products)

	}

}