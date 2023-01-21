package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shop"
)

type ProductsListPostgres struct {
	db *sqlx.DB
}

func NewProductListPostgres(db *sqlx.DB) *ProductsListPostgres {
	return &ProductsListPostgres{db: db}
}

func (r *ProductsListPostgres) CreateProduct(product shop.Product) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (title, price, description, category, image, rating) "+
			"values ($1, $2, $3, $4, decode($5, 'base64'), $6) RETURNING id", productsTable)
	row := r.db.QueryRow(query, product.Title, product.Price, product.Description, product.Category, product.Image, product.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductsListPostgres) GetAllProducts() ([]shop.Product, error) {
	var products []shop.Product
	query := fmt.Sprintf("SELECT id, title, price, description, category, translate(encode(image, 'base64'), E'\n', '') as image, rating FROM %s", productsTable)
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductsListPostgres) GetProductById(id int) (shop.Product, error) {
	var product shop.Product
	query := fmt.Sprintf("SELECT id, title, price, description, category, translate(encode(image, 'base64'), E'\n', '') as image, rating FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, id)
	return product, err
}
