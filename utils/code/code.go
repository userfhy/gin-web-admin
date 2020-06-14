package code

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
	TokenInvalid  = 401
	UnknownError  = 900

	ErrorAuthCheckTokenFail     = 20001
	ErrorAuthCheckTokenTimeout  = 20002
	ErrorAuthToken              = 20003
	ErrorAuth                   = 20004
	ErrorUserPasswordInvalid    = 20005
	AuthTokenInBlockList        = 20006
	ErrorUserOldPasswordInvalid = 20007
)

var MsgFlags = map[int]string{
	SUCCESS:                     "ok",
	ERROR:                       "fail",
	UnknownError:                "未知错误",
	InvalidParams:               "请求参数错误",
	TokenInvalid:                "Token参数无效或不存在",
	ErrorAuthCheckTokenFail:     "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:  "Token已超时",
	ErrorAuthToken:              "Token生成失败",
	ErrorAuth:                   "Token错误",
	ErrorUserPasswordInvalid:    "对应用户名或密码错误",
	AuthTokenInBlockList:        "该Token在blockList中已存在",
	ErrorUserOldPasswordInvalid: "对应用户原密码错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
