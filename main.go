package main

import (
	"github.com/kataras/iris/v12"

	"github.com/prodanlabs/kubeedge-wechat-examples/config"
	"github.com/prodanlabs/kubeedge-wechat-examples/handle"
	"github.com/prodanlabs/kubeedge-wechat-examples/utils"
)

// See https://github.com/kataras/iris/issues/1449
// Iris automatically binds the standard "context" context.Context to `iris.Context.Request().Context()`
// and any other structure that is not mapping to a registered dependency
// as a payload depends on the request, e.g XML, YAML, Query, Form, JSON.
//
// Useful to use gRPC services as Iris controllers fast and without wrappers.

func main() {
	app := config.NewAPP()

	// POST: https://localhost:443/protoreflect.FileDescriptor/method
	handle.RegisterGRPC(app)

	app.Any("/", handle.TextHandler)
	app.Any("/oauth", handle.WxOAuth)

	// The Iris server should ran under TLS (it's a gRPC requirement).
	addr := utils.GetEnv("SERVER_ADDR_PORT", "0.0.0.0:443")
	app.Run(iris.TLS(addr, "server.crt", "server.key"))
}
