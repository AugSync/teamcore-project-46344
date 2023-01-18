package httptransport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

	const (
		DDMMYYYY = "02-01-2006"
	)

	/* TODO: questions.Date cannot be parse to date cuz is received in an incorrect format */
	date := time.Now()

	questionsResponse := &internal.QuestionsResponse{
		Title:   "Preguntas de teamcore",
		Date:    date.Format(DDMMYYYY),
		Version: 1,
	}

	for _, v := range questions.Data {

		// string to int
		QuestionID, err := strconv.Atoi(v.QuestionID)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			fmt.Fprintf(ctx, "Client get failed: %s\n", err)
		}

		questionInfo := internal.QuestionInfo{
			QuestionID: QuestionID,
			Question:   v.Question,
		}

		questionsResponse.Data = append(questionsResponse.Data, questionInfo)
	}

	questionsResponseBody, _ := json.Marshal(questionsResponse)

	// then override already written body
	ctx.SetBody(questionsResponseBody)
}
