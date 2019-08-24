package errors

import (
	"dianying/src/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeUserSuccess = 200
)

var (
	ErrorUserSuccess = errors.New(
		config.ServiceNameUser,"操作成功",errorCodeUserSuccess,
	)

	ErrorUserAlready = errors.New(
		config.ServiceNameUser,"该邮箱已被注册",errorCodeUserSuccess,
	)

	ErrorUserFailed = errors.New(
		config.ServiceNameUser,"操作异常",errorCodeUserSuccess,
	)

	ErrorUserLoginFailed = errors.New(
		config.ServiceNameUser,"密码或者用户名错误~",errorCodeUserSuccess,
	)
)