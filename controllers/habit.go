package controllers

import (
	. "fish_habits/controllers/base"
	"fish_habits/controllers/forms"
	"fish_habits/services"
	"net/http"
	"strconv"
)

type HabitController struct {
	ApiController
	HabitService *services.HabitService
}

func (c *HabitController) Prepare() {
	c.ApiController.Prepare()
	c.HabitService = services.NewHabitService()  // todo service
}

func (c *HabitController) URLMapping() {
	c.Mapping("Store", c.Store)
	c.Mapping("StoreUserHabit", c.StoreUserHabit)
	c.Mapping("RemoveUserHabit", c.RemoveUserHabit)
	c.Mapping("Sign", c.Sign)
}

// @router / [post]
func (c *HabitController) Store() {
	var (
		f   forms.StoreHabitForm
		err error
		id  int64
	)
	if err = c.ParseForm(&f); err != nil {
		c.JsonReturn("解析参数发生错误:" + err.Error(), "", 500)
		return
	}
	if id, err = c.HabitService.Store(f); err != nil {
		c.JsonReturn("保存习惯接口发生错误:" + err.Error(), "", 500)
	} else{
		c.JsonReturn("保存习惯接口", id, http.StatusOK)
	}
}

// @router /:id/user/:uid [post]
func (c *HabitController) StoreUserHabit()  {
	var (
		habitId int
		userId int
		id int64
		err error
	)
	habitId, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))
	userId, _ = strconv.Atoi(c.Ctx.Input.Param(":uid"))
	if id, err = c.HabitService.StoreUserHabit(habitId, userId); err != nil {
		c.JsonReturn("保存用户习惯接口发生错误:" + err.Error(), "", 500)
	} else {
		c.JsonReturn("保存用户习惯接口", id, http.StatusOK)
	}
}

// @router /:id/user/:uid [delete]
func (c *HabitController) RemoveUserHabit()  {
	var (
		habitId int
		userId int
		id int
		err error
	)
	habitId, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))
	userId, _ = strconv.Atoi(c.Ctx.Input.Param(":uid"))
	if id, err = c.HabitService.RemoveUserHabit(habitId, userId); err != nil {
		c.JsonReturn("删除用户习惯接口发生错误:" + err.Error(), "", 500)
	} else {
		c.JsonReturn("删除用户习惯接口", id, http.StatusOK)
	}
}

// @router /:id/user/:user/sign
func (c *HabitController) Sign()  {
	var (
		habitId int
		userId int
		id int
		err error
	)
}