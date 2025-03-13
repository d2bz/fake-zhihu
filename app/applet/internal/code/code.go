package code

import "zhihu/pkg/xcode"

var (
	MobileEmpty       = xcode.Define(10001, "手机号不能为空")
	MobileHasRegister = xcode.Define(10002, "手机号已注册")
	PasswordEmpty     = xcode.Define(10003, "密码不能为空")
)
