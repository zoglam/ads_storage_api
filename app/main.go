package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"

    config "github.com/zoglam/ads_storage_api/config"
    handler "github.com/zoglam/ads_storage_api/handler"
    models "github.com/zoglam/ads_storage_api/models"
)

func main() {
    if err := models.OpenConnectionDataBase(); err != nil {
        log.Fatalf("%s - %s", "Database connection failed", err.Error())
    }
    defer models.CloseConnectionDataBase()

    server := &http.Server{
        Addr:    fmt.Sprintf(":%s", config.Params.ServerPort),
        Handler: handler.GetRouts(),
    }

    go func() {
        fmt.Printf("Server started on http://localhost:%s\n", config.Params.ServerPort)
        if err := server.ListenAndServe(); err != nil {
            print("Server stopped")
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop
    server.Shutdown(context.Background())
}
