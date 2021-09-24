package settings

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// load all of the env variables from .env
func LoadEnvFile() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
	}
}

// load a specified .env variable
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

// figure out how to pass the time as a env variable
func AccessTokenExpirationTime() time.Time {
	expiration_time := time.Now().Add(100 * time.Second)
	return expiration_time
}

func RefreshTokenExpirationTime() time.Time {
	expiration_time := time.Now().Add(1000 * time.Second)
	return expiration_time
}
