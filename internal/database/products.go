package database

import (
	"ecommerce/api/models"
)

func (d DB) GetProduct(id int) (*models.Product, error) {
	product := new(models.Product)
	row := d.conn.QueryRow("SELECT * FROM products where id=?", id)
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
		return nil, err
	}

	return product, nil
}

func (d DB) GetProducts() ([]models.Product, error) {
	rows, err := d.conn.Query("SELECT rowid, * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, err
}

func (d DB) Search(term string) ([]models.Product, error) {
	rows, err := d.conn.Query("SELECT * FROM products where name LIKE '%?%'", term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, err
}

func (d DB) DeleteProduct(id uint) (bool, error) {
	_, err := d.conn.Exec("DELETE from products where rowid=?", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
