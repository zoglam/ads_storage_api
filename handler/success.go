package handler

import (
    "encoding/json"
    "net/http"
)

func getStatusOKPage(w http.ResponseWriter, r *http.Request, details interface{}) {
    w.WriteHeader(http.StatusOK)
    message, _ := json.Marshal(
        Message{
            OK:       true,
            Response: details,
        })
    w.Write(message)
}
