package constants

import (
	"os"
)

func LoadEnv(key string) string {
	return os.Getenv(key)
}
