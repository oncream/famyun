package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type HandleFunc func(c *gin.Context) error

type Handler interface {
	RegisterRouters(router gin.IRouter)
}

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

type Params struct {
	fx.In
	Gin      *gin.Engine
	Handlers []Handler `group:"handlers"`
}

func RegisterHandlers(params Params) {
	fmt.Println("注册路由")
	api := params.Gin.Group("/api")
	for _, handler := range params.Handlers {
		handler.RegisterRouters(api)
	}
}
