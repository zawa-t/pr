package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"
)

var usage = "Usage: reviewcat --tool-name=[linter tool name] --ext=[file extension] < inputfile"

type Value struct {
	Name, FileExtension, Platform     string
	CustomTextFormat, AlternativeText *string
}

func NewValue() (value *Value) {
	// required
	var name string
	nameFlags := []string{"n", "name"}
	for _, f := range nameFlags {
		flag.StringVar(&name, f, "", "linter tool name. The flag is required.")
	}

	var fileExtension string
	fileExtFlags := []string{"ext", "extension"}
	for _, f := range fileExtFlags {
		flag.StringVar(&fileExtension, f, "", "file extension. The flag is required.")
	}

	var platform string
	flag.StringVar(&platform, "platform", "", "The flag is optional.")

	// optional
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

	flag.Parse()

	value = &Value{
		Name:          name,
		FileExtension: fileExtension,
		Platform:      platform,
	}

	if customTextFormat != "" {
		value.CustomTextFormat = &customTextFormat
	}

	if alternativeText != "" {
		value.AlternativeText = &alternativeText
	}

	value.validate()
	return
}

func (v *Value) validate() {
	if v.Name == "" || v.FileExtension == "" {
		slog.Error(usage)
		os.Exit(1)
	}

	allowedFormats := []string{"txt", "json"}
	if !slices.Contains(allowedFormats, v.FileExtension) {
		slog.Error(usage)
		os.Exit(1)
	}

	if v.AlternativeText != nil && v.CustomTextFormat != nil {
		slog.Error("Only one of AlternativeText and CustomTextFormat can be set.")
		os.Exit(1)
	}
}
