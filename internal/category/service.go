package category

type repository interface {
	FindById(id int) (*Category, error)
	FindAll() (*[]Category, error)
	Create(cat *Category) (*Category, error)
	UpdateDescription(id int, description string) (*Category, error)
	Delete(id int) error
}
type CategoryService struct {
	repo repository
}

func NewCategoryService(repo repository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) FindById(id int) (*CategoryResponse, error) {
	response, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return &CategoryResponse{Id: response.id, Name: response.name, Description: response.description}, nil
}

func (s *CategoryService) FindAll() (*[]CategoryResponse, error) {
	response, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []CategoryResponse
	for _, category := range *response {
		cat := CategoryResponse{Id: category.id, Name: category.name, Description: category.description}
		categories = append(categories, cat)
	}

	return &categories, nil
}

func (s *CategoryService) Create(cat *CategoryRequest) (*CategoryResponse, error) {
	newCat := &Category{
		name:        cat.Name,
		description: cat.Description,
	}

	category, err := s.repo.Create(newCat)
	if err != nil {
		return nil, err
	}
	return &CategoryResponse{
		Id:          category.id,
		Name:        category.name,
		Description: category.description,
	}, nil
}

func (s *CategoryService) UpdateDescription(id int, cat *CategoryRequest) (*CategoryResponse, error) {
	response, err := s.repo.UpdateDescription(id, cat.Description)
	if err != nil {
		return nil, err
	}

	return &CategoryResponse{Id: response.id, Name: response.name, Description: response.description}, nil
}

func (s *CategoryService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
