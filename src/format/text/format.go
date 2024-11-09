package text

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zawa-t/pr-commentator/src/flag"
	"github.com/zawa-t/pr-commentator/src/platform"
)

type ErrorFormat struct {
	File, Line, Column, Message bool
}

func Read(flagValue flag.Value, stdin *os.File) []platform.Content {
	contents := make([]platform.Content, 0)

	scanner := bufio.NewScanner(stdin)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}

	efm := *flagValue.ErrorFormat // ex) For example, line is "%f:%l:%c: %m"

	errorFormat := new(ErrorFormat)
	for _, v := range strings.Split(efm, ":") {
		switch strings.TrimSpace(v) {
		case "%f":
			errorFormat.File = true
		case "%l":
			errorFormat.Line = true
		case "%c":
			errorFormat.Column = true
		case "%m":
			errorFormat.Message = true
		default:
			slog.Error("The specified errorformat name is undefined.")
			os.Exit(1)
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		efm = strings.ReplaceAll(efm, "%f", `(?P<file>[^:]+)`)
		efm = strings.ReplaceAll(efm, "%l", `(?P<line>\d+)`)
		efm = strings.ReplaceAll(efm, "%c", `(?P<column>\d+)`)
		efm = strings.ReplaceAll(efm, "%m", `(?P<message>.+)`)

		re := regexp.MustCompile(efm)
		match := re.FindStringSubmatch(line)
		if match == nil {
			slog.Error(fmt.Sprintf("Failed to parse line: %s", line))
			os.Exit(1)
		}

		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i > 0 && name != "" {
				result[name] = match[i]
			}
		}

		if errorFormat.File && result["file"] == "" {
			slog.Error("A file name was specified in errorformat, but no value was found.", "text", line)
			os.Exit(1)
		}

		if errorFormat.Line && result["line"] == "" {
			slog.Error("A line number was specified in errorformat, but no value was found.", "text", line)
			os.Exit(1)
		}

		if errorFormat.Column && result["column"] == "" {
			slog.Error("A column number was specified in errorformat, but no value was found.", "text", line)
			os.Exit(1)
		}

		if errorFormat.Message && result["message"] == "" {
			slog.Error("A message was specified in errorformat, but no value was found.", "text", line)
			os.Exit(1)
		}

		var lineNum int
		var err error
		if result["line"] != "" {
			lineNum, err = strconv.Atoi(result["line"])
			if err != nil {
				slog.Error("Faild to strconv.Atoi().")
				os.Exit(1)
			}
		}

		var columnNum int
		if result["column"] != "" {
			columnNum, err = strconv.Atoi(result["column"])
			if err != nil {
				slog.Error("Faild to strconv.Atoi().")
				os.Exit(1)
			}
		}

		var message string
		if result["message"] != "" {
			if flagValue.AlternativeText == nil {
				message = strings.TrimSpace(result["message"])
			} else {
				message = *flagValue.AlternativeText
			}
		}

		contents = append(contents, platform.Content{
			Linter:    flagValue.Name,
			FilePath:  result["file"],
			LineNum:   uint(lineNum),
			ColumnNum: uint(columnNum),
			Message:   message,
		})
	}
	return contents
}
