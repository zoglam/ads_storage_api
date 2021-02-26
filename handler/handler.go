package handler

import (
    "net/http"

    "github.com/gorilla/mux"
)

// Message ...
type Message struct {
    OK               bool        `json:"ok"`
    ErrorCode        int         `json:"error_code,omitempty"`
    ErrorDescription interface{} `json:"description,omitempty"`
    Response         interface{} `json:"response,omitempty"`
}

// GetRouts ...
func GetRouts() *mux.Router {
    route := mux.NewRouter()
    api := route.PathPrefix("/api").Subrouter()
    adsAPI := api.PathPrefix("/ads").Subrouter()

    route.Headers("Content-Type", "application/json")
    route.NotFoundHandler = http.HandlerFunc(getEmptyPage)

    adsAPI.HandleFunc("/all", getAds).Methods("GET")
    adsAPI.HandleFunc("/get", getAd).Methods("GET")
    adsAPI.HandleFunc("/create", createAd).Methods("POST")

    return route
}
