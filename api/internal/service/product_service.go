package service

import (
	"github.com/jfilipedias/ecommerce/api/internal/entity"
	"github.com/jfilipedias/ecommerce/api/internal/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps *ProductService) Create(name, description string, price float64, categoryID, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := ps.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return product, err
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
