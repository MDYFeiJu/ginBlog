package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code =1000... 用户模块的错误
	ErrorUsernameUsed  = 1001
	ErrorPasswordWrong = 1002
	ErrorUserNotExist  = 1003
	ErrorTokenExist    = 1004
	ErrorTokenRuntime  = 1005
	ErrorTokenWrong    = 1006
	TokenTypeWrong     = 1007
	UserNoRight        = 1008
	// code =2000... 文章模块的错误
	ErrorExistArticle = 2001
	// code =3000... 分类模块的错误
	ErrorCategoryNameUsed = 3001

	// 测试用
	ErrorUnmatchData = 4001
)

var codemsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ErrorUsernameUsed:     "用户名已存在",
	ErrorPasswordWrong:    "密码错误",
	ErrorUserNotExist:     "用户不存在",
	ErrorTokenExist:       "TOKEN不存在",
	ErrorTokenRuntime:     "TOKEN已过期",
	ErrorTokenWrong:       "TOKEN错误",
	TokenTypeWrong:        "TOKEN格式错误",
	ErrorCategoryNameUsed: "分类已存在",
	ErrorExistArticle:     "文章已存在",
	UserNoRight:           "用户无权限",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
