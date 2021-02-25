package config

import "os"

// MariaConfig ...
type MariaConfig struct {
    DBUser string
    DBPass string
    DBHost string
    DBPort string
    DBName string
}

// Config ...
type Config struct {
    Maria      MariaConfig
    ServerPort string
}

// Params ...
var Params *Config

func init() {
    Params = new()
}

// New returns a new Config struct
func new() *Config {
    return &Config{
        Maria: MariaConfig{
            DBUser: getEnv("MARIA_USER", "root"),
            DBPass: getEnv("MARIA_PASSWORD", "secret"),
            DBHost: getEnv("MARIA_HOST", "127.0.0.1"),
            DBPort: getEnv("MARIA_PORT", "3306"),
            DBName: getEnv("MARIA_DB", "database"),
        },
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }

    return defaultVal
}
