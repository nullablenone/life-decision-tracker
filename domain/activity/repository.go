package activity

import "gorm.io/gorm"

type Repository interface {
	Create(activity *Activity) error
	FindAll() ([]Activity, error)
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

func (r *repository) FindAll() ([]Activity, error) {
	var activities []Activity
	err := r.db.Preload("Decision").Find(&activities).Error
	return activities, err
}
