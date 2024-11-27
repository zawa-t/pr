package golangcilint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"

	"github.com/zawa-t/pr/src/report"
)

type JSON struct {
	Issues []Issue `json:"Issues"`
}

type Issue struct {
	FromLinter  string   `json:"FromLinter"`
	Text        string   `json:"Text"`
	Severity    string   `json:"Severity"`
	SourceLines []string `json:"SourceLines"`
	// Replacement ? // NOTE: 型や使用方法が不明のため一旦コメントアウト
	Pos                  Pos    `json:"Pos"`
	ExpectNoLint         bool   `json:"ExpectNoLint"`
	ExpectedNoLintLinter string `json:"ExpectedNoLintLinter"`
}

type Pos struct {
	Filename string `json:"Filename"`
	Offset   uint   `json:"Offset"`
	Line     uint   `json:"Line"`
	Column   uint   `json:"Column"`
}

func Decode(stdin io.Reader) (*JSON, error) {
	var jsonData JSON
	if err := json.NewDecoder(stdin).Decode(&jsonData); err != nil {
		return nil, fmt.Errorf("failed to Decode() :%w", err)
	}
	return &jsonData, nil
}

func MakeContents(alternativeText, customMessageFormat *string, issues []Issue) ([]report.Content, error) {
	contents := make([]report.Content, 0)
	for _, v := range issues {
		var message report.Message
		if customMessageFormat != nil {
			tmpl, err := template.New("customMessageFormat").Parse(*customMessageFormat)
			if err != nil {
				return nil, fmt.Errorf("failed to Parse(): %w", err)
			}

			var result bytes.Buffer
			err = tmpl.Execute(&result, v)
			if err != nil {
				return nil, fmt.Errorf("failed to Execute(): %w", err)
			}
			message = report.CustomMessage(result.String())
		} else {
			var text string
			if alternativeText != nil {
				text = *alternativeText
			} else {
				text = v.Text
			}
			message = report.DefaultMessage(v.Pos.Filename, v.Pos.Line, v.FromLinter, text)
		}

		data := report.Content{
			ID:        report.NewID(v.Pos.Filename, v.Pos.Line, message),
			Linter:    v.FromLinter,
			FilePath:  v.Pos.Filename,
			LineNum:   v.Pos.Line,
			ColumnNum: v.Pos.Column,
			Message:   message,
		}
		contents = append(contents, data)
	}
	return contents, nil
}
