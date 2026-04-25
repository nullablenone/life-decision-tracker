package activity

type Service interface {
	AddActivity(title string, decisionID int) error
	GetActivities() ([]Activity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) AddActivity(title string, decisionID int) error {
	activity := &Activity{
		Title:      title,
		DecisionID: decisionID,
	}
	return s.repo.Create(activity)
}

func (s *service) GetActivities() ([]Activity, error) {
	return s.repo.FindAll()
}
