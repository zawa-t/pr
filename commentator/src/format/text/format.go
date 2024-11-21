package text

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zawa-t/pr/commentator/src/flag"
	"github.com/zawa-t/pr/commentator/src/review"
)

// errorformat 文字列から正規表現を生成する関数
func convertErrorFormatToRegex(efm string) (*regexp.Regexp, error) {
	efmPatterns := map[string]string{
		"%f": `(?P<File>(?:\./)?[a-zA-Z0-9_\-/]+(?:\.[a-zA-Z0-9]+)?)`, // ファイルパス形式の文字列にマッチ ex)「main.go」「./src/main.go」「src/dir/file.py」
		"%l": `(?P<Line>\d+)`,                                         // 数字にマッチ
		"%c": `(?P<Column>\d+)`,                                       // 数字にマッチ
		"%m": `(?P<Message>.+)`,                                       // 任意の文字列にマッチ
	}

	regexPattern := efm
	for placeholder, regexPart := range efmPatterns {
		regexPattern = strings.ReplaceAll(regexPattern, placeholder, regexPart)
	}

	return regexp.Compile("^" + regexPattern + "$")
}

// efm パターンでファイルをパースして Issue を抽出する関数
func Read(flagValue flag.Value, stdin *os.File) []review.Content {
	if flagValue.ErrorFormat == nil {
		slog.Error("errorformatの指定がありません")
		os.Exit(1)
	}

	efm := *flagValue.ErrorFormat
	regex, err := convertErrorFormatToRegex(efm)
	if err != nil {
		slog.Error("errorformat を正規表現に変換できません", "error", err.Error())
		os.Exit(1)
	}

	contents := make([]review.Content, 0)
	var currentContent *review.Content

	scanner := bufio.NewScanner(stdin)

	lineCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// 正規表現によるエラーメッセージ行の解析
		if matches := regex.FindStringSubmatch(line); len(matches) > 0 {
			// エラーメッセージ行の場合、新しい Issue を初期化
			if currentContent != nil {
				contents = append(contents, *currentContent)
			}
			currentContent = &review.Content{}
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
					currentContent.LineNum = toUint(value)
				case "Column":
					currentContent.ColumnNum = toUint(value)
				case "Message":
					currentContent.Text = value
				}
			}
			currentContent.Linter = flagValue.Name
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
		slog.Error("読み込みエラー", "error", err.Error())
		os.Exit(1)
	}

	if len(contents) == 0 {
		slog.Info("No rows matched the specified errorformat.")
	}

	return contents
}

// 文字列を整数に変換するためのヘルパー関数
func toUint(s string) uint {
	num, err := strconv.Atoi(s)
	if err != nil {
		slog.Error("failed to exec strconv.Atoi()", "error", err.Error())
		os.Exit(1)
	}
	return uint(num)
}
