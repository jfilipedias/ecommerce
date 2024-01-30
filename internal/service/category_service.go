package service

import (
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/database"
	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (cs *CategoryService) Create(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryDB.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetAll() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetByID(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
