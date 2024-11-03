package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"

	"github.com/zawa-t/pr-commentator/platform"
)

var usage = "Usage: pr-comment --name=[tool name] --ext=[file extension] --platform=[platform name] < inputfile"

type Required struct {
	Name, FileExtension, Platform string
}

type Optional struct {
	CustomTextFormat, AlternativeText *string
}

type Value struct {
	Required
	Optional
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
		Required: Required{
			Name:          name,
			FileExtension: fileExtension,
			Platform:      platform,
		},
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

func (r *Required) validate() {
	if r.Name == "" || r.FileExtension == "" || r.Platform == "" {
		slog.Error(usage)
		os.Exit(1)
	}

	allowedExtensions := []string{"txt", "json"}
	if !slices.Contains(allowedExtensions, r.FileExtension) {
		slog.Error("The specified extension is not supported.")
		os.Exit(1)
	}

	allowedPlatforms := []string{platform.Bitbucket, platform.Github}
	if !slices.Contains(allowedPlatforms, r.Platform) {
		slog.Error("The specified platform is not supported.")
		os.Exit(1)
	}
}
