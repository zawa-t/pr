package flag

import (
	"flag"
	"log/slog"
	"os"
	"slices"

	"github.com/zawa-t/pr/reporter/src/format"
	"github.com/zawa-t/pr/reporter/src/report/role"
)

var usage = `
Usage: reporter [options]

Required Flags:
	-n, --tool-name         	  Specify the tool name for static code analysis.
	-f, --input-format      	  Specify the input format (default: text).
	-r, --role-name         	  Specify the role name.

Optional Flags:
	-cus, --custom-message-format Specify the custom message format (json input only).
	-alt, --alternative-text 	  Specify the alternative text.
	-t, --format-type        	  Specify the format type (required for json input).
	-efm                     	  Specify the error format (required for text input).
`

type useableFlag struct {
	toolName            string
	inputFormat         string
	roleName            string
	customMessageFormat string
	alternativeText     string
	formatType          string
	errorFormat         string
}

func newUseableFlag() *useableFlag {
	useableFlag := new(useableFlag)

	// Required
	toolNameFlags := []string{"n", "tool-name"}
	for _, f := range toolNameFlags {
		flag.StringVar(&useableFlag.toolName, f, "", "Specify the tool name for static code analysis. The flag is required.")
	}

	inputFormatFlags := []string{"f", "input-format"}
	for _, f := range inputFormatFlags {
		flag.StringVar(&useableFlag.inputFormat, f, "text", "Specify the input format. The flag is required (default value: text).")
	}

	roleFlags := []string{"r", "role-name"}
	for _, f := range roleFlags {
		flag.StringVar(&useableFlag.roleName, f, "", "Specify the role name. The flag is required.")
	}

	// Optional
	customMessageFormatFlags := []string{"cus", "custom-message-format"}
	for _, f := range customMessageFormatFlags {
		flag.StringVar(&useableFlag.customMessageFormat, f, "", "Specify the custom message format. The flag is optional. This flag is only available when input-format is set to json.")
	}

	alternativeTextFlags := []string{"alt", "alternative-text"}
	for _, f := range alternativeTextFlags {
		flag.StringVar(&useableFlag.alternativeText, f, "", "Specify the alternative text. The flag is optional.")
	}

	formatTypeFlags := []string{"t", "format-type"}
	for _, f := range formatTypeFlags {
		flag.StringVar(&useableFlag.formatType, f, "", "Specify the format type. The flag is optional. This is required when input-format is set to json.")
	}

	flag.StringVar(&useableFlag.errorFormat, "efm", "", "Specify the error format. This is required when input-format is set to text.")

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
		slog.Error("Invalid configuration: The specified input-format is not supported.", "input-format", f.inputFormat)
		os.Exit(1)
	}

	if f.inputFormat == format.Text && f.errorFormat == "" {
		slog.Error("Invalid configuration: When the input-format flag is `text`, the `efm` flag must be specified.")
		os.Exit(1)
	}

	if f.inputFormat == format.JSON {
		if f.formatType == "" {
			slog.Error("Invalid configuration: When the input-format flag is json, format-type flag must be specified.")
			os.Exit(1)
		}

		allowedFormatTypes := []string{"golangci-lint"}
		if !slices.Contains(allowedFormatTypes, f.formatType) {
			slog.Error("Invalid configuration: The specified format-type is not supported.", "format-type", f.formatType)
			os.Exit(1)
		}
	}

	if _, ok := role.NameList[f.roleName]; !ok {
		useableRoleNames := make([]string, 0, len(role.NameList))
		for roleName := range role.NameList {
			useableRoleNames = append(useableRoleNames, roleName)
		}
		slog.Error("Invalid configuration: The specified role is not supported.", "role", f.roleName, "useableRoleNames", useableRoleNames)
		os.Exit(1)
	}
}

type Optional struct {
	CustomMessageFormat, AlternativeText, ErrorFormat, FormatType *string
}

type Value struct {
	ToolName, InputFormat string
	Role                  int
	Optional
}

func (v *Value) addOptionalValue(customMessageFormat, alternativeText, errorFormat, formatType string) {
	if customMessageFormat != "" {
		if v.InputFormat == format.JSON { // NOTE: customMessageFormat は json 形式の場合のみ利用可能
			v.CustomMessageFormat = &customMessageFormat
		} else {
			slog.Warn("When input-format flag is not in json format, customMessageFormat cannot be used.")
		}
	}

	if alternativeText != "" {
		v.AlternativeText = &alternativeText
	}

	if errorFormat != "" {
		v.ErrorFormat = &errorFormat
	}

	if formatType != "" {
		v.FormatType = &formatType
	}
}

func NewValue() (value *Value) {
	useableFlag := newUseableFlag()

	value = &Value{
		ToolName:    useableFlag.toolName,
		InputFormat: useableFlag.inputFormat,
		Role:        role.NameList[useableFlag.roleName],
	}

	value.addOptionalValue(useableFlag.customMessageFormat, useableFlag.alternativeText, useableFlag.errorFormat, useableFlag.formatType)

	return
}
