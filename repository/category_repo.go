package repository

import (
	"toko_kue/helper"
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

func (r *Repo) IsCodeUnique(code string) (bool, error) {
	data := models.Category{}
	res := r.DB.Where("code = ?", code).First(&data)
	if res.Error == nil {
		return false, nil
	}
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, res.Error
}

func (r *Repo) GetAllCategory() ([]models.Category, error) {
	data := []models.Category{}
	res := r.DB.Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (r *Repo) GetById(id int) (*models.Category, error) {
	data := models.Category{}
	res := r.DB.First(&data, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}

func (r *Repo) AddCategory(cat *models.Category) (*models.Category, error) {
	codeUnique, err := r.IsCodeUnique(cat.Code)
	if err != nil || !codeUnique {
		return nil, helper.ErrCodeExists
	}

	// code is unique
	res := r.DB.Create(cat)
	if res.Error != nil {
		return nil, res.Error
	}

	return cat, nil
}

func (r *Repo) UpdateCategory(cat *models.Category) (*models.Category, error) {
	data := models.Category{}
	res := r.DB.First(&data, cat.ID)
	if res.Error != nil {
		return nil, res.Error
	}

	if data.Code != cat.Code {
		codeUnique, err := r.IsCodeUnique(cat.Code)
		if err != nil || !codeUnique {
			return nil, helper.ErrCodeExists
		}
		data.Code = cat.Code

	}
	data.Name = cat.Name

	res = r.DB.Save(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}

func (r *Repo) DeleteCategory(id int) (*models.Category, error) {
	data := models.Category{}
	res := r.DB.First(&data, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, helper.ErrNoData
		}
		return nil, res.Error
	}
	res = r.DB.Delete(&data)

	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}
