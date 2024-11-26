package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToRegex(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		type testCase struct {
			name     string
			input    efm
			expected string
		}

		tests := []testCase{
			{
				name:     "`%f` の場合",
				input:    efm("%f"),
				expected: `^(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+))$`,
			},
			{
				name:     "`%l` の場合",
				input:    efm("%l"),
				expected: `^(?P<Line>\d+)$`,
			},
			{
				name:     "`%c` の場合",
				input:    efm("%c"),
				expected: `^(?P<Column>\d+)$`,
			},
			{
				name:     "`%m` の場合",
				input:    efm("%m"),
				expected: `^(?P<Text>.+)$`,
			},
			{
				name:     "`%f:%l:%c: %m` の場合",
				input:    efm("%f:%l:%c: %m"),
				expected: `^(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+)):(?P<Line>\d+):(?P<Column>\d+): (?P<Text>.+)$`,
			},
			{
				name:     "`%f:%l` の場合",
				input:    efm("%f:%l"),
				expected: `^(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+)):(?P<Line>\d+)$`,
			},
			{
				name:     "`%f:%l:%c` の場合",
				input:    efm("%f:%l:%c"),
				expected: `^(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+)):(?P<Line>\d+):(?P<Column>\d+)$`,
			},
			{
				name:     "` %m:%c:%l:%f` の場合",
				input:    efm(" %m:%c:%l:%f"),
				expected: `^ (?P<Text>.+):(?P<Column>\d+):(?P<Line>\d+):(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+))$`,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				regex, err := tt.input.convertToRegex()
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, regex.String())
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		type testCase struct {
			name  string
			input efm
		}

		tests := []testCase{
			{
				name:  "エラーフォーマットに `%f`、`%l`、`%c`、`%m` 以外のプレースホルダーが設定されている場合",
				input: efm("%w:%x:%y: %z"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				regex, err := tt.input.convertToRegex()
				assert.Nil(t, regex)
				assert.Error(t, err)
			})
		}
	})
}
