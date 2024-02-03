package service

import (
	"github.com/jfilipedias/ecommerce/api/internal/entity"
	"github.com/jfilipedias/ecommerce/api/internal/repository"
)

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

func (cs *CategoryService) Create(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryRepository.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetAll() ([]*entity.Category, error) {
	categories, err := cs.CategoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetByID(id string) (*entity.Category, error) {
	category, err := cs.CategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
