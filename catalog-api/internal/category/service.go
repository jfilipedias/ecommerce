package category

type CategoryService struct {
	CategoryRepository CategoryRepository
}

func NewService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

func (cs *CategoryService) Create(name string) (*Category, error) {
	category := NewCategory(name)
	_, err := cs.CategoryRepository.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetAll() ([]*Category, error) {
	categories, err := cs.CategoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetByID(id string) (*Category, error) {
	category, err := cs.CategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
