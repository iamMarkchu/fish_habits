package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 习惯表
type Habit struct {
	Id           int       `orm:"auto"`
	CateId       int       `orm:"description(类别id)"`
	Name         string    `orm:"unique;description(习惯名称)"`
	Status       uint8     `orm:"default(1);description(用户状态)"`
	DisplayOrder uint8     `orm:"default(99);description(排序)"`
	UserId       int       `orm:"default(0);column(user_id);description(创建人)"`
	Created      time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated      time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func (m *Habit) Store() (int64, error) {
	var (
		o = orm.NewOrm()
	)
	return o.Insert(m)
}
