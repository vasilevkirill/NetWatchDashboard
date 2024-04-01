package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        uint      `orm:"auto"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	LastEnter time.Time `orm:"auto_now;type(datetime)"`
	Login     string    `orm:"size(100);unique"`
	Password  string    `orm:"size(100)"`
	Role      uint      `orm:"size(2)"`
	Enabled   bool      `orm:"default(true);description(this is status)"`
}

func GetAllUsers() ([]*User, error) {
	var users []*User
	o := orm.NewOrmUsingDB("config")
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UserAuth(login, password string) (User, error) {
	var u User
	o := orm.NewOrmUsingDB("config")
	err := o.QueryTable("user").Filter("Login", login).Filter("Enabled", true).One(&u)
	if err != nil {
		return User{}, err
	}
	if u.Id == 0 {
		return User{}, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return User{}, nil
	}
	return u, nil
}
