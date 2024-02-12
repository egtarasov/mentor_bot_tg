package commands

import "fmt"

const (
	CancelMessage = "Отмена"
)

var (
	ErrCanceled = fmt.Errorf("the command was canceled by the user")
)
