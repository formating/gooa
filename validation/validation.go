package validation

import (
	"github.com/astaxie/beego/validation"
)

//表单验证
func ValidUser(m *BUser) string {
	v := validation.Validation{}
	v.Required(m.Username, "username").Message("邮箱不能为空")
	v.Required(m.Password, "password").Message("密码不能为空")
	v.Email(m.Username, "username").Message("邮箱格式不正确")
	if v.HasErrors() {
		// 如果有错误信息，证明验证没通过
		return utils.WrapErrors(v.Errors)
	}
	return ""
}

//表单验证
func ValidCompany(m *BCompany) string {
	v := validation.Validation{}
	v.Required(m.Name, "name").Message("单位名称不能为空")
	v.Required(m.Phone, "phone").Message("电话不能为空")
	v.Required(m.Address, "address").Message("地址不能为空")
	if v.HasErrors() {
		// 如果有错误信息，证明验证没通过
		return utils.WrapErrors(v.Errors)
	}
	return ""
}
