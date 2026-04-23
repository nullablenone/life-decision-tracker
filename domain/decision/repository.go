package decision

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]DecisionCategory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]DecisionCategory, error) {
	var categories []DecisionCategory
	err := r.db.Find(&categories).Error
	return categories, err
}
