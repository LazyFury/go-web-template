package response

import (
	"net/http"
	"time"
)

// Result Result
type Result struct {
	Code    ErrCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	BuildBy time.Time   `json:"build_by"`
}

// ErrCode 错误码类型
type ErrCode int

const (

	// Errors 失败
	Errors ErrCode = -1
	// NoRoute NoRoute
	NoRoute ErrCode = http.StatusNotFound
	// NoMethod NoMethod
	NoMethod ErrCode = http.StatusMethodNotAllowed
)
const (
	// Success Success
	Success ErrCode = iota + 240
	// LoginSuccess 登陆成功
	LoginSuccess

	StatusCreated ErrCode = http.StatusCreated
)
const (
	// AuthedError 认证失败
	AuthedError ErrCode = -iota - 100
	// NotFound 没有数据
	NotFound
	// RepeatEmail 邮箱已存在
	RepeatEmail
	// RepeatUserName 用户名已存在
	RepeatUserName
	// BindJSONErr 绑定json失败
	BindJSONErr

	InvalidJSONData
	InvalidQueryData
	InvalidFormData
)

type ErrorCodeTextInterface map[ErrCode]string

// ErrorCodeText 错误提示
var ErrorCodeText = ErrorCodeTextInterface{
	// base
	Success: "获取成功",
	Errors:  "遇到错误",

	// business
	LoginSuccess:  "登陆成功",
	StatusCreated: "添加成功",

	AuthedError:    "登陆超时",
	NotFound:       "没有数据",
	RepeatEmail:    "邮箱已存在",
	RepeatUserName: "用户名已存在",
	BindJSONErr:    "绑定失败,请检查参数",

	InvalidJSONData:  "JSON参数错误",
	InvalidFormData:  "Form参数错误",
	InvalidQueryData: "Url参数错误",
	// system
	NoRoute:  "路由不存在",
	NoMethod: "方法不存在",
}

// BuildBy BuildBy
var BuildBy = time.Now()

// StatusText StatusText
func StatusText(code ErrCode) string {
	var msg = ""
	for _, errCodeMap := range ErrCodeTextArray {
		_msg := errCodeMap[code]
		if _msg != "" {
			msg = _msg
		}
	}
	if msg == "" {
		msg = http.StatusText(int(code))
	}
	if msg == "" {
		msg = "未知错误码"
	}
	return msg
}

// JSON JSON
func JSON(code ErrCode, message string, data interface{}) Result {
	if message == "" {
		message = StatusText(code)
	}
	return Result{
		Code:    code,
		Message: message,
		Data:    data,
		BuildBy: BuildBy,
	}
}

// JSONSuccess 成功
func JSONSuccess(message string, data interface{}) Result {
	return JSON(Success, message, data)
}

// JSONError JSONError
func JSONError(message string, data interface{}) Result {
	return JSON(Errors, message, data)
}
