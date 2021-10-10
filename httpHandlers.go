package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeData(w http.ResponseWriter, _ *http.Request) {
	var data = readJSON()
	if len(data) > 0 {
		d, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		buildResponseHeader(w.Header(), "application/json", len(d))

		_, err = w.Write(d)
		if err != nil {
			errorCode := http.StatusServiceUnavailable
			log.Printf("%d %s; %s\n", errorCode, http.StatusText(errorCode), err.Error())
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

var respDataCUD = ResponseData{
	ResponseContentType: "text/plain",
	ContentLength:       0,
	RequestContentType:  "application/json",
	FailureCode:         http.StatusInternalServerError,
}

func (s *server) ReceiveData(w http.ResponseWriter, r *http.Request) {
	respData := respDataCUD
	respData.Request = r
	respData.ResponseWriter = w
	respData.SuccessCode = http.StatusCreated
	buildResponseCUD(ResponseDataCUD{
		ResponseData: &respData,
		SuccessComp:  writeJSON,
	})
}

func (s *server) UpdateData(w http.ResponseWriter, r *http.Request) {
	respData := respDataCUD
	respData.Request = r
	respData.ResponseWriter = w
	respData.SuccessCode = http.StatusNoContent
	buildResponseCUD(ResponseDataCUD{
		ResponseData: &respData,
		SuccessComp:  updateJSON,
	})
}

func (s *server) DeleteData(w http.ResponseWriter, r *http.Request) {
	respData := respDataCUD
	respData.Request = r
	respData.ResponseWriter = w
	respData.SuccessCode = http.StatusNoContent
	buildResponseCUD(ResponseDataCUD{
		ResponseData: &respData,
		SuccessComp:  deleteJSON,
	})
}
