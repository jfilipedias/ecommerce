package service

import (
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/database"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductDB(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) Create(name, description string, price float64, categoryID, imageURL string) (string, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	productID, err := ps.ProductDB.Create(product)
	if err != nil {
		return "", err
	}
	return productID, err
}

func (ps *ProductService) GetAll() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetByID(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
