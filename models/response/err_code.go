package response

import "github.com/gin-gonic/gin"

// ErrCode 自定义ErrCode类型，不用gin自带的ErrorType类型
type ErrCode int

const (
	SettingError  ErrCode = 1001 // 系统错误
	ArgumentError ErrCode = 1002 // 参数错误
)

//var (
//	CustomErrorMap = map[ErrCode]string{
//		SettingError:  "系统错误",
//		ArgumentError: "参数错误",
//	}
//)

// GinErrorMap gin自带ErrorType类型
var (
	GinErrorMap = map[gin.ErrorType]string{
		gin.ErrorTypeBind:    "参数绑定错误",
		gin.ErrorTypeRender:  "渲染错误",
		gin.ErrorTypePrivate: "私有错误",
		gin.ErrorTypePublic:  "公共错误",
	}
)
