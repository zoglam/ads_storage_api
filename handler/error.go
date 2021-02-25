package handler

import (
    "encoding/json"
    "net/http"
)

type errorHandler struct{}

// Error ...
var Error errorHandlerInterface

func init() {
    Error = &errorHandler{}
}

// GetEmptyPage ...
func (e *errorHandler) GetEmptyPage(w http.ResponseWriter, r *http.Request) {
    httpStatus := http.StatusNotFound
    w.WriteHeader(httpStatus)
    message, _ := json.Marshal(
        Message{
            OK:               false,
            ErrorCode:        httpStatus,
            ErrorDescription: "Not Found",
        })
    w.Write(message)
}

func (e *errorHandler) getErrorPage(w http.ResponseWriter, r *http.Request, errorDetails error) {
    httpStatus := http.StatusBadRequest
    w.WriteHeader(httpStatus)
    message, _ := json.Marshal(
        Message{
            OK:               false,
            ErrorCode:        httpStatus,
            ErrorDescription: errorDetails.Error(),
        })
    w.Write(message)
}
