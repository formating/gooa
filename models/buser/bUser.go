package buser

import (
	"github.com/astaxie/beego/orm"
	. "github.com/baiyuxiong/gooa/models"
	"log"
)

func Insert(m *BUser) (int64, error) {
	return orm.NewOrm().Insert(m)
}

func Read(m *BUser, fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func Update(m *BUser, fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func Delete(m *BUser) error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func IsUserEmailExists(email string) bool {
	user, _ := FindUserByEmail(email)
	return user.Uid > 0
}

//返回的user肯定不为空，因为函数里初始化了
func FindUserByEmail(email string) (BUser, error) {
	log.Println("FindUserByEmail " + email)
	user := BUser{Email: email}
	err := orm.NewOrm().Read(&user, "email")
	return user, err
}
