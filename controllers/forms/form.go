package forms

// StoreCategoryForm
type StoreCategoryForm struct {
	CateName     string `form:"cate_name"`
	ParentId     int    `form:"parent_id"`
	DisplayOrder uint8  `form:"display_order"`
}

// StoreHabitForm
type StoreHabitForm struct {
	CateId       int    `form:"cate_id"`
	Name         string `form:"name"`
	DisplayOrder uint8  `form:"display_order"`
}
