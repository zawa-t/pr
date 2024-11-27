package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/zawa-t/pr/src/dependency"
	"github.com/zawa-t/pr/src/flag"
	"github.com/zawa-t/pr/src/format"
	"github.com/zawa-t/pr/src/format/json"
	"github.com/zawa-t/pr/src/format/text"
	"github.com/zawa-t/pr/src/report"
)

/*
以下、動作確認用コマンド
```
$ go build -o reporter

<json>
$ ./reporter -n=golangci-lint -f=json -t=golangci-lint -r=local-comment < sample/golangci-lint_json.json

<text>
$ ./reporter -n=golangci-lint -efm="%f:%l:%c: %m" -r=local-comment < sample/golangci-lint_line-number.txt
```
*/

/*
TODO:
・-vオプションでのバージョン表示
・github-pr-checkとgithub-checkの整備
・出力されるログおよびログレベルの整理（slog でカスタムの JSON フォーマッタを作成含む） ※出力されるエラーの整理も
・httpパッケージまわりの整備
・BitbucketのGetComment()の並行処理化
・githubPRReviewerの開発
*/

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	stdin, err := stdin()
	if err != nil {
		return err
	}

	flagValue := flag.NewValue()
	data := newData(*flagValue, stdin)

	if len(data.Contents) == 0 {
		slog.Info("The processing is complete, and no comments were added because there was no data to comment on.")
		return nil
	}

	if err := dependency.NewReporter(flagValue.Role).Report(context.Background(), data); err != nil {
		return fmt.Errorf("failed to add comments: %w", err)
	}
	slog.Info("The processing is complete, and comments were added.")
	return nil
}

func stdin() (*os.File, error) {
	stdin := os.Stdin
	stat, err := os.Stdin.Stat() // MEMO: 標準入力の「ファイル情報（ファイルのモードやサイズ、変更日時など）」取得
	if err != nil {
		return nil, fmt.Errorf("stdin could not be verified: %w", err)
	}

	// MEMO: stat.Mode()を実行することでファイルのモード情報（ファイルの種類やアクセス権）を取得。それによって設定される os.ModeCharDevice の値を用いて、
	// 入力がキャラクタデバイス（通常、ターミナル）であるか否かを確認。現時点では、標準入力がパイプやリダイレクトのみ受け付けたいため、ターミナルからの入力の場合（0 でない場合）は処理終了。
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, errors.New("only data from standard input can be accepted")
	}
	return stdin, nil
}

func newData(flagValue flag.Value, stdin io.Reader) report.Data {
	data := report.Data{
		Name: flagValue.ToolName,
	}

	switch flagValue.InputFormat {
	case format.JSON:
		config, err := json.NewConfig(flagValue.ToolName, flagValue.FormatType, flagValue.CustomMessageFormat, flagValue.AlternativeText)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		data.Contents, err = json.Decode(stdin, *config)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	case format.Text:
		config, err := text.NewConfig(flagValue.ToolName, flagValue.ErrorFormat, flagValue.AlternativeText)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		data.Contents, err = text.Read(stdin, *config)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	default:
		slog.Error("The specified input-format is not supported.")
		os.Exit(1)
	}
	return data
}
