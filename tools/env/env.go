package env

import (
	"os"

	"github.com/small-entropy/go-backbone/pkg/constants"
)

// GetMongoUriFromEnv
// Функция чтения из переменных окружения URI для MongoDB
func GetMongoUriFromEnv() string {
	uri := os.Getenv(constants.EnvDbUri)
	return uri
}

// GetServerAddressFromEnv
// Функция чтения из переменных окружения адреса сервера
func GetServerAddressFromEnv() string {
	address := os.Getenv(constants.EnvDbAddress)
	return address
}

// GetDatabaseFromEnv
// Функция чтения из переменных окружения имя базы данных
func GetDatabaseFromEnv() string {
	dbName := os.Getenv(constants.EnvDbName)
	return dbName
}
