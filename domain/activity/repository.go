package activity

import "gorm.io/gorm"

type Repository interface {
	Create(activity *Activity) error
	FindAll(boardID string) ([]Activity, error)
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
