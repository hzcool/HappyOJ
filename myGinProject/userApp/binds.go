package userApp

import (
	"strings"
)

//登陆绑定
type LoginBind struct {
	Username string `form:"username"  validate:"lte=20,required"`
	Password string `form:"password"  validate:"gte=6,lte=16,required,printascii"`
}

//字符的合法性
func (lg *LoginBind) CharsValidate() (bool,string){
	if strings.ContainsAny(lg.Username," \n\t\r") {
		return false, "Username 不能包含空字符"
	}
	if strings.ContainsAny(lg.Password," \n\t\r") {
		return false, "Password 不能包含空字符"
	}
	return true, ""
}

//注册绑定
type RegisterBind struct {
	LoginBind
	Email    string	`form:"email"  validate:"email,required"`
	School 	 string `form:"school" validate:"lte=20,required"`
	Description string `form:"description" validate:"lte=255"`
}

// 修改绑定,update需要的与Register一样
type UpdateBind = RegisterBind


// 修改密码绑定
type UpdatePasswordBind struct {
	OldPassword string `form:"oldPassword"  validate:"gte=6,lte=16,required,printascii"`
	NewPassword string `form:"newPassword"  validate:"gte=6,lte=16,required,printascii"`
}



