package envvar

import (
	"errors"
	"os"
)

func GetVar(key string) (string, error){
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("config: error reading env variable")
	}
	return value, nil
}