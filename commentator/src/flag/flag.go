package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"

	"github.com/zawa-t/pr/commentator/src/format"
	"github.com/zawa-t/pr/commentator/src/platform"
)

var usage = "Usage: commentator --name=[tool name] --input-format=[input format] --reviewer=[reviewer name] < inputfile"

type Required struct {
	Name, InputFormat, Reviewer string
}

type Optional struct {
	CustomTextFormat, AlternativeText, ErrorFormat, FormatType *string
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
		flag.StringVar(&inputFormat, f, "text", "input format. The flag is required. json, text")
	}

	var reviewer string
	reviewerFlags := []string{"r", "reviewer"}
	for _, f := range reviewerFlags {
		flag.StringVar(&reviewer, f, "", "The flag is required.")
	}

	// Optional
	var customTextFormat string
	customTextFormatFlags := []string{"cus", "custom-text-format"}
	for _, f := range customTextFormatFlags {
		flag.StringVar(&customTextFormat, f, "", "The flag is optional. input-format が json の場合にのみ使用可能")
	}

	var alternativeText string
	alternativeTextFlags := []string{"alt", "alternative-text"}
	for _, f := range alternativeTextFlags {
		flag.StringVar(&alternativeText, f, "", "The flag is optional.")
	}

	var formatType string
	formatTypeFlags := []string{"t", "format-type"}
	for _, f := range formatTypeFlags {
		flag.StringVar(&formatType, f, "", "format type. The flag is optional. input-format が json の場合は必須。 golangci-lint")
	}

	var errorFormat string
	flag.StringVar(&errorFormat, "efm", "", "Error format pattern. input-format が text の場合にのみ使用可能。%f:%l:%c: %m")

	flag.Parse()

	value = &Value{
		Required: Required{
			Name:        name,
			InputFormat: inputFormat,
			Reviewer:    reviewer,
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

	if errorFormat != "" {
		value.ErrorFormat = &errorFormat
	}

	if formatType != "" {
		value.FormatType = &formatType
	}

	value.validate()
	return
}

func (v *Value) validate() {
	if v.Name == "" || v.InputFormat == "" || v.Reviewer == "" {
		slog.Error(usage)
		os.Exit(1)
	}

	allowedInputFormats := []string{format.Text, format.JSON}
	if !slices.Contains(allowedInputFormats, v.InputFormat) {
		slog.Error("The specified input-format is not supported.", "input-format", v.InputFormat)
		os.Exit(1)
	}

	if v.InputFormat == format.Text && v.ErrorFormat == nil {
		slog.Error("If the input-format flag is text, efm flag must be specified.")
		os.Exit(1)
	}

	if v.InputFormat == format.JSON {
		if v.FormatType == nil {
			slog.Error("If the input-format flag is json, format-type flag must be specified.")
			os.Exit(1)
		}

		allowedFormatTypes := []string{"golangci-lint"}
		if !slices.Contains(allowedFormatTypes, *v.FormatType) {
			slog.Error("The specified format-type is not supported.", "name", *v.FormatType)
			os.Exit(1)
		}
	}

	allowedReviewers := []string{platform.BitbucketReviewer, platform.GithubReviewer, platform.LocalReviewer}
	if !slices.Contains(allowedReviewers, v.Reviewer) {
		slog.Error("The specified reviewer is not supported.", "reviewer", v.Reviewer)
		os.Exit(1)
	}
}
