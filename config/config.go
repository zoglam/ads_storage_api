package config

import "os"

var (
	ServerPort, _ = os.LookupEnv("SERVER_PORT")

	DBUser, _ = os.LookupEnv("MARIA_USER")
	DBPass, _ = os.LookupEnv("MARIA_PASSWORD")
	DBHost, _ = os.LookupEnv("MARIA_HOST")
	DBPort, _ = os.LookupEnv("MARIA_PORT")
	DBName, _ = os.LookupEnv("MARIA_DB")
)
