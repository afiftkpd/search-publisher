package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"publisher/models"
)

type productRepository struct {
	DB dbCommand
}

type dbCommand interface {
	QueryRow(query string, args ...any) *sql.Row
}

func NewProductRepository(db dbCommand) ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) Get(ctx context.Context, id int64) (models.Product, error) {
	result := models.Product{}
	err := p.DB.QueryRow("SELECT id, name, price, rating, image_url, stock, description FROM products WHERE id = $1", id).Scan(&result.ID, &result.Name, &result.Price, &result.Rating, &result.ImageURL, &result.Stock, &result.Description)
	if err != nil {
		fmt.Println("geterr")
		return result, err
	}

	return result, nil
}

func (p *productRepository) Update(ctx context.Context, product models.Product) error {
	return p.DB.QueryRow(`UPDATE products SET name=$1, price=$2, image_url=$3, rating=$4, stock=$5, description=$6 WHERE id = $6`, product.Name, product.Price, product.ImageURL, product.Rating, product.Stock, product.ID, product.Description).Err()
}

func (p *productRepository) Delete(ctx context.Context, id int64) error {
	return p.DB.QueryRow(`DELETE FROM products WHERE id = $1`, id).Err()
}

func (p *productRepository) Store(ctx context.Context, product models.Product) (int, error) {
	id := 0
	err := p.DB.QueryRow(`INSERT INTO products (name, price, rating, image_url, stock, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, product.Name, product.Price, product.Rating, product.ImageURL, product.Stock, product.Description).Scan(&id)
	return id, err
}
