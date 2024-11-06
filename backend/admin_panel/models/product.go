package models

import "admin_panel/config"

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   string `json:"createdAt"`
}

func (p *Product) CreateProduct() error {
	query := "INSERT INTO products (name, price, description, image, created_at) VALUES (?, ?, ?, ?, ?)"
	err := config.DB.QueryRow(query, p.Name, p.Price, p.Description, p.Image, p.CreatedAt).Scan(&p.ID)

	return err
}

func GetAllProducts() ([]Product, error) {
	rows, err := config.DB.Query("SELECT id, name, description, price, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
