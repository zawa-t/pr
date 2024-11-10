package env

import (
	"fmt"
	"log/slog"
	"os"
)

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		slog.Warn(fmt.Sprintf("The environment variable named %s is empty.", name))
	}
	return v
}

var local env = "local"

type env string

func (e env) IsLocal() bool {
	return e == local
}

var (
	Env env = env(os.Getenv("ENV"))
)
