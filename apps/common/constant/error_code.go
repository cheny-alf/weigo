package constant

const (
	ErrorUserNotFound = 1001
)

var errorMsg = map[int]string{
	ErrorUserNotFound: "用户名/密码错误",
}

func GetErrorMsg(code int) string {
	msg, ok := errorMsg[code]
	if ok {
		return msg
	}

	return ""
}
