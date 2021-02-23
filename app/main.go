package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/mux"

    config "github.com/zoglam/ads_storage_api/config"
    handler "github.com/zoglam/ads_storage_api/handler"
    models "github.com/zoglam/ads_storage_api/models"
)

func getRouts() *mux.Router {
    route := mux.NewRouter()
    api := route.PathPrefix("/api").Subrouter()
    adsAPI := api.PathPrefix("/ads").Subrouter()

    route.Headers("Content-Type", "application/json")
    route.NotFoundHandler = http.HandlerFunc(handler.GetEmptyPage)

    adsAPI.HandleFunc("/all", handler.GetAds).Methods("GET")
    adsAPI.HandleFunc("/get", handler.GetAd).Methods("GET")
    adsAPI.HandleFunc("/create", handler.CreateAd).Methods("POST")

    return route
}

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

    if err := models.OpenConnectionDataBase(dsn); err != nil {
        log.Fatalf("%s - %s", "Database connection failed", err.Error())
    }
    defer models.CloseConnectionDataBase()

    server := &http.Server{Addr: fmt.Sprintf(":%s", conf.ServerPort), Handler: getRouts()}

    go func() {
        fmt.Printf("Server started on http://localhost:%s\n", conf.ServerPort)
        if err := server.ListenAndServe(); err != nil {
            print("Server stopped")
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    server.Shutdown(ctx)
}
