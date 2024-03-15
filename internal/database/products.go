package product

import (
	"ecommerce/api/models"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetProduct(id int) (models.Product, error) {
	product := new(models.Product)

	


	return *product, nil
}
