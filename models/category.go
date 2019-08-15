package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 类别表
type Category struct {
	Id          int       `orm:"auto"`
	Name        string    `orm:"unique;column(cate_name);description(类别名称)"`
	ParentId    int       `orm:"column(parent_id);description(父类别id);default(0);"`
	Description string    `orm:"type(text);description(类别描述)"`
	DisplayOrder uint8    `orm:"default(99);description(排序)"`
	Status      uint8     `orm:"default(1);description(用户状态)"`
	UserId      int       `orm:"default(0);column(user_id);description(创建人)"`
	Created     time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated     time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

type Categories []*Category

func (c *Category) Store() (int64, error) {
	var (
		o        = orm.NewOrm()
		err      error
		insertId int64
	)
	insertId, err = o.Insert(c)
	return insertId, err
}

func (c *Category) GetList() ([]*Category, error) {
	var (
		o          = orm.NewOrm()
		err        error
		categories Categories
	)
	_, err = o.QueryTable(c).All(&categories)
	return categories, err
}

func (c *Category) Remove() error {
	var (
		o          = orm.NewOrm()
		err        error
	)
	c.Status = STATUS_BANNED
	_, err = o.Update(c, "Status")
	return err
}

//func (c *Category) GetById(id int) (*Category, error) {
//	var (
//		o          = orm.NewOrm()
//		err        error
//	)
//	c.Id = id
//	err = o.Read(c)
//	if err != nil {
//		if err == orm.ErrNoRows {
//			return nil, errors.New("查询不到")
//		}
//	}
//	return c, _
//}
