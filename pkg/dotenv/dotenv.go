package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("ENV") != "prod" {
		godotenv.Load()
	}
}
