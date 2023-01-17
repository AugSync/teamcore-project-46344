package internal

import (
	"github.com/valyala/fasthttp"
)

type QuestionsData struct {
	Date string     `json:"date"`
	Data []Question `json:"data"`
}

type Question struct {
	QuestionID string   `json:"question_id"`
	Question   string   `json:"question"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	AnswerID string `json:"answer_id"`
	Answer   string `json:"answer"`
}

type Service interface {
	GetQuestions(*fasthttp.RequestCtx) (QuestionsData, error)
}
