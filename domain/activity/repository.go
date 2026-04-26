package activity

import "gorm.io/gorm"

type Repository interface {
	Create(activity *Activity) error
	FindAll(boardID string) ([]Activity, error)
	CreateBoard(id string) error
	CheckBoardExists(id string) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(activity *Activity) error {
	return r.db.Create(activity).Error
}

func (r *repository) FindAll(boardID string) ([]Activity, error) {
	var activities []Activity
	err := r.db.Preload("Decision").Where("board_id = ?", boardID).Order("created_at desc").Find(&activities).Error
	return activities, err
}

func (r *repository) CreateBoard(id string) error {
	board := Board{ID: id}
	return r.db.Create(&board).Error
}

func (r *repository) CheckBoardExists(id string) bool {
	var count int64
	r.db.Model(&Board{}).Where("id = ?", id).Count(&count)
	return count > 0 // Kalau > 0 berarti ID asli buatan sistem
}
