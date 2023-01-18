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

type QuestionsResponse struct {
	Title   string         `json:"titulo"`
	Date    string         `json:"dia"`
	Data    []QuestionInfo `json:"info"`
	Version int            `json:"api_version"`
}

type QuestionInfo struct {
	QuestionID int    `json:"pregunta_id"`
	Question   string `json:"pregunta"`
}

type Service interface {
	GetQuestions(*fasthttp.RequestCtx) (QuestionsData, error)
}
