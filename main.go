package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/zawa-t/pr-commentator/src/env"
	"github.com/zawa-t/pr-commentator/src/flag"
	"github.com/zawa-t/pr-commentator/src/linter/golangci"
	"github.com/zawa-t/pr-commentator/src/linter/txt"
	"github.com/zawa-t/pr-commentator/src/log"
	"github.com/zawa-t/pr-commentator/src/platform"
	"github.com/zawa-t/pr-commentator/src/platform/bitbucket"
	bitbucketClient "github.com/zawa-t/pr-commentator/src/platform/bitbucket/client"
	"github.com/zawa-t/pr-commentator/src/platform/github"
	githubClient "github.com/zawa-t/pr-commentator/src/platform/github/client"
	"github.com/zawa-t/pr-commentator/src/platform/http"
	"github.com/zawa-t/pr-commentator/src/test/custommock"
)

/*
以下、動作確認用コマンド
```
$ go build -o pr-comment
$ ./pr-comment -n=golangci-lint -ext=json --platform=bitbucket < sample.json
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

	data := platform.Data{
		Name: flagValue.Name,
	}

	switch flagValue.Name {
	case "golangci-lint":
		data.RawDatas = golangci.MakeInputDatas(*flagValue, stdin)
	default:
		data.RawDatas = txt.Read(*flagValue, stdin)
	}
	log.PrintJSON("platform.Data", data)

	if err := newPullRequest(flagValue.Platform).AddComments(context.Background(), data); err != nil {
		slog.Error("Failed to add comments.", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("The pull request comments were successfully added.")
}

func newPullRequest(pf string) (pr *platform.PullRequest) {
	switch pf {
	case platform.Bitbucket:
		if env.Env.IsLocal() {
			pr = platform.NewPullRequest(bitbucket.NewPullRequest(custommock.DefaultBitbucketReview))
		} else {
			pr = platform.NewPullRequest(bitbucket.NewPullRequest(bitbucketClient.NewCustomClient(http.NewClient())))
		}
	case platform.Github:
		if env.Env.IsLocal() {
			pr = platform.NewPullRequest(github.NewPullRequest(custommock.DefaultGithubReview))
		} else {
			pr = platform.NewPullRequest(github.NewPullRequest(githubClient.NewCustomClient(http.NewClient())))
		}
	default:
		slog.Error("Unsupported platform was set.")
		os.Exit(1)
	}
	return
}
