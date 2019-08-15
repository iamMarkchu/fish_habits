package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	STATUS_BLANK = iota
	STATUS_NORMAL
	STATUS_BANNED

	TEST_USER_ID = 1

	DEFAULT_DISPLAY_ORDER = 99
)

// 类别习惯表
//type CategoryHabit struct {
//	Id      int       `orm:"auto"`
//	CateId  int       `orm:"description(类别id)"`
//	HabitId int       `orm:"description(习惯id)"`
//	Status  uint8     `orm:"default(1);description(类别下习惯的状态)"`
//	Created time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
//	Updated time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
//}

// 用户习惯表
type UserHabit struct {
	Id      int       `orm:"auto"`
	UserId  int       `orm:"description(用户ID)"`
	HabitId int       `orm:"description(习惯ID)"`
	Status  uint8     `orm:"default(1);description(状态)"`
	Created time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func (m *UserHabit) Store() (int64, error) {
	var (
		o        = orm.NewOrm()
		err      error
		insertId int64
	)
	insertId, err = o.Insert(m)
	return insertId, err
}

func (m *UserHabit) Fetch() error {
	var (
		o        = orm.NewOrm()
		err      error
	)
	err = o.Read(m, "HabitId", "UserId")
	return err
}

func (m *UserHabit) Remove() (int64, error) {
	var (
		o        = orm.NewOrm()
	)
	m.Status = STATUS_BANNED
	return o.Update(m, "Status")
}

// 打卡表
type Sign struct {
	Id          int       `orm:"auto"`
	UserHabitId int       `orm:"column(uh_id);description(用户习惯关系表ID)"`
	SignDay     int    `orm:"description(签到日期)"`
	Created     time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated     time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func (m *Sign) Fetch() (error) {
	var (
		o        = orm.NewOrm()
		err error
	)
	err = o.Read(m, "UserHabitId", "SignDay")
	return err
}

func (m *Sign) Store() (int64, error) {
	var (
		o        = orm.NewOrm()
		err error
		id int64
	)
	id, err = o.Insert(m)
	return id, err
}

// 用户表
type User struct {
	Id       int       `orm:"auto"`
	Name     string    `orm:"unique;description(用户名)"`
	NickName string    `orm:"default();description(昵称)"`
	Password string    `orm:"default();description(密码)"`
	Status   uint8     `orm:"default(1);description(用户状态)"`
	Created  time.Time `orm:"auto_now_add;column(created_at);type(datetime)"`
	Updated  time.Time `orm:"auto_now;column(updated_at);type(datetime)"`
}

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	go CheckError(err, "[RegisterDriver Error]")
	err = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	go CheckError(err, "[RegisterDataBase Error]")
	orm.RegisterModel(new(User), new(Category), new(Habit), new(UserHabit), new(Sign))
	err = orm.RunSyncdb("default", false, true)
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
