package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ResponseData struct {
	// Golang way to do inheritance, the workaround: http://stackoverflow.com/questions/25051299/ddg#25051421
	Request             *http.Request
	ResponseWriter      http.ResponseWriter
	ResponseContentType string
	ContentLength       int
	RequestContentType  string
	SuccessCode         int
	FailureCode         int
}

type CUDTask func(Entry) bool
type ResponseDataCUD struct {
	*ResponseData
	SuccessComp CUDTask
}

type RTask func() []Entry
type ResponseDataR struct {
	*ResponseData
	SuccessComp RTask
}

// buildResponseCUD handles Create, Update and Delete Responses (for their high similarity).
func buildResponseCUD(r ResponseDataCUD) {
	buildResponseHeader(r.ResponseWriter.Header(), r.ResponseContentType, r.ContentLength)
	if r.Request.Header.Get("Content-Type") != r.RequestContentType {
		r.ResponseWriter.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	e, ok := parseRequestBody(r.Request.Body)
	if !ok {
		r.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	var timestamp time.Time
	if r.SuccessCode == http.StatusCreated {
		timestamp = time.Now()
	} else {
		timestamp = e.Timestamp
	}

	if r.SuccessComp(Entry{e.Data, timestamp}) {
		r.ResponseWriter.WriteHeader(http.StatusNoContent)
	} else {
		r.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func buildResponseR(r ResponseDataR) {
	// TODO: handle Read with buildResponseR
}

func buildResponseHeader(h http.Header, c string, l int) {
	h.Set("Content-Type", c)
	h.Set("Content-Length", fmt.Sprintf("%d", l))
}

func parseRequestBody(body io.ReadCloser) (Entry, bool) {
	var e Entry
	decoder := json.NewDecoder(body)
	//decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("Error: " + err.Error() + "\n")
		return Entry{}, false
	}
	return e, true
}
