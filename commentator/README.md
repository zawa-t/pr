# commentator
commentator は、GitHub のプルリクエストに自動的にコメントを追加するツールです。

## 機能
プルリクエストの自動コメント: プルリクエストが作成または更新された際に、指定したコメントを自動的に追加します。

## 使用方法
1. 設定の確認: リポジトリの設定で、Actions の権限が「Read and write permissions」になっていることを確認してください。
2. ワークフローの設定: .github/workflows ディレクトリに適切なワークフローを配置し、プルリクエストのイベントに応じて commentator を実行するよう設定します。

## コマンドラインオプション
commentator は、以下のコマンドラインオプションをサポートしています。

**（必須）**
- `-n`, `--tool-name`：静的コード解析ツールの名前を指定します。このオプションは必須です。

- `-f`, `--input-format`：入力フォーマットを指定します。このオプションは必須です（指定なしの場合のデフォルト値: `text`）。
    - サポートされている値： `text`, `json`

- `-r`, `--role-name`：役割名を指定します。このオプションは必須です。
    - サポートされている値： `local-comment`、`bitbucket-pr-comment`、`github-pr-comment`、`github-check`

**（任意）**
- `-cus`, `--custom-message-format`：カスタムメッセージフォーマットを指定します。input-format が `json` の場合にのみ使用可能です。

- `-alt`, `--alternative-text`：代替テキストを指定します。

- `-t`, `--format-type`：フォーマットの種類を指定します。input-format が `json` の場合は必須です。
    - サポートされている値： `golangci-lint`

- `-efm`：エラーフォーマットのパターンを指定します。input-format が `text` の場合にのみ使用可能です。
    - 例） `%f:%l:%c: %m`

#### ＜使用例＞
input-format を `text` に設定し、エラーフォーマットを指定する場合:

```
commentator --tool-name=my-tool --input-format=text -efm="%f:%l:%c: %m" --role-name=github-pr-comment
```

input-format を `json` に設定し、フォーマットタイプを指定する場合:

```
commentator --tool-name=my-tool --input-format=json --format-type=golangci-lint --role-name=github-pr-comment
```

## 注意事項
- 権限設定: リポジトリの設定で、Actions permissions の Workflow permissions が「Read repository contents and packages permissions」になっていると、"403 Resource not accessible by integration" というエラーが発生します。そのため、設定を「Read and write permissions」に変更する必要があります。
ライセンス
このプロジェクトは MIT ライセンスの下で提供されています。詳細は LICENSE ファイルをご参照ください。

## 貢献
バグ報告や機能提案は、Issue を通じてお知らせください。プルリクエストも歓迎します。

## 開発者
zawa-t

## 参考資料
GitHub Actions の権限設定に関するドキュメント

https://vim-jp.org/vimdoc-en/quickfix.html#error-file-format
