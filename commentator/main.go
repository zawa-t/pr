package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/zawa-t/pr/commentator/src/flag"
	"github.com/zawa-t/pr/commentator/src/format"
	"github.com/zawa-t/pr/commentator/src/format/json"
	"github.com/zawa-t/pr/commentator/src/format/text"
	"github.com/zawa-t/pr/commentator/src/platform"
	"github.com/zawa-t/pr/commentator/src/platform/bitbucket"
	bitbucketClient "github.com/zawa-t/pr/commentator/src/platform/bitbucket/client"
	"github.com/zawa-t/pr/commentator/src/platform/github"
	githubClient "github.com/zawa-t/pr/commentator/src/platform/github/client"
	"github.com/zawa-t/pr/commentator/src/platform/http"
	"github.com/zawa-t/pr/commentator/src/platform/local"
)

/*
以下、動作確認用コマンド
```
$ go build -o pr-commentator

<json>
$ ./pr-commentator -n=golangci-lint -f=json -t=golangci-lint --platform=local < sample/sample.json

<text>
$ ./pr-commentator -n=golangci-lint -efm="%f:%l:%c: %m" --platform=local < sample/golangci-lint_line-number.txt
```
*/

/*
TODO:
・githubにanotationコメントを入れられるようにする
・githubも同じコメントは1回しか入らないようにする
・出力されるログおよびログレベルの整理（slog でカスタムの JSON フォーマッタを作成含む）
・httpパッケージまわりの整備
・BitbucketのGetComment()の並行処理化
・環境変数不足時のログが必要な分しかwarnが出ないように変更（githubのときはgithubで必要な環境変数だけ）
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
	data := newData(*flagValue, stdin)

	if len(data.Contents) == 0 {
		slog.Info("No comments were added. This is because there is no data to comment on.")
	} else {
		if err := newPullRequest(flagValue.PlatformName).AddComments(context.Background(), data); err != nil {
			slog.Error("Failed to add comments.", "error", err.Error())
			os.Exit(1)
		}

		slog.Info("The pull request comments were successfully added.")
	}
}

func newData(flagValue flag.Value, stdin *os.File) platform.Data {
	data := platform.Data{
		Name: flagValue.Name,
	}

	switch flagValue.InputFormat {
	case format.JSON:
		data.Contents = json.Decode(flagValue, stdin)
	case format.Text:
		data.Contents = text.Read(flagValue, stdin)
	default:
		slog.Error("The specified input-format is not supported.")
		os.Exit(1)
	}
	return data
}

func newPullRequest(platformName string) (pr *platform.PullRequest) {
	switch platformName {
	case platform.Local:
		pr = platform.NewPullRequest(local.NewReview())
	case platform.Github:
		pr = platform.NewPullRequest(github.NewReview(githubClient.NewCustomClient(http.NewClient())))
	case platform.Bitbucket:
		pr = platform.NewPullRequest(bitbucket.NewReview(bitbucketClient.NewCustomClient(http.NewClient())))
	default:
		slog.Error("Unsupported platform was set.")
		os.Exit(1)
	}
	return
}
