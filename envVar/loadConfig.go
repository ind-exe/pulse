package envvar

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	envVarPath = "/etc/oob-server/config.env"
)

func LoadConfig() {
	err := godotenv.Load(envVarPath)
	if err != nil {
		log.Fatal("config: error reading config.env file")
	}
}