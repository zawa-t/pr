package json

import (
	"fmt"
	"io"

	"github.com/zawa-t/pr/commentator/src/errors"
	golangcilint "github.com/zawa-t/pr/commentator/src/format/json/golangci-lint"

	"github.com/zawa-t/pr/commentator/src/review"
)

type Config struct {
	ToolName         string
	FormatType       string
	CustomTextFormat *string
	AlternativeText  *string
}

func NewConfig(toolName string, formatType *string, customTextFormat, alternativeText *string) (*Config, error) {
	if toolName == "" || formatType == nil || *formatType == "" {
		return nil, fmt.Errorf("when using the json format, the values for toolName and formatType are required. toolName=%s, formatType=%v :%w", toolName, formatType, errors.ErrMissingRequiredParams)
	}
	return &Config{
		ToolName:         toolName,
		FormatType:       *formatType,
		CustomTextFormat: customTextFormat,
		AlternativeText:  alternativeText,
	}, nil
}

func Decode(stdin io.Reader, config Config) (contents []review.Content, err error) {
	switch config.FormatType {
	case "golangci-lint":
		jsonData, err := golangcilint.Decode(stdin)
		if err != nil {
			return nil, fmt.Errorf("failed to golangcilint.Decode(): %w", err)
		}
		return golangcilint.MakeContents(config.AlternativeText, config.CustomTextFormat, jsonData.Issues)
	}
	return
}
