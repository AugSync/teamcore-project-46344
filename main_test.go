package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/augsync/teamcore-project-46344/internal"
	"github.com/augsync/teamcore-project-46344/internal/httptransport"
	"github.com/augsync/teamcore-project-46344/internal/service"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

// serve serves http request using provided fasthttp handler
func serve(handler fasthttp.RequestHandler, req *http.Request) (*http.Response, error) {
	godotenv.Load()

	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	return client.Do(req)
}

func TestIndexHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "http://test/", nil)
	if err != nil {
		t.Error(err)
	}

	s := service.New()

	res, err := serve(httptransport.Handler(s).Handler, r)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(fmt.Errorf("The status code we got is: %v\n", res.StatusCode))
	}

	fmt.Println(string(body))
}

func TestQuestionsHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "http://test/questions", nil)
	if err != nil {
		t.Error(err)
	}

	s := service.New()

	res, err := serve(httptransport.Handler(s).Handler, r)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(fmt.Errorf("The status code we got is: %v\n", res.StatusCode))
	}

	questions := internal.QuestionsResponse{}
	json.Unmarshal([]byte(body), &questions)

	if questions.Date == "" {
		t.Error(fmt.Errorf("Date key missing: %v\n", questions.Date))
	}

	if questions.Data == nil {
		t.Error(fmt.Errorf("Data key missing: %v\n", questions.Data))
	}

	fmt.Println(questions)
}
