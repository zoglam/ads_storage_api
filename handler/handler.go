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

type adsHandlerInterface interface {
    GetAds(w http.ResponseWriter, r *http.Request)
    GetAd(w http.ResponseWriter, r *http.Request)
    CreateAd(w http.ResponseWriter, r *http.Request)
}

type errorHandlerInterface interface {
    GetEmptyPage(w http.ResponseWriter, r *http.Request)
    getErrorPage(w http.ResponseWriter, r *http.Request, errorDetails error)
}

type successHandlerInterface interface {
    getStatusOKPage(w http.ResponseWriter, r *http.Request, details interface{})
}

// GetRouts ...
func GetRouts() *mux.Router {
    route := mux.NewRouter()
    api := route.PathPrefix("/api").Subrouter()
    adsAPI := api.PathPrefix("/ads").Subrouter()

    route.Headers("Content-Type", "application/json")
    route.NotFoundHandler = http.HandlerFunc(Error.GetEmptyPage)

    adsAPI.HandleFunc("/all", Ads.GetAds).Methods("GET")
    adsAPI.HandleFunc("/get", Ads.GetAd).Methods("GET")
    adsAPI.HandleFunc("/create", Ads.CreateAd).Methods("POST")

    return route
}
