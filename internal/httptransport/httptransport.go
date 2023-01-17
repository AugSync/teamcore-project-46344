package httptransport

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type handler struct {
}

func Handler() *router.Router {
	h := &handler{}
	r := router.New()
	r.GET("/", h.Index)
	r.GET("/hello/{name}", h.Hello)

	return r
}

func (h *handler) Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome to teamcore-project-46344!")
}

func (h *handler) Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}
