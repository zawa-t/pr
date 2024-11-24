package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"

	"github.com/zawa-t/pr/commentator/src/format"
	"github.com/zawa-t/pr/commentator/src/report/role"
)

var usage = "Usage: commentator --tool-name=[tool name] --input-format=[input format] --role-name=[role name]"

type useableFlag struct {
	toolName         string
	inputFormat      string
	roleName         string
	customTextFormat string
	alternativeText  string
	formatType       string
	errorFormat      string
}

func newUseableFlag() *useableFlag {
	useableFlag := new(useableFlag)

	// Required
	toolNameFlags := []string{"n", "tool-name"}
	for _, f := range toolNameFlags {
		flag.StringVar(&useableFlag.toolName, f, "", "The tool name for static code analysis. The flag is required.")
	}

	inputFormatFlags := []string{"f", "input-format"}
	for _, f := range inputFormatFlags {
		flag.StringVar(&useableFlag.inputFormat, f, "text", "input format. The flag is required. json, text")
	}

	roleFlags := []string{"r", "role-name"}
	for _, f := range roleFlags {
		flag.StringVar(&useableFlag.roleName, f, "", "The flag is required.")
	}

	// Optional
	customTextFormatFlags := []string{"cus", "custom-text-format"}
	for _, f := range customTextFormatFlags {
		flag.StringVar(&useableFlag.customTextFormat, f, "", "The flag is optional. input-format が json の場合にのみ使用可能")
	}

	alternativeTextFlags := []string{"alt", "alternative-text"}
	for _, f := range alternativeTextFlags {
		flag.StringVar(&useableFlag.alternativeText, f, "", "The flag is optional.")
	}

	formatTypeFlags := []string{"t", "format-type"}
	for _, f := range formatTypeFlags {
		flag.StringVar(&useableFlag.formatType, f, "", "format type. The flag is optional. input-format が json の場合は必須。 golangci-lint")
	}

	flag.StringVar(&useableFlag.errorFormat, "efm", "", "Error format pattern. input-format が text の場合にのみ使用可能。%f:%l:%c: %m")

	flag.Parse()

	useableFlag.validate()
	return useableFlag
}

func (f *useableFlag) validate() {
	if f.toolName == "" || f.inputFormat == "" || f.roleName == "" {
		slog.Error(usage)
		os.Exit(1)
	}

	allowedInputFormats := []string{format.Text, format.JSON}
	if !slices.Contains(allowedInputFormats, f.inputFormat) {
		slog.Error("The specified input-format is not supported.", "input-format", f.inputFormat)
		os.Exit(1)
	}

	if f.inputFormat == format.Text && f.errorFormat == "" {
		slog.Error("If the input-format flag is text, efm flag must be specified.")
		os.Exit(1)
	}

	if f.inputFormat == format.JSON {
		if f.formatType == "" {
			slog.Error("If the input-format flag is json, format-type flag must be specified.")
			os.Exit(1)
		}

		allowedFormatTypes := []string{"golangci-lint"}
		if !slices.Contains(allowedFormatTypes, f.formatType) {
			slog.Error("The specified format-type is not supported.", "format-type", f.formatType)
			os.Exit(1)
		}
	}

	_, ok := role.NameList[f.roleName]
	if !ok {
		slog.Error("The specified role is not supported.", "role", f.roleName)
		os.Exit(1)
	}
}

type Required struct {
	ToolName, InputFormat string
	Role                  int
}

type Optional struct {
	CustomTextFormat, AlternativeText, ErrorFormat, FormatType *string
}

type Value struct {
	Required
	Optional
}

func NewValue() (value *Value) {
	useableFlag := newUseableFlag()

	value = &Value{
		Required: Required{
			ToolName:    useableFlag.toolName,
			InputFormat: useableFlag.inputFormat,
			Role:        role.NameList[useableFlag.roleName],
		},
	}

	if useableFlag.customTextFormat != "" {
		if value.InputFormat == format.JSON { // NOTE: customTextFormat は json 形式の場合のみ利用可能
			value.CustomTextFormat = &useableFlag.customTextFormat
		} else {
			slog.Warn("If input-format flag is not in json format, customTextFormat cannot be used.")
		}
	}

	if useableFlag.alternativeText != "" {
		value.AlternativeText = &useableFlag.alternativeText
	}

	if useableFlag.errorFormat != "" {
		value.ErrorFormat = &useableFlag.errorFormat
	}

	if useableFlag.formatType != "" {
		value.FormatType = &useableFlag.formatType
	}

	return
}
