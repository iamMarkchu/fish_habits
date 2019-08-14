package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 类别表
type Category struct {
	Id          int       `orm:"auto"`
	Name        string    `orm:"unique;column(cate_name);description(类别名称)"`
	ParentId    int       `orm:"column(parent_id);description(父类别id);default(0);"`
	Description string    `orm:"type(text);description(类别描述)"`
	Status      uint8     `orm:"default(1);description(用户状态)"`
	UserId      int       `orm:"default(0);column(user_id);description(创建人)"`
	Created     time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated     time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

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
		categories []*Category
	)
	_, err = o.QueryTable(c).All(&categories)
	return categories, err
}

// 习惯表
type Habit struct {
	Id      int       `orm:"auto"`
	Name    string    `orm:"unique;description(习惯名称)"`
	Status  uint8     `orm:"default(1);description(用户状态)"`
	UserId  int       `orm:"default(0);column(user_id);description(创建人)"`
	Created time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

// 用户习惯表
type UserHabit struct {
	Id      int       `orm:"auto"`
	UserId  int       `orm:"description(用户ID)"`
	HabitId int       `orm:"description(习惯ID)"`
	Status  uint8     `orm:"default(1);description(状态)"`
	Created time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

// 打卡表
type Sign struct {
	Id          int       `orm:"auto"`
	UserHabitId int       `orm:"column(uh_id);description(用户习惯关系表ID)"`
	SignDay     uint16    `orm:"description(签到日期)"`
	Created     time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated     time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

// 用户表
type User struct {
	Id       int          `orm:"auto"`
	Name     string       `orm:"unique;description(用户名)"`
	NickName string       `orm:"default();description(昵称)"`
	Password string       `orm:"default();description(密码)"`
	Status   uint8        `orm:"default(1);description(用户状态)"`
	Created  time.Time    `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated  time.Time    `orm:"auto_now;column(updated_at);type(datetime)"`
}

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	go CheckError(err, "[RegisterDriver Error]")
	err = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	go CheckError(err, "[RegisterDataBase Error]")
	orm.RegisterModel(new(User), new(Category), new(Habit), new(UserHabit), new(Sign))
	err = orm.RunSyncdb("default", true, true)
	go CheckError(err, "[RunSyncdb Error]")
	orm.Debug, err = beego.AppConfig.Bool("ormdebug")
	go CheckError(err, "[orm Debug Error]")
}

// 检查错误
func CheckError(err error, msg string) {
	if err != nil {
		logs.Info("["+msg+"]:", err.Error())
	}
}
