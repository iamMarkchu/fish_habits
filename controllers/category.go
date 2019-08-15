package controllers

import (
	. "fish_habits/controllers/base"
	"fish_habits/controllers/forms"
	"fish_habits/models"
	"fish_habits/services"
	"net/http"
	"strconv"
)

// category controller
type CategoryController struct {
	ApiController
	CategoryService *services.CategoryService
}

func (c *CategoryController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Store", c.Store)
	c.Mapping("Remove", c.Remove)
}

func (c *CategoryController) Prepare()  {
	c.ApiController.Prepare()
	c.CategoryService = services.NewCategoryServices()
}

// @router / [get]
func (c *CategoryController) Index() {
	var categories []*models.Category
	categories, _ = c.CategoryService.GetList()
	c.JsonReturn("获取品类列表接口", categories, http.StatusOK)
}

// @router / [post]
func (c *CategoryController) Store() {
	f := forms.StoreCategoryForm{}
	if err:= c.ParseForm(&f); err != nil {
		c.JsonReturn("解析参数发生错误:" + err.Error(), "", 500)
		return
	}
	if id, err := c.CategoryService.Store(f); err != nil {
		c.JsonReturn("保存类别发生错误:" + err.Error(), "", 500)
	} else {
		c.JsonReturn("保存类别接口", id, http.StatusOK)
	}
}

// @router /:id [delete]
func (c *CategoryController) Remove()  {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err := c.CategoryService.Remove(id); err != nil {
		c.JsonReturn("删除类别失败:" + err.Error(), "", 500)
	} else {
		c.JsonReturn("删除类别接口", id, http.StatusOK)
	}
}