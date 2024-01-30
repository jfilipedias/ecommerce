package repository

import (
	"database/sql"

	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(product *entity.Product) (string, error) {
	_, err := pr.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES(?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}

func (pr *ProductRepository) GetAll() ([]*entity.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pr *ProductRepository) GetByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := pr.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) GetByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
