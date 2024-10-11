package repository

import (
	"toko_kue/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type CategoryRepoI interface {
	GetAllCategory() ([]models.Category, error)
	GetById(id int) (*models.Category, error)
	AddCategory(cat *models.Category) (*models.Category, error)
	UpdateCategory(cat *models.Category) (*models.Category, error)
	DeleteCategory(id int) (*models.Category, error)
}

func (r *Repo) GetAllCategory() ([]models.Category, error) {
	return nil, nil
}

func (r *Repo) GetById(id int) (*models.Category, error) {
	return nil, nil
}

func (r *Repo) AddCategory(cat *models.Category) (*models.Category, error) {
	return nil, nil
}

func (r *Repo) UpdateCategory(cat *models.Category) (*models.Category, error) {
	return nil, nil
}

func (r *Repo) DeleteCategory(id int) (*models.Category, error) {
	return nil, nil
}
