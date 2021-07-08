package envvars

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "binance-api-bot"

// LoadEnv loads env vars from .env (https://github.com/joho/godotenv/issues/43#issuecomment-503183127)
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Problem loading .env file: %e\n", err)
	}
}
