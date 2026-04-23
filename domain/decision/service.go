package decision

type Service interface {
	GetDecisionCategories() ([]DecisionCategory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetDecisionCategories() ([]DecisionCategory, error) {
	return s.repository.FindAll()
}
