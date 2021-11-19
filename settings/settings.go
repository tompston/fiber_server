package settings

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// load all of the variables from .env
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

func AccessTokenExpirationTime() time.Time {
	env_time, _ := strconv.Atoi(Config("JWT_ACCESS_TOKEN_DURATION"))
	duration := time.Duration(env_time)
	expiration_time := time.Now().Add(duration * time.Minute)
	return expiration_time
}

func RefreshTokenExpirationTime() time.Time {
	env_time, _ := strconv.Atoi(Config("JWT_REFRESH_TOKEN_DURATION"))
	duration := time.Duration(env_time)
	expiration_time := time.Now().Add(duration * time.Minute)
	return expiration_time
}
