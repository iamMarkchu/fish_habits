package services

import (
	"errors"
	"fish_habits/controllers/forms"
	"fish_habits/models"
	"strconv"
	"time"
)

type HabitService struct {
}

func NewHabitService() *HabitService {
	return &HabitService{}
}

func (c *HabitService) Store(f forms.StoreHabitForm) (int64, error) {
	var (
		habit    models.Habit
		// category models.Category
	)
	if f.CateId == 0 {
		return 0, errors.New("类别ID需要提供")
	} else {
	//	category.GetById(f.CateId)
	}

	habit.CateId = f.CateId
	habit.Name = f.Name
	if f.DisplayOrder == 0 {
		habit.DisplayOrder = models.DEFAULT_DISPLAY_ORDER
	} else {
		habit.DisplayOrder = f.DisplayOrder
	}
	habit.Status = models.STATUS_NORMAL
	habit.UserId = models.TEST_USER_ID
	return habit.Store()
}

func (c *HabitService) StoreUserHabit(habitId int, userId int) (int64, error) {
	var (
		userHabit models.UserHabit
	)

	userHabit.UserId = userId
	userHabit.HabitId = habitId
	userHabit.Status = models.STATUS_NORMAL
	return userHabit.Store()
}

func (c *HabitService) RemoveUserHabit(habitId int, userId int) (int, error) {
	var (
		userHabit models.UserHabit
		err error
	)

	userHabit.UserId = userId
	userHabit.HabitId = habitId
	if err = userHabit.Fetch(); err != nil {
		return 0, errors.New("该用户查询不到该习惯")
	}
	if _, err = userHabit.Remove(); err != nil {
		return 0, errors.New("删除习惯失败")
	}
	return userHabit.Id, nil
}

func (c *HabitService) Sign(habitId int, userId int) (int64, error) {
	var (
		userHabit models.UserHabit
		sign models.Sign
		err error
	)

	userHabit.UserId = userId
	userHabit.HabitId = habitId
	if err = userHabit.Fetch(); err != nil {
		return 0, errors.New("该用户查询不到该习惯")
	}

	sign.UserHabitId = userHabit.Id
	signDay,_ := strconv.Atoi(time.Now().Format("20060102"))
	sign.SignDay = signDay
	if err = sign.Fetch(); err == nil {
		return 0, errors.New("已签到")
	}

	return sign.Store()
}
