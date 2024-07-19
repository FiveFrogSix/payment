package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("GO_ENV")
	if env != "" {
		env = fmt.Sprintf(".%s", env)
	}
	envFile := fmt.Sprintf(".env%s", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
}
