package httptransport

import (
	"encoding/json"
	"fmt"

	"github.com/augsync/teamcore-project-46344/internal"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type handler struct {
	internal.Service
}

func Handler(s internal.Service) *router.Router {
	h := &handler{s}
	r := router.New()
	r.GET("/", h.index)
	r.GET("/questions", h.getQuestions)

	return r
}

func (h *handler) index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome to teamcore-project-46344!")
}

func (h *handler) getQuestions(ctx *fasthttp.RequestCtx) {

	questions, err := h.GetQuestions(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "Client get failed: %s\n", err)
		return
	}

	questionsData, _ := json.Marshal(questions)

	// then override already written body
	ctx.SetBody(questionsData)
}
