package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBConnection string
	DBHost       string
	DBPort       string
	DBName       string
	DBUsername   string
	DBPassword   string
	DBSSLMode    string
)

func SetEnvFileToVariable() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBConnection = os.Getenv("DB_CONNECTION")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_DATABASE")
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBSSLMode = os.Getenv("DB_SSLMODE")
	fmt.Println("DBConnection               : ", DBConnection)
	fmt.Println("DBHost                     : ", DBHost)
	fmt.Println("DBPort                     : ", DBPort)
	fmt.Println("DBName                     : ", DBName)
	fmt.Println("DBUsername                 : ", DBUsername)
	fmt.Println("DBPassword                 : ", DBPassword)
	fmt.Println("DBSSLMode                  : ", DBSSLMode)

}
