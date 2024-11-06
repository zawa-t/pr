package golangci

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"os"

	"github.com/zawa-t/pr-commentator/src/flag"

	"github.com/zawa-t/pr-commentator/src/platform"
)

func MakeInputDatas(flagValue flag.Value, stdin *os.File) (datas []platform.Raw) {
	switch flagValue.InputFormat {
	case "txt":
		// input.Datas = readText(*flagValue)
	case "json":
		jsonData := decodeJSONData(stdin)
		if flagValue.AlternativeText != nil {
			jsonData.Issues = replaceText(*flagValue.AlternativeText, jsonData.Issues)
		}
		datas = makeInputDatas(flagValue.CustomTextFormat, jsonData.Issues)
	default:
		slog.Error("The specified input-format could not be processed because it is not supported.")
		os.Exit(1)
	}
	return
}

func replaceText(alternativeText string, issues []Issue) []Issue {
	for i, v := range issues {
		v.Text = alternativeText
		issues[i] = v
	}

	return issues
}

func makeInputDatas(customTextFormat *string, issues []Issue) []platform.Raw {
	datas := make([]platform.Raw, 0)
	for _, v := range issues {
		data := platform.Raw{
			Linter:   v.FromLinter,
			FilePath: v.Pos.Filename,
			LineNum:  v.Pos.Line,
			Message:  v.Text,
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
