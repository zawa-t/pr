//go:generate moq -rm -out $GOPATH/src/reporter/reporter/test/mock/$GOFILE -pkg mock . Reporter
package report

import (
	"context"
	"fmt"
)

// Reporter ...
type Reporter interface {
	Report(ctx context.Context, data Data) error
}

type ID string

func NewID(filePath string, lineNum uint, message Message) ID {
	return ID(fmt.Sprintf("%s:%d:%s", filePath, lineNum, message.String()))
}

func ReNewID(filePath string, lineNum uint, message string) ID {
	return ID(fmt.Sprintf("%s:%d:%s", filePath, lineNum, message))
}

type Message string

func (m Message) String() string {
	return string(m)
}

func DefaultMessage(filePath string, lineNum uint, linter string, text string) Message {
	return Message(fmt.Sprintf("[Automatic Comment]  \n・File: %s（%d）  \n・Linter: %s  \n・Details: %s", filePath, lineNum, linter, text)) // NOTE: 改行する際には、「空白2つ+`/n`（  \n）」が必要な点に注意
}

func CustomMessage(customMessageFormat string) Message {
	return Message(fmt.Sprintf("[Automatic Comment]  \n%s", customMessageFormat))
}

// Data ...
type Data struct {
	Name     string
	Contents []Content
}

// Content ...
type Content struct {
	ID        ID
	Linter    string
	FilePath  string
	LineNum   uint
	ColumnNum uint
	CodeLine  string
	Indicator string
	Message   Message
}
