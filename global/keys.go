package g

import "errors"

// 定义全局用到的Key

const (
	CTX_DB        = "_db_field"
	CTX_USER_AUTH = "_user_auth"
)

var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)
