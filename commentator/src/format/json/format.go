package json

import (
	"os"

	"github.com/zawa-t/pr/commentator/src/flag"
	golangcilint "github.com/zawa-t/pr/commentator/src/format/json/golangci-lint"

	"github.com/zawa-t/pr/commentator/src/review"
)

func Decode(flagValue flag.Value, stdin *os.File) (contents []review.Content) {
	switch flagValue.Name {
	case "golangci-lint":
		jsonData := golangcilint.Decode(stdin)
		if flagValue.AlternativeText != nil {
			jsonData.Issues = replaceText(*flagValue.AlternativeText, jsonData.Issues)
		}
		contents = golangcilint.MakeContents(flagValue.CustomTextFormat, jsonData.Issues)
	}
	return
}

func replaceText(alternativeText string, issues []golangcilint.Issue) []golangcilint.Issue {
	for i, v := range issues {
		v.Text = alternativeText
		issues[i] = v
	}
	return issues
}
