package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"

    config "github.com/zoglam/ads_storage_api/config"
    handler "github.com/zoglam/ads_storage_api/handler/http"
    models "github.com/zoglam/ads_storage_api/models"
)

func main() {
    conf := config.New()
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s",
        conf.Maria.DBUser,
        conf.Maria.DBPass,
        conf.Maria.DBHost,
        conf.Maria.DBPort,
        conf.Maria.DBName,
    )
    err := models.OpenConnectionDataBase(dsn)
    if err != nil {
        log.Fatalf("%s - %s\n%s", "Database connection failed", err.Error(), dsn)
    }
    defer models.CloseConnectionDataBase()

    r := mux.NewRouter()
    s := r.PathPrefix("/api").Subrouter()
    r.Headers("Content-Type", "application/json")
    s.NotFoundHandler = http.HandlerFunc(handler.GetEmptyPage)

    s.HandleFunc("/ads/all", handler.GetAds).Methods("GET")
    s.HandleFunc("/ads/get", handler.GetAd).Methods("GET")
    s.HandleFunc("/ads/create", handler.CreateAd).Methods("POST")

    fmt.Printf("Server started on http://localhost:%s\n", conf.ServerPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), r))
}
