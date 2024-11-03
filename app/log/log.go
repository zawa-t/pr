package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

func PrintJSON(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		slog.Error("Faild to exec json.Marshal().", "error", err.Error())
		os.Exit(1)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	if err != nil {
		slog.Error("Faild to exec json.Indent().", "error", err.Error())
		os.Exit(1)
	}
	fmt.Println(buf.String())
}
