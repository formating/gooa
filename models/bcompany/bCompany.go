package bcompany

import (
	"github.com/astaxie/beego/orm"
	. "github.com/baiyuxiong/gooa/models"
)

func Insert(m *BCompany) (int64, error) {
	return orm.NewOrm().Insert(m)
}

func Read(m *BCompany, fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func Update(m *BCompany, fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func Delete(m *BCompany) error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
