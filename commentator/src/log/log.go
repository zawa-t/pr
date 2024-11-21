package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
)

func PrintJSON(msg string, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		slog.Error("Faild to exec json.Marshal().", "error", err.Error())
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	if err != nil {
		slog.Error("Faild to exec json.Indent().", "error", err.Error())
	}
	slog.Debug(msg)
	fmt.Println(buf.String())
}
