package config

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func NewAPP() *iris.Application {
	// 创建app结构体对象
	app := iris.New()
	// 配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	// 配置日志
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		// 为空则从 `ctx.Values().Get(MessageContextKey)` 获取内容
		MessageContextKeys: []string{"logger_message", ""},
		// 为空则从 `ctx.Values().Get(MessageHeaderKey)` 获取头信息
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)
	app.Logger().SetLevel("debug")
	return app
}
