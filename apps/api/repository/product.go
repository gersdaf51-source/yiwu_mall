package repository


import (
	"database/sql"

	"yiwu-api/model"
)



func GetProducts(db *sql.DB) ([]model.Product,error){


	rows,err:=db.Query(
		"SELECT id,name,description,price,stock,image_url FROM products",
	)


	if err!=nil{
		return nil,err
	}


	defer rows.Close()


	var products []model.Product


	for rows.Next(){

		var p model.Product


		err:=rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.ImageURL,
		)


		if err!=nil{
			return nil,err
		}


		products=append(products,p)

	}


	return products,nil
}