package forms

// StoreCategoryForm
type StoreCategoryForm struct {
	CateName    string `form:"cate_name"`
	ParentId    int    `form:"parent_id"`
}
