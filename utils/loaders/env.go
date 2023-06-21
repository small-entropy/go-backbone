package loaders

import (
	"os"

	utils_constants "github.com/small-entropy/go-backbone/constants/utils"
)

// Функция чтения из переменных окружения URI для MongoDB
func GetMongoUriFromEnv() string {
	uri := os.Getenv(utils_constants.ENV_DB_URI)
	return uri
}

// Функция чтения из переменных окружения адреса сервера
func GetServerAddressFromEnv() string {
	address := os.Getenv(utils_constants.ENV_DB_ADDRESS)
	return address
}

// Функция чтения из переменных окружения имя базы данных
func GetDatabaseFromEnv() string {
	database_name := os.Getenv(utils_constants.ENV_DB_NAME)
	return database_name
}
