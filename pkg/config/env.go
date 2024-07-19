package config

import (
	"fmt"
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
		fmt.Println(err)
	}
}
