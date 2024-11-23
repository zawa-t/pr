package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zawa-t/pr/commentator/src/format/json"
	"github.com/zawa-t/pr/commentator/src/test/helper"
)

func Test_json_NewConfig(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		type input struct {
			toolName         string
			formatType       *string
			customTextFormat *string
			alternativeText  *string
		}
		type testCase struct {
			name     string
			input    input
			expected *json.Config
		}

		tests := []testCase{
			{
				name: "すべての引数に値がある",
				input: input{
					toolName:         "linter",
					formatType:       helper.ToPtr("formatType"),
					customTextFormat: helper.ToPtr("customTextFormat"),
					alternativeText:  helper.ToPtr("alternativeText"),
				},
				expected: &json.Config{
					ToolName:         "linter",
					FormatType:       "formatType",
					CustomTextFormat: helper.ToPtr("customTextFormat"),
					AlternativeText:  helper.ToPtr("alternativeText"),
				},
			},
			{
				name: "customTextFormatおよびcustomTextFormatがnil",
				input: input{
					toolName:         "linter",
					formatType:       helper.ToPtr("formatType"),
					customTextFormat: nil,
					alternativeText:  nil,
				},
				expected: &json.Config{
					ToolName:         "linter",
					FormatType:       "formatType",
					CustomTextFormat: nil,
					AlternativeText:  nil,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := json.NewConfig(tt.input.toolName, tt.input.formatType, tt.input.customTextFormat, tt.input.alternativeText)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, got)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		type input struct {
			toolName         string
			formatType       *string
			customTextFormat *string
			alternativeText  *string
		}
		type testCase struct {
			name  string
			input input
		}

		tests := []testCase{
			{
				name: "toolNameが空文字",
				input: input{
					toolName:   "",
					formatType: helper.ToPtr("formatType"),
				},
			},
			{
				name: "formatTypeがnil",
				input: input{
					toolName:   "linter",
					formatType: nil,
				},
			},
			{
				name: "formatTypeが空文字",
				input: input{
					toolName:   "linter",
					formatType: helper.ToPtr(""),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := json.NewConfig(tt.input.toolName, tt.input.formatType, tt.input.customTextFormat, tt.input.alternativeText)
				assert.Nil(t, got)
				assert.Error(t, err)
			})
		}
	})
}
