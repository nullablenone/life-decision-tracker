package activity

type Service interface {
	AddActivity(boardID string, title string, decisionID int) error
	GetActivities(boardID string) ([]Activity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) AddActivity(boardID string, title string, decisionID int) error {
	activity := &Activity{
		BoardID:    boardID,
		Title:      title,
		DecisionID: decisionID,
	}
	return s.repo.Create(activity)
}

func (s *service) GetActivities(boardID string) ([]Activity, error) {
	return s.repo.FindAll(boardID)
}
