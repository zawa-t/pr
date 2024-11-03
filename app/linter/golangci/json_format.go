package golangci

import (
	"encoding/json"
	"log/slog"
	"os"
)

type JSONFormat struct {
	Issues []Issue `json:"Issues"`
}

type Issue struct {
	FromLinter  string   `json:"FromLinter"`
	Text        string   `json:"Text"`
	Severity    string   `json:"Severity"`
	SourceLines []string `json:"SourceLines"`
	// Replacement ? // NOTE: 型や使用方法が不明のため一旦コメントアウト
	Pos          Pos  `json:"Pos"`
	ExpectNoLint bool `json:"ExpectNoLint"`
}

type Pos struct {
	Filename             string `json:"Filename"`
	Offset               uint   `json:"Offset"`
	Line                 uint   `json:"Line"`
	Column               uint   `json:"Column"`
	ExpectedNoLintLinter string `json:"ExpectedNoLintLinter"`
}

func decodeJSONData(stdin *os.File) JSONFormat {
	var jsonData JSONFormat
	decoder := json.NewDecoder(stdin)
	if err := decoder.Decode(&jsonData); err != nil {
		slog.Error("Failed to JSON Decode.", "error", err.Error())
		os.Exit(1)
	}

	return jsonData
}
