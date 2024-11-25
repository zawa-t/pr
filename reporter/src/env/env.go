package env

import (
	"fmt"
	"log/slog"
	"os"
)

type env string

var (
	Env  env = env(os.Getenv("ENV"))
	Lang     = os.Getenv("LANG")
)

func (e env) IsLocal() bool {
	return e == "local"
}

func (e env) IsTest() bool {
	return e == "test"
}

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		slog.Warn(fmt.Sprintf("The environment variable named %s is empty.", name))
	}
	return v
}
