package golangci

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"os"

	"github.com/zawa-t/pr-commentator/flag"

	"github.com/zawa-t/pr-commentator/platform"
)

func MakeInputDatas(flagValue flag.Value, stdin *os.File) (datas []platform.Data) {
	switch flagValue.FileExtension {
	case "txt":
		// input.Datas = readText(*flagValue)
	case "json":
		jsonData := decodeJSONData(stdin)
		if flagValue.AlternativeText != nil {
			jsonData.Issues = replaceText(*flagValue.AlternativeText, jsonData.Issues)
		}
		datas = makeInputDatas(flagValue.CustomTextFormat, jsonData.Issues)
	default:
		slog.Error("The specified extension could not be processed because it is not supported.")
		os.Exit(1)
	}
	return
}

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

func replaceText(alternativeText string, issues []Issue) []Issue {
	for i, v := range issues {
		v.Text = alternativeText
		issues[i] = v
	}

	return issues
}

func makeInputDatas(customTextFormat *string, issues []Issue) []platform.Data {
	datas := make([]platform.Data, 0)
	for _, v := range issues {
		data := platform.Data{
			Linter:   v.FromLinter,
			FilePath: v.Pos.Filename,
			LineNum:  v.Pos.Line,
			Summary:  v.Text,
			Details:  v.Text,
		}
		if customTextFormat != nil {
			tmpl, err := template.New("customTextFormat").Parse(*customTextFormat) // HACK: 本来はfor文のたびにParseをする必要はないため、for文の外でParseするようにできないか検討
			if err != nil {
				fmt.Println("Error parsing template:", err)
				os.Exit(1)
			}

			var result bytes.Buffer
			err = tmpl.Execute(&result, v)
			if err != nil {
				fmt.Println("Error executing template:", err)
				os.Exit(1)
			}
			text := result.String()
			data.CustomCommentText = &text // NOTE: 利用者から指定されたテキストに置き換え
		}
		datas = append(datas, data)
	}
	return datas
}
