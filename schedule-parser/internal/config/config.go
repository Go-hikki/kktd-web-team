package config

import (
    "github.com/joho/godotenv"
    "os"
)

type Config struct {
    DBHost       string
    DBPort       string
    DBUser       string
    DBPassword   string
    DBName       string
    DBSSLMode    string
    ExcelFilePath string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    return &Config{
        DBHost:       os.Getenv("DB_HOST"),
        DBPort:       os.Getenv("DB_PORT"),
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBName:       os.Getenv("DB_NAME"),
        DBSSLMode:    os.Getenv("DB_SSLMODE"),
        ExcelFilePath: os.Getenv("EXCEL_FILE_PATH"),
    }, nil
}