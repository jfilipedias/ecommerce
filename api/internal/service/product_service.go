package service

import (
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/entity"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductRepository(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps *ProductService) Create(name, description string, price float64, categoryID, imageURL string) (string, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	productID, err := ps.ProductRepository.Create(product)
	if err != nil {
		return "", err
	}
	return productID, err
}

func (ps *ProductService) GetAll() ([]*entity.Product, error) {
	products, err := ps.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetByID(id string) (*entity.Product, error) {
	product, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductRepository.GetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
