package listing

type Service interface {
	Create(listing *Listing) error
	Update(listing *Listing) error
	Delete(id uint) error
	GetByID(id uint) (*Listing, error)
	GetAll() ([]Listing, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(listing *Listing) error {
	return s.repo.Create(listing)
}
func (s *service) Update(listing *Listing) error {
	return s.repo.Update(listing)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) GetByID(id uint) (*Listing, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetAll() ([]Listing, error) {
	return s.repo.FindAll()
}
