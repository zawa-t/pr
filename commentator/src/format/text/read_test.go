package text_test

import (
	"bytes"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zawa-t/pr/commentator/src/format/text"
	"github.com/zawa-t/pr/commentator/src/review"
	"github.com/zawa-t/pr/commentator/src/test/helper"
)

func Test_Text_NewConfig(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		type input struct {
			toolName        string
			errorFormat     *string
			alternativeText *string
		}
		type testCase struct {
			name     string
			input    input
			expected *text.Config
		}

		tests := []testCase{
			{
				name: "errorFormatに `%f:%l:%c: %m` を指定",
				input: input{
					toolName:    "Tool name",
					errorFormat: helper.ToPtr("%f:%l:%c: %m"),
				},
				expected: &text.Config{
					ToolName:    "Tool name",
					ErrorFormat: "%f:%l:%c: %m",
				},
			},
			{
				name: "errorFormatに任意の文字列を指定",
				input: input{
					toolName:    "Tool name",
					errorFormat: helper.ToPtr("任意の文字列"),
				},
				expected: &text.Config{
					ToolName:    "Tool name",
					ErrorFormat: "任意の文字列",
				},
			},
			{
				name: "alternativeTextに任意の文字列を指定",
				input: input{
					toolName:        "Tool name",
					errorFormat:     helper.ToPtr("任意の文字列"),
					alternativeText: helper.ToPtr("alternativeText"),
				},
				expected: &text.Config{
					ToolName:        "Tool name",
					ErrorFormat:     "任意の文字列",
					AlternativeText: helper.ToPtr("alternativeText"),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				config, err := text.NewConfig(tt.input.toolName, tt.input.errorFormat, tt.input.alternativeText)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, config)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		type input struct {
			toolName        string
			errorFormat     *string
			alternativeText *string
		}
		type testCase struct {
			name  string
			input input
		}

		tests := []testCase{
			{
				name: "toolNameが空文字",
				input: input{
					toolName:    "",
					errorFormat: helper.ToPtr("%f:%l:%c: %m"),
				},
			},
			{
				name: "errorFormatが空文字",
				input: input{
					toolName:    "Tool name",
					errorFormat: helper.ToPtr(""),
				},
			},
			{
				name: "errorFormatがnil",
				input: input{
					toolName:    "Tool name",
					errorFormat: nil,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				config, err := text.NewConfig(tt.input.toolName, tt.input.errorFormat, tt.input.alternativeText)
				assert.Nil(t, config)
				assert.Error(t, err)
			})
		}
	})
}

