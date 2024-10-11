package repository

import (
	"toko_kue/helper"
	"toko_kue/models"

	"gorm.io/gorm"
)

type ProductRepoI interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id int) (*models.Product, error)
	AddProduct(cat *models.Product) (*models.Product, error)
	UpdateProduct(cat *models.Product) (*models.Product, error)
	DeleteProduct(id int) (*models.Product, error)
}

func (r *Repo) IsCodeProductUnique(code string) (bool, error) {
	data := models.Product{}
	res := r.DB.Where("code = ?", code).First(&data)
	if res.Error == nil {
		return false, nil
	}
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, res.Error
}

func (r *Repo) GetAllProduct() ([]models.Product, error) {
	data := []models.Product{}
	res := r.DB.Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

func (r *Repo) GetProductById(id int) (*models.Product, error) {
	data := models.Product{}
	res := r.DB.First(&data, id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, helper.ErrNoData
		}
		return nil, res.Error
	}
	return &data, nil
}

func (r *Repo) AddProduct(cat *models.Product) (*models.Product, error) {
	codeUnique, err := r.IsCodeUnique(cat.Code)
	if err != nil || !codeUnique {
		return nil, helper.ErrCodeExists
	}

	// check if id for category is exist
	categoryDB := models.Category{ID: cat.CategoryID}
	res := r.DB.Find(&categoryDB)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, helper.ErrNoData
		}
		return nil, res.Error
	}

	// code is unique
	res = r.DB.Create(cat)
	if res.Error != nil {
		return nil, res.Error
	}

	return cat, nil
}

func (r *Repo) UpdateProduct(cat *models.Product) (*models.Product, error) {
	data := models.Product{}
	res := r.DB.First(&data, cat.ID)
	if res.Error != nil {
		return nil, res.Error
	}

	// check if id for category is exist
	categoryDB := models.Category{ID: cat.CategoryID}
	res = r.DB.Find(&categoryDB)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, helper.ErrNoData
		}
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
	data.CategoryID = cat.CategoryID
	data.Stock = cat.Stock
	data.Price = cat.Price

	res = r.DB.Save(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return &data, nil
}

func (r *Repo) DeleteProduct(id int) (*models.Product, error) {
	data := models.Product{}
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
