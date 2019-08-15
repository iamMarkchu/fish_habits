package services

import (
	"fish_habits/controllers/forms"
	"fish_habits/models"
)

type CategoryService struct {
}

func (c *CategoryService) Store(f forms.StoreCategoryForm) (int64, error) {
	var (
		category models.Category
	)

	category.Name = f.CateName
	category.ParentId = f.ParentId
	if f.DisplayOrder == 0 {
		category.DisplayOrder = models.DEFAULT_DISPLAY_ORDER
	} else {
		category.DisplayOrder = f.DisplayOrder
	}
	category.UserId = models.TEST_USER_ID
	category.Status = models.STATUS_NORMAL

	id, err := category.Store()
	return id, err
}

func (c *CategoryService) GetList() ([]*models.Category, error) {
	var (
		category   models.Category
		categories []*models.Category
		err        error
	)

	categories, err = category.GetList()
	return categories, err
}

func (c *CategoryService) Remove(id int) error {
	var (
		category   models.Category
	)
	category.Id = id
	return category.Remove()
}

func NewCategoryServices() *CategoryService {
	return &CategoryService{}
}