func Test_Read(t *testing.T) {
	type input struct {
		config text.Config
		text   string
	}
	type testCase struct {
		name     string
		input    input
		expected []review.Content
	}

	testExecutor := func(tests []testCase) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				stdin := bytes.NewBufferString(tt.input.text)

				contents, err := text.Read(stdin, tt.input.config)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, contents)
				}
			})
		}
	}

	t.Run("正常系 - 指定されたエラーフォーマットと一致するテキストがあった場合、その内容を反映させた[]review.Content型のスライスを返す", func(t *testing.T) {
		tests := []testCase{
			{
				name: "エラーフォーマットが `%f:%l:%c: %m` の場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text: `example.go:10:5: Error message here
example.go:20:6: Another error
`,
				},
				expected: []review.Content{
					{
						ID:        review.NewID("example.go", 10, review.DefaultMessage("example.go", 10, "TestLinter", "Error message here")),
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						CodeLine:  "",
						Indicator: "",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 10, "TestLinter", "Error message here"),
					},
					{
						ID:        review.NewID("example.go", 20, review.DefaultMessage("example.go", 20, "TestLinter", "Another error")),
						FilePath:  "example.go",
						LineNum:   20,
						ColumnNum: 6,
						CodeLine:  "",
						Indicator: "",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 20, "TestLinter", "Another error"),
					},
				},
			},
			{
				name: "config.AlternativeTextがある場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", helper.ToPtr("AlternativeText")),
					text: `example.go:10:5: Error message here
example.go:20:6: Another error
`,
				},
				expected: []review.Content{
					{
						ID:        review.NewID("example.go", 10, review.DefaultMessage("example.go", 10, "TestLinter", "AlternativeText")),
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						CodeLine:  "",
						Indicator: "",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 10, "TestLinter", "AlternativeText"),
					},
					{
						ID:        review.NewID("example.go", 20, review.DefaultMessage("example.go", 20, "TestLinter", "AlternativeText")),
						FilePath:  "example.go",
						LineNum:   20,
						ColumnNum: 6,
						CodeLine:  "",
						Indicator: "",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 20, "TestLinter", "AlternativeText"),
					},
				},
			},
			{
				name: "エラーフォーマットが `%f:%l:%c: %m` かつCodeLine、Indicatorがある場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text: `example.go:10:5: Error message here
var a = 10
        ^
example.go:20:6: Another error
var b = 20
        ^
`,
				},
				expected: []review.Content{
					{
						ID:        review.NewID("example.go", 10, review.DefaultMessage("example.go", 10, "TestLinter", "Error message here")),
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						CodeLine:  "var a = 10",
						Indicator: "        ^",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 10, "TestLinter", "Error message here"),
					},
					{
						ID:        review.NewID("example.go", 20, review.DefaultMessage("example.go", 20, "TestLinter", "Another error")),
						FilePath:  "example.go",
						LineNum:   20,
						ColumnNum: 6,
						CodeLine:  "var b = 20",
						Indicator: "        ^",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 20, "TestLinter", "Another error"),
					},
				},
			},
			{
				name: "エラーフォーマットが `%f:%l: %m` かつCodeLine、Indicatorがある場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l: %m", nil),
					text: `example.go:10: Error message here
var a = 10
        ^
example.go:20: Another error
var b = 20
        ^
`,
				},
				expected: []review.Content{
					{
						ID:        review.NewID("example.go", 10, review.DefaultMessage("example.go", 10, "TestLinter", "Error message here")),
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 0,
						CodeLine:  "var a = 10",
						Indicator: "        ^",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 10, "TestLinter", "Error message here"),
					},
					{
						ID:        review.NewID("example.go", 20, review.DefaultMessage("example.go", 20, "TestLinter", "Another error")),
						FilePath:  "example.go",
						LineNum:   20,
						ColumnNum: 0,
						CodeLine:  "var b = 20",
						Indicator: "        ^",
						Linter:    "TestLinter",
						Message:   review.DefaultMessage("example.go", 20, "TestLinter", "Another error"),
					},
				},
			},
		}

		// t.Logf("Input content: %s", input)
		testExecutor(tests)
	})

	t.Run("正常系 - 正規表現にマッチするテキストがない場合、空のスライス（[]review.Content{}）を返す", func(t *testing.T) {
		tests := []testCase{
			{
				name: "テキストの内容がエラーフォーマットの形式と一致しているものの、%fにあたる部分がファイルパス形式（yyy.拡張子）ではない場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `example:10:5: Error message here`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致しているものの、%lにあたる部分が数値ではない場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `example.go:xxxx:5: Error message here`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致しているものの、%cにあたる部分が数値ではない場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `example.go:10:xxxx: Error message here`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致していない場合",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `xxxx`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致していない場合（指定した%fにあたる部分がない）",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `10:5: Error message here`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致していない場合（指定した%lまたは%cにあたる部分が両方ない）",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `example.go:5: Error message here`,
				},
				expected: []review.Content{},
			},
			{
				name: "テキストの内容がエラーフォーマットの形式と一致していない場合（指定した%mにあたる部分がない）",
				input: input{
					config: NewConfigWithNoError("TestLinter", "%f:%l:%c: %m", nil),
					text:   `example.go:10:5`,
				},
				expected: []review.Content{},
			},
		}
		testExecutor(tests)
	})
}

func NewConfigWithNoError(toolName string, errorFormat string, alternativeText *string) text.Config {
	config, err := text.NewConfig(toolName, &errorFormat, alternativeText)
	if err != nil {
		slog.Error("Faild to exec text.NewConfig()")
		os.Exit(1)
	}
	return *config
}
