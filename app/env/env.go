package env

import (
	"log/slog"
	"os"
)

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		slog.Warn("failed to get environment variable.", "name", name)
	}
	return v
}

var local env = "local"

type env string

func (e env) IsLocal() bool {
	return e == local
}

var (
	Env env = env(getEnv("ENV"))
)
