package config

import "strconv"

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + "-" + e.Message
}

var ErrArr = []*Error{
	{
		Code:    1000001,
		Message: "操作失败",
	},
	{
		Code:    100000001,
		Message: "超时未响应",
	},
	{
		Code:    100000014,
		Message: "未录入授权TOKEN",
	},
	{
		Code:    100000017,
		Message: "签名不正确",
	},
}

func buildErr(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
