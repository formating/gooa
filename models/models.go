package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

//用户表
type BUser struct {
	Uid                     int64 `orm:"pk"`
	Cid                     int64
	Ip_address              string
	Username                string
	Password                string
	Salt                    string
	Email                   string
	Forgotten_password_code string
	Forgotten_password_time int64
	Remember_code           int64
	Created_on              int64
	Last_login              int64
	Active                  int
}

//单位表
type BCompany struct {
	Cid        int64 `orm:"pk"`
	Name       string
	Phone      string
	Phone1     string
	Address    string
	Createdate int64
}

func (m *BUser) TableName() string {
	return "b_user"
}
func (m *BCompany) TableName() string {
	return "b_company"
}

func init() {
	log.Println("---------- models init called.-----------------")
	orm.RegisterModelWithPrefix("", new(BUser), new(BCompany))
}
