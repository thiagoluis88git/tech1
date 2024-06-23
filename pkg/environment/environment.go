package environment

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	RedocFolderPath *string = flag.String("PATH_REDOC_FOLDER", "/docs/swagger.json", "Swagger docs folder")

	localDev = flag.String("localDev", "false", "local development")

	singleton *Environment
)

const (
	QRCodeGatewayRootURL = "QR_CODE_GATEWAY_ROOT_URL"
	QRCodeGatewayToken   = "QR_CODE_GATEWAY_TOKEN"
	DBHost               = "DB_HOST"
	DBUser               = "POSTGRES_USER"
	DBPassword           = "POSTGRES_PASSWORD"
	DBPort               = "DB_PORT"
	DBName               = "POSTGRES_DB"
)

type Environment struct {
	qrCodeGatewayRootURL string
	qrCodeGatewayToken   string
	dbHost               string
	dbPort               string
	dbName               string
	dbUser               string
	dbPassword           string
}

func LoadEnvironmentVariables() {
	flag.Parse()

	if localFlag := *localDev; localFlag != "false" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err.Error())
		}
	}

	qrCodeGatewayRootURL := getEnvironmentVariable(QRCodeGatewayRootURL)
	qrCodeGatewayToken := getEnvironmentVariable(QRCodeGatewayToken)
	dbHost := getEnvironmentVariable(DBHost)
	dbPort := getEnvironmentVariable(DBPort)
	dbUser := getEnvironmentVariable(DBUser)
	dbPassword := getEnvironmentVariable(DBPassword)
	dbName := getEnvironmentVariable(DBName)

	once := &sync.Once{}

	once.Do(func() {
		singleton = &Environment{
			qrCodeGatewayRootURL: qrCodeGatewayRootURL,
			qrCodeGatewayToken:   qrCodeGatewayToken,
			dbHost:               dbHost,
			dbPort:               dbPort,
			dbUser:               dbUser,
			dbPassword:           dbPassword,
			dbName:               dbName,
		}
	})
}

func getEnvironmentVariable(key string) string {
	value, hashKey := os.LookupEnv(key)

	if !hashKey {
		log.Fatalf("There is no %v environment variable", key)
	}

	return value
}

func GetQRCodeGatewayRootURL() string {
	if singleton != nil {
		return singleton.qrCodeGatewayRootURL
	}

	return getEnvironmentVariable(QRCodeGatewayRootURL)
}

func GetQRCodeGatewayToken() string {
	if singleton != nil {
		return singleton.qrCodeGatewayToken
	}

	return getEnvironmentVariable(QRCodeGatewayToken)
}

func GetDBHost() string {
	if singleton != nil {
		return singleton.dbHost
	}

	return getEnvironmentVariable(DBHost)
}

func GetDBPort() string {
	if singleton != nil {
		return singleton.dbPort
	}

	return getEnvironmentVariable(DBPort)
}

func GetDBName() string {
	if singleton != nil {
		return singleton.dbName
	}

	return getEnvironmentVariable(DBName)
}

func GetDBUser() string {
	if singleton != nil {
		return singleton.dbUser
	}

	return getEnvironmentVariable(DBUser)
}

func GetDBPassword() string {
	if singleton != nil {
		return singleton.dbPassword
	}

	return getEnvironmentVariable(DBPassword)
}
