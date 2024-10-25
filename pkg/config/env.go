package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("GO_ENV")
	fmt.Println(env)
	if env != "" {
		env = fmt.Sprintf(".%s", env)
	}
	envFile := fmt.Sprintf(".env%s", env)
	fmt.Println(envFile)
	if err := godotenv.Load(envFile); err != nil {
		fmt.Println(err)
	}
	// if err := godotenv.Load(".env"); err != nil {
	// 	fmt.Println(err)
	// }

}
