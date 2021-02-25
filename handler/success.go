package handler

import (
    "encoding/json"
    "net/http"
)

type successHandler struct{}

// Success ...
var success successHandlerInterface

func init() {
    success = &successHandler{}
}

func (s *successHandler) getStatusOKPage(w http.ResponseWriter, r *http.Request, details interface{}) {
    w.WriteHeader(http.StatusOK)
    message, _ := json.Marshal(
        Message{
            OK:       true,
            Response: details,
        })
    w.Write(message)
}
