package main

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/zawa-t/pr/commentator/src/flag"
	"github.com/zawa-t/pr/commentator/src/format"
	"github.com/zawa-t/pr/commentator/src/format/json"
	"github.com/zawa-t/pr/commentator/src/format/text"
	bitbucketClient "github.com/zawa-t/pr/commentator/src/platform/bitbucket/client"
	githubClient "github.com/zawa-t/pr/commentator/src/platform/github/client"
	"github.com/zawa-t/pr/commentator/src/platform/http"
	"github.com/zawa-t/pr/commentator/src/review"
	"github.com/zawa-t/pr/commentator/src/review/role"
)

/*
以下、動作確認用コマンド
```
$ go build -o pr-postman

<json>
$ ./pr-postman -n=golangci-lint -f=json -t=golangci-lint -r=local-commentator < sample/sample.json

<text>
$ ./pr-postman -n=golangci-lint -efm="%f:%l:%c: %m" -r=local-commentator < sample/golangci-lint_line-number.txt
```
*/

/*
TODO:
・github-pr-checkerとgithub-checkerの整備
・テスト追加
・出力されるログおよびログレベルの整理（slog でカスタムの JSON フォーマッタを作成含む） ※出力されるエラーの整理も
・httpパッケージまわりの整備
・BitbucketのGetComment()の並行処理化
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
		if err := newReviewer(flagValue.Role).Review(context.Background(), data); err != nil {
			slog.Error("Failed to add comments.", "error", err.Error())
			os.Exit(1)
		}

		slog.Info("completed.")
	}
}

func newData(flagValue flag.Value, stdin io.Reader) review.Data {
	data := review.Data{
		Name: flagValue.ToolName,
	}

	switch flagValue.InputFormat {
	case format.JSON:
		config, err := json.NewConfig(flagValue.ToolName, flagValue.FormatType, flagValue.CustomTextFormat, flagValue.AlternativeText)
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

func newReviewer(roleName string) (reviewer review.Reviewer) {
	switch roleName {
	case role.LocalCommentator:
		reviewer = role.NewLocalCommentator()
	case role.BitbucketPRCommentator:
		reviewer = role.NewBitbucketPRCommentator(bitbucketClient.NewCustomClient(http.NewClient()))
	case role.GithubPRCommentator:
		reviewer = role.NewGithubPRCommentator(githubClient.NewCustomClient(http.NewClient()))
	case role.GithubPRChecker:
		reviewer = role.NewGithubPRChecker(githubClient.NewCustomClient(http.NewClient()))
	case role.GithubChecker:
		reviewer = role.NewGithubChecker(githubClient.NewCustomClient(http.NewClient()))
	default:
		slog.Error("Unsupported role was set.")
		os.Exit(1)
	}
	return
}
