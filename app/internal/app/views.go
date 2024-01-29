package app

import (
	"fmt"
	"sort"
	"strings"
)

func todoMessage(todos []Todo) string {
	var message strings.Builder
	message.WriteString("Список задач в твоем чек-листе:\n")

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Priority > todos[j].Priority
	})
	for i, todo := range todos {
		message.WriteString(fmt.Sprintf("%v. %s\n", i+1, todo.Label))
	}

	return message.String()
}
