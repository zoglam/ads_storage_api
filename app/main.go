package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	config "github.com/zoglam/ads_storage_api/config"
	handler "github.com/zoglam/ads_storage_api/handler"
	models "github.com/zoglam/ads_storage_api/models"
)

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	models.ConnectDataBase(dsn)
	defer models.CloseConnectionDataBase()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.Headers("Content-Type", "application/json")
	s.NotFoundHandler = http.HandlerFunc(handler.GetErrorPage)

	s.HandleFunc("/ads/all", handler.GetAds).Methods("GET")
	s.HandleFunc("/ads/get", handler.GetAd).Methods("GET")
	s.HandleFunc("/ads/create", handler.CreateAd).Methods("POST")

	log.Fatal(http.ListenAndServe(config.ServerPort, r))
}
