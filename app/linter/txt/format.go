package txt

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/zawa-t/pr-commentator/flag"
	"github.com/zawa-t/pr-commentator/platform"
)

func Read(flagValue flag.Value, stdin *os.File) []platform.Raw {
	datas := make([]platform.Raw, 0)

	var last bool
	reader := bufio.NewReader(stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF { // TODO: err != io.EOF に errors.Is() が使えるやり方を検討
			slog.Error("Error reading from stdin.", "error", err.Error())
			os.Exit(1)
		}
		if err == io.EOF { // TODO: err != io.EOF に errors.Is() が使えるやり方を検討
			if line != "" {
				last = true
			} else {
				break
			}
		}

		count := 0
		for _, v := range line {
			if string(v) == ":" {
				count++
			}
		}
		if count != 2 {
			slog.Error("For txt files, two colons are required per line.")
			os.Exit(1)
		}

		columns := strings.Split(strings.TrimSpace(line), ":") // ex) For example, line is "src/user.go:30:test message\n".

		lineNum, err := strconv.Atoi(columns[1])
		if err != nil {
			slog.Error("Error convert to int.", "error", err.Error())
			os.Exit(1)
		}

		var text string
		if flagValue.AlternativeText == nil {
			text = columns[2]
		} else {
			text = *flagValue.AlternativeText
		}

		datas = append(datas, platform.Raw{
			Linter:   flagValue.Name,
			FilePath: columns[0],
			LineNum:  uint(lineNum),
			Summary:  text,
			Details:  text,
		})

		if last {
			break
		}
	}
	return datas
}
