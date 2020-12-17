module github.com/Treblex/go-web-template

go 1.15

require (
	github.com/Treblex/go-echo-demo/server v0.0.0-20201216024905-8679e4d0afc5
	github.com/aliyun/aliyun-oss-go-sdk v2.1.5+incompatible
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace github.com/Treblex/go-echo-demo/server => ../go-echo-demo/server
