package json

// type GolangciLintJSON struct {
// 	Issues []Issue `json:"Issues"`
// }

// type Issue struct {
// 	FromLinter  string   `json:"FromLinter"`
// 	Text        string   `json:"Text"`
// 	Severity    string   `json:"Severity"`
// 	SourceLines []string `json:"SourceLines"`
// 	// Replacement ? // NOTE: 型や使用方法が不明のため一旦コメントアウト
// 	Pos          Pos  `json:"Pos"`
// 	ExpectNoLint bool `json:"ExpectNoLint"`
// }

// type Pos struct {
// 	Filename             string `json:"Filename"`
// 	Offset               uint   `json:"Offset"`
// 	Line                 uint   `json:"Line"`
// 	Column               uint   `json:"Column"`
// 	ExpectedNoLintLinter string `json:"ExpectedNoLintLinter"`
// }

// func decodeGolangciLintJSON(stdin *os.File) GolangciLintJSON {
// 	var jsonData GolangciLintJSON
// 	decoder := json.NewDecoder(stdin)
// 	if err := decoder.Decode(&jsonData); err != nil {
// 		slog.Error("Failed to JSON Decode.", "error", err.Error())
// 		os.Exit(1)
// 	}
// 	return jsonData
// }

// func makeContents(customTextFormat *string, issues []Issue) []platform.Content {
// 	contents := make([]platform.Content, 0)
// 	for _, v := range issues {
// 		data := platform.Content{
// 			Linter:   v.FromLinter,
// 			FilePath: v.Pos.Filename,
// 			LineNum:  v.Pos.Line,
// 			Message:  v.Text,
// 		}
// 		if customTextFormat != nil {
// 			tmpl, err := template.New("customTextFormat").Parse(*customTextFormat) // HACK: 本来はfor文のたびにParseをする必要はないため、for文の外でParseするようにできないか検討
// 			if err != nil {
// 				fmt.Println("Error parsing template:", err)
// 				os.Exit(1)
// 			}

// 			var result bytes.Buffer
// 			err = tmpl.Execute(&result, v)
// 			if err != nil {
// 				fmt.Println("Error executing template:", err)
// 				os.Exit(1)
// 			}
// 			text := result.String()
// 			data.CustomCommentText = &text // NOTE: 利用者から指定されたテキストに置き換え
// 		}
// 		contents = append(contents, data)
// 	}
// 	return contents
// }
