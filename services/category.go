package services

import (
	"fish_habits/controllers/forms"
	"fish_habits/models"
)

type CategoryService struct {
}

func (c *CategoryService) Store(f forms.StoreCategoryForm) (int64, error) {
	category := models.Category{}
	category.Name = f.CateName
	category.ParentId = f.ParentId
	category.UserId = 1
	category.Status = 1
	id, err := category.Store()
	return id, err
}

func (c *CategoryService) GetList() ([]*models.Category, error) {
	category := models.Category{}
	var categories []*models.Category
	categories, err := category.GetList()
	return categories, err
}

func NewCategoryServices() *CategoryService {
	return &CategoryService{}
}
