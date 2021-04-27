package gwt

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	CorsConfig struct {
		AllowOrigins     []string
		AllowMethods     []string
		AllowHeaders     []string
		AllowCredentials bool
	}
)

var (
	DefaultAllowHeaders = []string{"authorization", "token", "content-type", "x-requested-with"}
	DefaultAllowMethods = []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch}
)

func DefaultCorsConfig() *CorsConfig {
	return &CorsConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     DefaultAllowHeaders,
		AllowMethods:     DefaultAllowMethods,
	}
}

func DefaultCors(c *gin.Context) {
	Cors(c, DefaultCorsConfig())
}

//Cors 自己尝试的cors配置实现
func Cors(c *gin.Context, config *CorsConfig) {
	req := c.Request
	origin := c.Request.Host
	if len(origin) == 0 {
		// request is not a CORS request
		return
	}

	// host := c.Request.Host

	// if origin == "http://"+host || origin == "https://"+host {
	// 	// request is not a CORS request but have origin header.
	// 	// for example, use fetch api
	// 	return
	// }

	allowOrigins := []string{"*"}
	if len(config.AllowOrigins) > 0 {
		allowOrigins = config.AllowOrigins
	}

	// 如果请求的域名在放行名单中
	inAllow := 0
	for i, o := range allowOrigins {
		if origin == o || o == "*" {
			inAllow = i + 1
		}
	}

	// 如果不在名单中  且 没有开放任意域名通过
	if inAllow == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// set header
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ","))
	c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(config.AllowCredentials))
	c.Header("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ","))

	if req.Method == http.MethodOptions {
		c.Status(http.StatusNoContent)
	}

	c.Next()
}
