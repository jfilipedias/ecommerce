package product

type ProductService struct {
	ProductRepository ProductRepository
}

func NewService(productRepository ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps *ProductService) Create(name, description string, price float64, categoryID, imageURL string) (*Product, error) {
	product := NewProduct(name, description, price, categoryID, imageURL)
	_, err := ps.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (ps *ProductService) GetAll() ([]*Product, error) {
	products, err := ps.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetByID(id string) (*Product, error) {
	product, err := ps.ProductRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetByCategoryID(categoryID string) ([]*Product, error) {
	products, err := ps.ProductRepository.GetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
