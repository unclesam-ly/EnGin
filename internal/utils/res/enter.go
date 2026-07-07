package res

import (
	"EnGin/internal/utils/validata"

	"github.com/gin-gonic/gin"
)

// Code 自定义状态码结构
type Code struct {
	Code int
	Msg  string
}

// Success 预定义的成功状态码
var (
	Success = Code{0, "成功"}
)

// 预定义的错误状态码

var (
	UnknownError      = Code{1000, "未知错误"}
	ParamError        = Code{1001, "参数错误"}
	UnauthorizedError = Code{1002, "未认证"}
	ForbiddenError    = Code{1003, "权限不足"}
	NotFoundError     = Code{1004, "资源不存在"}
	DatabaseError     = Code{1005, "数据库错误"}
	ValidateError     = Code{1006, "验证失败"}
	ContentError      = Code{1007, "内容不合法"}
	DefaultError      = Code{7, "失败"} // 默认错误状态码
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
	TraceID string `json:"trace_id"`
}

// 成功响应

func Ok(data any, code Code, msg string, c *gin.Context) {
	message := code.Msg
	if msg != "" {
		message = msg
	}

	// 获取 TraceID
	traceID := c.GetString("trace_id")

	c.JSON(200, Response{
		Code:    code.Code,
		Data:    data,
		Msg:     message,
		TraceID: traceID,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, Success, msg, c)
}

func OkWithList(list any, count int, c *gin.Context) {
	Ok(map[string]any{
		"list":  list,
		"count": count,
	}, Success, Success.Msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Ok(data, Success, Success.Msg, c)
}

// 失败响应

func Fail(errorCode Code, msg string, c *gin.Context) {
	// 如果提供了特定的错误信息,则使用该信息,否则使用预定义的错误信息
	message := errorCode.Msg
	if msg != "" {
		message = msg
	}

	statusCode := 200
	// 根据错误代码设置合适的HTTP状态码
	switch errorCode {
	case ParamError:
		statusCode = 400 // Bad Request
	case UnauthorizedError:
		statusCode = 401 // Unauthorized
	case ForbiddenError:
		statusCode = 403 // Forbidden
	case NotFoundError:
		statusCode = 404 // Not Found
	case ValidateError, ContentError:
		statusCode = 422
	case DatabaseError, UnknownError:
		statusCode = 500 // Internal Server Error

	}

	// TraceID
	traceID := c.GetString("trace_id")

	c.JSON(statusCode, Response{
		Code:    errorCode.Code,
		Data:    gin.H{},
		Msg:     message,
		TraceID: traceID,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(DefaultError, msg, c)
}

func FailWithError(err error, c *gin.Context) {
	msg := validata.ValidateErr(err)
	Fail(ValidateError, msg, c)
}

// 添加专门的错误响应方法

func BadRequest(msg string, c *gin.Context) {
	Fail(ParamError, msg, c)
}

func Unauthorized(msg string, c *gin.Context) {
	Fail(UnauthorizedError, msg, c)
}

func Forbidden(msg string, c *gin.Context) {
	Fail(ForbiddenError, msg, c)
}

func NotFound(msg string, c *gin.Context) {
	Fail(NotFoundError, msg, c)
}

func InternalError(msg string, c *gin.Context) {
	Fail(UnknownError, msg, c)
}

func InvalidContent(msg string, c *gin.Context) {
	Fail(ContentError, msg, c)
}
