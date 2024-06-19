package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
	"gvb_server/utils/valid"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json() string {
	byteData, _ := json.Marshal(r)
	return string(byteData)
}

type ListResponse struct {
	Count int64 `json:"count"`
	List  any   `json:"list"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDataSSE(data any, c *gin.Context) {
	content := Response{
		Code: SUCCESS,
		Data: data,
		Msg:  "成功",
	}.Json()
	c.SSEvent("", content)
}

func OkWithSSE(data any, msg string, c *gin.Context) {
	content := Response{
		Code: SUCCESS,
		Data: data,
		Msg:  msg,
	}.Json()
	c.SSEvent("", content)
}

func OkWithList(list any, count int64, c *gin.Context) {
	lr := ListResponse{
		Count: count,
		List:  list,
	}
	OkWithData(lr, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithMessageSSE(msg string, c *gin.Context) {
	data := Response{
		Code: ERROR,
		Data: map[string]any{},
		Msg:  msg,
	}.Json()
	c.SSEvent("", data)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

func FailWithValidError(err error, c *gin.Context) {
	msg := valid.Error(err)
	FailWithMessage(msg, c)
}

func FailWithCode(code gin.ErrorType, c *gin.Context) {
	if msg, ok := GinErrorMap[code]; ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
