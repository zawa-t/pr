package json

import (
	"os"

	"github.com/zawa-t/pr-commentator/src/flag"

	"github.com/zawa-t/pr-commentator/src/platform"
)

func Decode(flagValue flag.Value, stdin *os.File) (datas []platform.Raw) {
	switch flagValue.Name {
	case "golangci-lint":
		jsonData := decodeGolangciLintJSON(stdin)
		if flagValue.AlternativeText != nil {
			jsonData.Issues = replaceText(*flagValue.AlternativeText, jsonData.Issues)
		}
		datas = makeInputDatas(flagValue.CustomTextFormat, jsonData.Issues)
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
