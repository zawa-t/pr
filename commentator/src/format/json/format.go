package json

import (
	"os"

	"github.com/zawa-t/pr-commentator/commentator/src/flag"

	"github.com/zawa-t/pr-commentator/commentator/src/platform"
)

func Decode(flagValue flag.Value, stdin *os.File) (contents []platform.Content) {
	switch flagValue.Name {
	case "golangci-lint":
		jsonData := decodeGolangciLintJSON(stdin)
		if flagValue.AlternativeText != nil {
			jsonData.Issues = replaceText(*flagValue.AlternativeText, jsonData.Issues)
		}
		contents = makeContents(flagValue.CustomTextFormat, jsonData.Issues)
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
