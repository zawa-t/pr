package text

import (
	"bufio"
	stdErr "errors"
	"fmt"
	"io"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/zawa-t/pr/reporter/src/errors"
	"github.com/zawa-t/pr/reporter/src/report"
)

type efm string

// convertToRegex 文字列から正規表現を生成する関数
func (e *efm) convertToRegex() (*regexp.Regexp, error) {
	efmPatterns := map[string]string{
		"%f": `(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+))`, // ファイルパス形式の文字列にマッチ ex)「main.go」「./src/main.go」「src/dir/file.py」
		"%l": `(?P<Line>\d+)`,                                        // 数字にマッチ
		"%c": `(?P<Column>\d+)`,                                      // 数字にマッチ
		"%m": `(?P<Text>.+)`,                                         // 任意の文字列にマッチ
	}

	regexPattern := string(*e)
	for _, splitedEFM := range strings.Split(regexPattern, ":") {
		_, ok := efmPatterns[strings.TrimSpace(splitedEFM)]
		if !ok {
			return nil, stdErr.New("unsuported errorformat is specified")
		}
	}

	for placeholder, regexPart := range efmPatterns {
		regexPattern = strings.ReplaceAll(regexPattern, placeholder, regexPart)
	}

	return regexp.Compile("^" + regexPattern + "$")
}

type Config struct {
	ToolName        string
	ErrorFormat     efm
	AlternativeText *string
}

func NewConfig(toolName string, errorFormat, alternativeText *string) (*Config, error) {
	if toolName == "" || errorFormat == nil || *errorFormat == "" {
		err := fmt.Errorf("when using the text format, the values for toolName and errorFormat are required. toolName=%s, errorFormat=%v", toolName, errorFormat)
		return nil, errors.NewAppError(errors.InvalidParams, err)
	}
	return &Config{
		ToolName:        toolName,
		ErrorFormat:     efm(*errorFormat),
		AlternativeText: alternativeText,
	}, nil
}

// efm パターンでファイルをパースして Issue を抽出する関数
func Read(stdin io.Reader, config Config) ([]report.Content, error) {
	regex, err := config.ErrorFormat.convertToRegex()
	if err != nil {
		return nil, err
	}

	var currentContent *report.Content
	lineCounter := 0

	contents := make([]report.Content, 0)
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		var text string
		if matches := regex.FindStringSubmatch(line); len(matches) > 0 {
			// エラーメッセージ行の場合、新しい Issue を初期化
			if currentContent != nil {
				contents = append(contents, *currentContent)
			}
			currentContent = &report.Content{}
			lineCounter = 1

			// キャプチャグループから Issue を生成
			for i, name := range regex.SubexpNames() {
				if i == 0 || name == "" {
					continue
				}
				value := matches[i]
				switch name {
				case "File":
					currentContent.FilePath = value
				case "Line":
					lineNum, _ := strconv.Atoi(value) // MEMO: 正規表現で数値の場合だけ値を取得しているため基本的にエラーは出ない想定
					currentContent.LineNum = uint(lineNum)
				case "Column":
					columnNum, _ := strconv.Atoi(value) // MEMO: 正規表現で数値の場合だけ値を取得しているため基本的にエラーは出ない想定
					currentContent.ColumnNum = uint(columnNum)
				case "Text":
					if config.AlternativeText != nil {
						text = *config.AlternativeText
					} else {
						text = value
					}
				}
			}
			currentContent.Linter = config.ToolName
			currentContent.Message = report.DefaultMessage(currentContent.FilePath, currentContent.LineNum, currentContent.Linter, text)
			currentContent.ID = report.NewID(currentContent.FilePath, currentContent.LineNum, currentContent.Message)
		} else if currentContent != nil && lineCounter == 1 {
			// 2行目：該当コード行
			currentContent.CodeLine = line
			lineCounter++
		} else if currentContent != nil && lineCounter == 2 {
			// 3行目：インジケータ行
			currentContent.Indicator = line
			contents = append(contents, *currentContent)

			// 初期化
			currentContent = nil
			lineCounter = 0
		}
	}

	if currentContent != nil {
		contents = append(contents, *currentContent)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(contents) == 0 {
		slog.Info("No rows matched the specified errorformat.")
	}
	return contents, nil
}
