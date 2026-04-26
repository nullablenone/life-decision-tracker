package activity

import (
	"crypto/rand"
	"encoding/hex"
)

type Service interface {
	AddActivity(boardID string, title string, decisionID int) error
	GetActivities(boardID string) ([]Activity, error)
	GenerateNewBoard() (string, error)
	IsBoardValid(id string) bool
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

func (s *service) GenerateNewBoard() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	boardID := hex.EncodeToString(bytes)

	err := s.repo.CreateBoard(boardID) // Daftarin ke DB
	return boardID, err
}

func (s *service) IsBoardValid(id string) bool {
	return s.repo.CheckBoardExists(id)
}
