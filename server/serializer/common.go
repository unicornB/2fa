package serializer

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	//签名错误
	CodeSignError = 40002
	//时间戳错误
	CodeTimestampError = 40003
	//随机数错误
	CodeNonceError = 40004
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = fmt.Sprintf("%+v", err)
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

func Success(msg string, data interface{}) Response {
	return Response{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}
func SuccessNoData() Response {
	return Response{
		Code: 200,
		Msg:  "success",
	}
}
func Error(msg string) Response {
	return Response{
		Code: 0,
		Msg:  msg,
	}
}
func SignError(msg string) Response {
	return Response{
		Code: CodeSignError,
		Msg:  msg,
	}
}
func TimestampError(msg string) Response {
	return Response{
		Code: CodeTimestampError,
		Msg:  msg,
	}
}
func NonceError(msg string) Response {
	return Response{
		Code: CodeNonceError,
		Msg:  msg,
	}
}
