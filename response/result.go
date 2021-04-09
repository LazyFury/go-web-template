package response

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// LogError LogError
func LogError(err interface{}) {
	log.Printf("\n\x1b[31m[Custom Debug Result]: %v \x1b[0m\n\n", err)
}

// DuplicateEntryKey DuplicateEntryKey
var DuplicateEntryKey, _ = regexp.Compile(
	`Duplicate entry '(.*)' for key '(.*)\.(.*)'`,
)

// ParseError ParseError
func ParseError(r interface{}) Result {
	result := JSON(http.StatusInternalServerError, "", nil)
	//普通错误
	if err, ok := r.(error); ok {
		if err == gorm.ErrRecordNotFound {
			result.Code = NotFound
			result.Message = StatusText(result.Code)
		} else {
			result.Message = err.Error()
		}

		if arr := DuplicateEntryKey.FindSubmatch([]byte(err.Error())); len(arr) > 0 {
			result.Message = string(arr[3]) + " 是不可重复字段,已存在相同的数据"
		}
	}
	//错误提示
	if err, ok := r.(string); ok {
		result.Message = err
	}
	//错误码
	if err, ok := r.(ErrCode); ok {
		result.Code = err
		result.Message = StatusText(err)
	}
	if err, ok := r.(int); ok {
		result.Code = ErrCode(err)
		result.Message = StatusText(ErrCode(err))
	}
	//完整错误类型
	if data, ok := r.(Result); ok {
		result = data
	}

	return result
}

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
	// Success Success
	Success ErrCode = 1
	// Errors 失败
	Errors ErrCode = -1
	// NoRoute NoRoute
	NoRoute ErrCode = http.StatusNotFound
	// NoMethod NoMethod
	NoMethod ErrCode = http.StatusMethodNotAllowed
)
const (
	// LoginSuccess 登陆成功
	LoginSuccess ErrCode = iota + 100
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
)

// ErrorCodeText 错误提示
var ErrorCodeText = map[ErrCode]string{
	// base
	Success: "获取成功",
	Errors:  "遇到错误",

	// business
	LoginSuccess:   "登陆成功",
	AuthedError:    "登陆超时",
	NotFound:       "没有数据",
	RepeatEmail:    "邮箱已存在",
	RepeatUserName: "用户名已存在",
	BindJSONErr:    "绑定失败,请检查参数",

	// system
	NoRoute:  "路由不存在",
	NoMethod: "方法不存在",
}

// BuildBy BuildBy
var BuildBy = time.Now()

// StatusText StatusText
func StatusText(code ErrCode) string {
	msg := ErrorCodeText[code]
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
