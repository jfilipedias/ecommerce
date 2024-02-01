package repository

import (
	"database/sql"

	"github.com/jfilipedias/ecommerce-fullcycle/api/internal/entity"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) Create(category *entity.Category) (string, error) {
	_, err := cr.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}

func (cr *CategoryRepository) GetAll() ([]*entity.Category, error) {
	rows, err := cr.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category

	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (cr *CategoryRepository) GetByID(id string) (*entity.Category, error) {
	var category entity.Category
	err := cr.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
