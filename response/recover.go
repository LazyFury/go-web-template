package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var RecoverErrHtml = false
var RecoverErrTemplateName = "err/error.html"

// Recover 使用defer调用阻止panic中止程序
func Recover(c *gin.Context) {
	defer func(c *gin.Context) {
		if r := recover(); r != nil {
			result := ParseError(r)

			var code = c.Writer.Status()
			result.Code = ErrCode(code)

			//返回内容
			if IsReqFromHTML(c) && RecoverErrHtml {
				c.HTML(code, RecoverErrTemplateName, result)
			} else {
				c.JSON(code, result)
			}
			// c.JSON(code, result)
			LogError(fmt.Sprintf("URL:%s ;\nErr: %v", c.Request.URL.RequestURI(), result))

			// "打断response继续写入内容"
			c.Abort()
		}
	}(c)

	c.Next()
}
