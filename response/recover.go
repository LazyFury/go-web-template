package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var RecoverErrHtml = false

type RecoverRenderType func(c *gin.Context, code int, result *Result)

var RecoverRender RecoverRenderType = func(c *gin.Context, code int, result *Result) {
	c.HTML(code, "err.html", result)
}

// Recover 使用defer调用阻止panic中止程序
func Recover(c *gin.Context) {
	defer func(c *gin.Context) {
		if r := recover(); r != nil {
			result := ParseError(r)

			var code = c.Writer.Status()
			// result.Code = ErrCode(code)

			//返回内容
			if IsReqFromHTML(c) && RecoverErrHtml {
				RecoverRender(c, code, &result)
			} else {
				c.JSON(code, result)
			}
			// c.JSON(code, result)
			LogError(fmt.Sprintf("URL:%s ;\nErr: %v", c.Request.URL.RequestURI(), result), "Error")

			// "打断response继续写入内容"
			c.Abort()
		}
	}(c)

	c.Next()
}
