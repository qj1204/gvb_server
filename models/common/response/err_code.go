package response

type ErrCode int

const (
	SettingError  ErrCode = 1001 // 系统错误
	ArgumentError ErrCode = 1002 // 参数错误
)

var (
	ErrorMap = map[ErrCode]string{
		SettingError:  "系统错误",
		ArgumentError: "参数错误",
	}
)
