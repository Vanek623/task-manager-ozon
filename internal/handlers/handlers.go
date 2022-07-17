package handlers

import (
	"TaskAlertBot/internal/storage"
	"github.com/pkg/errors"
	"strings"
)

const (
	helpCmd = "help"
	listCmd = "list"

	addCmd = "add"
	//editCmd = "edit"
	//delCmd = "delete"
	//getCmd = "get"
)

var BadArgumentErr = errors.New("Bad argument")

func listFunc(s string) string {
	data := storage.Tasks()
	res := make([]string, 0, len(data))
	for _, c := range data {
		res = append(res, c.String())
	}

	return strings.Join(res, "\n")
}

func helpFunction(s string) string {
	return `/help - list commands\n`
}
