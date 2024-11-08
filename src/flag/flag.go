package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"

	"github.com/zawa-t/pr-commentator/src/format"
	"github.com/zawa-t/pr-commentator/src/platform"
)

var usage = "Usage: pr-commentator --name=[tool name] --input-format=[input format] --platform=[platform name] < inputfile"

type Required struct {
	Name, InputFormat, Platform string
}

type Optional struct {
	CustomTextFormat, AlternativeText *string
	Local                             bool
}

type Value struct {
	Required
	Optional
}

func NewValue() (value *Value) {
	// Required
	var name string
	nameFlags := []string{"n", "name"}
	for _, f := range nameFlags {
		flag.StringVar(&name, f, "", "The tool name for static code analysis. The flag is required.")
	}

	var inputFormat string
	inputFormatFlags := []string{"f", "input-format"}
	for _, f := range inputFormatFlags {
		flag.StringVar(&inputFormat, f, "", "input format. The flag is required. json, text")
	}

	var platform string
	flag.StringVar(&platform, "platform", "", "The flag is optional.")

	// Optional
	var customTextFormat string
	customTextFormatFlags := []string{"cus", "custom-text-format"}
	for _, f := range customTextFormatFlags {
		flag.StringVar(&customTextFormat, f, "", "The flag is optional.")
	}

	var alternativeText string
	alternativeTextFlags := []string{"alt", "alternative-text"}
	for _, f := range alternativeTextFlags {
		flag.StringVar(&alternativeText, f, "", "The flag is optional.")
	}

	var local bool
	flag.BoolVar(&local, "local", false, "The flag is optional.")

	flag.Parse()

	value = &Value{
		Required: Required{
			Name:        name,
			InputFormat: inputFormat,
			Platform:    platform,
		},
	}

	if customTextFormat != "" {
		if value.InputFormat == format.JSON { // NOTE: customTextFormat は json 形式の場合のみ利用可能
			value.CustomTextFormat = &customTextFormat
		} else {
			slog.Warn("If input-format flag is not in json format, customTextFormat cannot be used.")
		}
	}

	if alternativeText != "" {
		value.AlternativeText = &alternativeText
	}

	value.Local = local

	value.validate()
	return
}

func (r *Required) validate() {
	if r.Name == "" || r.InputFormat == "" || r.Platform == "" {
		slog.Error(usage)
		os.Exit(1)
	}

	allowedInputFormats := []string{format.Text, format.JSON}
	if !slices.Contains(allowedInputFormats, r.InputFormat) {
		slog.Error("The specified input-format is not supported.", "input-format", r.InputFormat)
		os.Exit(1)
	}

	if r.InputFormat == format.JSON {
		allowedNames := []string{"golangci-lint"}
		if !slices.Contains(allowedNames, r.Name) {
			slog.Error("The specified tool cannot use json format data.", "name", r.Name)
			os.Exit(1)
		}
	}

	allowedPlatforms := []string{platform.Bitbucket, platform.Github}
	if !slices.Contains(allowedPlatforms, r.Platform) {
		slog.Error("The specified platform is not supported.", "platform", r.Platform)
		os.Exit(1)
	}
}
