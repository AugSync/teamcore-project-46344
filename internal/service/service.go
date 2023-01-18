package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/augsync/teamcore-project-46344/internal"
	"github.com/valyala/fasthttp"
)

type service struct {
}

func New() internal.Service {
	return &service{}
}

func (s *service) GetQuestions(ctx *fasthttp.RequestCtx) (internal.QuestionsData, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	questions := internal.QuestionsData{}

	req.SetRequestURI("https://us-central1-teamcore-retail.cloudfunctions.net/test_mobile/api/questions")
	// fasthttp does not automatically request a gzipped response.
	// We must explicitly ask for it.

	token := os.Getenv("API_TOKEN")
	if token == "" {
		return questions, fmt.Errorf("Missing api token")
	}

	var bearer = "Bearer " + token

	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Authorization", bearer)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	err := fasthttp.Do(req, resp)
	if err != nil {
		return questions, fmt.Errorf("Client get failed: %s\n", err)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return questions, fmt.Errorf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
	}

	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		return questions, fmt.Errorf("Expected content type application/json but got %s\n", contentType)
	}

	// Do we need to decompress the response?
	contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte
	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}

	json.Unmarshal([]byte(body), &questions)
	fmt.Printf("Request successfully: %s \n", questions)

	return questions, nil
}
