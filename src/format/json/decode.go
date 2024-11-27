package json

import (
	"fmt"
	"io"

	golangcilint "github.com/zawa-t/pr/reporter/format/json/golangci-lint"
	"github.com/zawa-t/pr/reporter/report"
)

type Config struct {
	ToolName            string
	FormatType          string
	CustomMessageFormat *string
	AlternativeText     *string
}

func NewConfig(toolName string, formatType *string, customMessageFormat, alternativeText *string) (*Config, error) {
	if toolName == "" || formatType == nil || *formatType == "" {
		return nil, fmt.Errorf("when using the json format, the values for toolName and formatType are required. toolName=%s, formatType=%v", toolName, formatType)
	}
	return &Config{
		ToolName:            toolName,
		FormatType:          *formatType,
		CustomMessageFormat: customMessageFormat,
		AlternativeText:     alternativeText,
	}, nil
}

func Decode(stdin io.Reader, config Config) (contents []report.Content, err error) {
	switch config.FormatType {
	case "golangci-lint":
		jsonData, err := golangcilint.Decode(stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to golangcilint.Decode(): %w", err)
		}
		return golangcilint.MakeContents(config.AlternativeText, config.CustomMessageFormat, jsonData.Issues)
	}
	return
}
