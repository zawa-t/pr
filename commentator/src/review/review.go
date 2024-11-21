//go:generate moq -rm -out $GOPATH/app/src/test/mock/$GOFILE -pkg mock . Reviewer
package review

import (
	"context"
	"fmt"
)

// Reviewer ...
type Reviewer interface {
	Review(ctx context.Context, data Data) error
}

// Data ...
type Data struct {
	Name     string
	Contents []Content
}

// Content ...
type Content struct {
	ID        string
	Linter    string
	FilePath  string
	LineNum   uint
	ColumnNum uint
	CodeLine  string
	Indicator string
	Text      string
}

func DefaultMessage(filePath string, lineNum uint, linter string, text string) string {
	return fmt.Sprintf("[Automatic PR Comment]  \n・File: %s（%d）  \n・Linter: %s  \n・Details: %s", filePath, lineNum, linter, text) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
}

func CustomMessage(customText string) string {
	return fmt.Sprintf("[Automatic PR Comment]  \n%s", customText)
}
