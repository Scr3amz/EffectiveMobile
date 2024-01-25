package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerPort string
	HostDB     string
	DriverDB   string
	UserDB     string
	PasswordDB string
	NameDB     string
	PortDB     string
}

// LoadMainConfig search for config file on confPath
// and returns Config structure and error.
func LoadMainConfig(confName, confType string) (Config, error) {

	err := godotenv.Load(fmt.Sprintf("%s.%s", confName, confType))
	if err != nil {
		return Config{}, fmt.Errorf("failed to load .env file: %v", err)
	}
	config := Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		HostDB:     os.Getenv("DB_HOST"),
		DriverDB:   os.Getenv("DB_DRIVER"),
		UserDB:     os.Getenv("DB_USER"),
		PasswordDB: os.Getenv("DB_PASSWORD"),
		NameDB:     os.Getenv("DB_NAME"),
		PortDB:     os.Getenv("DB_PORT"),
	}

	return config, nil
}
