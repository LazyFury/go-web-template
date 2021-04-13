package response

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IsReqFromHTML ReqFromHtml
func IsReqFromHTML(c *gin.Context) bool {
	req := c.Request
	reqAccept := strings.Split(req.Header.Get("Accept"), ",")[0]
	return reqAccept == "text/html"
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
