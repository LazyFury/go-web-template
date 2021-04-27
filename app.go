package gwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

type App struct {
	*gin.Engine

	NoMethodHandleFunc []gin.HandlerFunc
	NoRouteHandleFunc  []gin.HandlerFunc

	PreHandleFunc    []gin.HandlerFunc
	MiddleHandleFunc []gin.HandlerFunc
	LastHandleFunc   []gin.HandlerFunc

	Router []func(c *gin.RouterGroup)
}

func (a *App) Run(addr string) (err error) {
	a.Engine.Use(a.PreHandleFunc...)

	a.Engine.Use(a.MiddleHandleFunc...)
	// HandleMethodNotAllowed
	a.Engine.NoMethod(a.NoMethodHandleFunc...)
	a.Engine.NoRoute(a.NoRouteHandleFunc...)

	for _, route := range a.Router {
		route(&a.RouterGroup)
	}

	a.Engine.Use(a.LastHandleFunc...)

	return a.Engine.Run(addr)
}

// InitRouter 保证全局中间价注册完成之后才会初始化路由
// 路由也是中间件的形式实现，use注册的顺序保证执行的顺序
func (a *App) InitRouter(fn ...func(c *gin.RouterGroup)) {
	a.Router = append(a.Router, fn...)
}
func (a *App) PreUse(midd ...gin.HandlerFunc) {
	a.PreHandleFunc = append(a.PreHandleFunc, midd...)
}
func (a *App) Use(midd ...gin.HandlerFunc) {
	a.MiddleHandleFunc = append(a.MiddleHandleFunc, midd...)
}

func (a *App) LastUse(midd ...gin.HandlerFunc) {
	a.LastHandleFunc = append(a.LastHandleFunc, midd...)
}
func (a *App) NoRouteUse(midd ...gin.HandlerFunc) {
	a.NoRouteHandleFunc = append(a.NoRouteHandleFunc, midd...)
}
func (a *App) NoMethodUse(midd ...gin.HandlerFunc) {
	a.NoMethodHandleFunc = append(a.NoMethodHandleFunc, midd...)
}

// New 初始化
func New() *App {
	app := &App{
		Engine: gin.New(),
		PreHandleFunc: []gin.HandlerFunc{
			response.Recover,
			gin.Logger(),
		},
		LastHandleFunc: []gin.HandlerFunc{},
		NoMethodHandleFunc: []gin.HandlerFunc{
			NoMethod,
		},
		NoRouteHandleFunc: []gin.HandlerFunc{
			NoRoute,
		},
	}
	app.HandleMethodNotAllowed = true
	// 移除多余斜杠 /api//v1/doSomething/ => /api/v1/doSomething
	app.RemoveExtraSlash = true
	// 重定向请求移除斜杠请求
	app.RedirectTrailingSlash = true

	// 错误码配置
	response.RecoverErrHtml = true
	return app
}
func NoRoute(c *gin.Context) {
	if c.Request.URL.Path != "/favicon.ico" {
		response.Error(response.NoRoute)
	}
}
func NoMethod(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	response.Error(response.NoMethod)
}
