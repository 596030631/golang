package main

const (
	RECODE_SUCCESS  = 0
	RECODE_DATABASE = 401
	RECODE_SESSION  = 501
	RECODE_UNKNOWN  = 405
)

var msg = map[int]string{
	RECODE_SUCCESS:  "成功",
	RECODE_DATABASE: "数据库异常",
	RECODE_SESSION:  "Session异常",
	RECODE_UNKNOWN:  "未知错误",
}

func RMsg(code int) string {
	m, s := msg[code]
	if s {
		return m
	}
	return msg[RECODE_UNKNOWN]
}
