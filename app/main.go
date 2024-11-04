package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/zawa-t/pr-commentator/env"
	"github.com/zawa-t/pr-commentator/flag"
	"github.com/zawa-t/pr-commentator/linter/golangci"
	"github.com/zawa-t/pr-commentator/linter/txt"
	"github.com/zawa-t/pr-commentator/log"
	"github.com/zawa-t/pr-commentator/platform"
	"github.com/zawa-t/pr-commentator/platform/bitbucket"
	"github.com/zawa-t/pr-commentator/platform/bitbucket/client"
	"github.com/zawa-t/pr-commentator/platform/http"
	"github.com/zawa-t/pr-commentator/test/custommock"
)

/*
以下、動作確認用コマンド
```
$ go build -o pr-comment
$ ./pr-comment -n=golangci-lint -ext=json < sample.json
```
*/

/*
TODO:
*/

func main() {
	stdin := os.Stdin
	stat, err := os.Stdin.Stat() // MEMO: 標準入力の「ファイル情報（ファイルのモードやサイズ、変更日時など）」取得
	if err != nil {
		slog.Error("Stdin could not be verified.", "error", err.Error())
		os.Exit(1)
	}

	// MEMO:
	// stat.Mode()を実行することでファイルのモード情報（ファイルの種類やアクセス権）を取得。それによって設定される os.ModeCharDevice の値を用いて、
	// 入力がキャラクタデバイス（通常、ターミナル）であるか否かを確認。現時点では、標準入力がパイプやリダイレクトのみ受け付けたいため、ターミナルからの入力の場合（0 でない場合）は処理終了。
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		slog.Error("Only data from standard input can be accepted.")
		os.Exit(1)
	}

	flagValue := flag.NewValue()
	input := platform.Input{
		Name: flagValue.Name,
	}

	switch flagValue.Name {
	case "golangci-lint":
		input.Datas = golangci.MakeInputDatas(*flagValue, stdin)
	default:
		input.Datas = txt.Read(*flagValue, stdin)
	}

	log.PrintJSON("platform.Input", input)

	var pf *platform.Platform
	switch flagValue.Platform {
	case platform.Bitbucket:
		if env.Env.IsLocal() {
			pf = platform.New(bitbucket.NewPullRequest(custommock.DefaultCustomClient))
		} else {
			pf = platform.New(bitbucket.NewPullRequest(client.NewCustomClient(http.NewClient())))
		}
	case platform.Github:
		// TODO: 処理追加
	}

	if err := pf.PullRequest.AddComments(context.Background(), input); err != nil {
		slog.Error("Failed to add comments.", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("The pull request comments were successfully added.")
}

func a() {
fmt.Println("xxx")
}

func b() {

}
