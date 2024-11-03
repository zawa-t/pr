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
