package views

import (
	"fmt"
	"strings"
	"telegrambot_new_emploee/internal/models"
)

func ShowTodo(todos []models.Todo) string {
	var msg strings.Builder
	msg.WriteString("Список задач в твоем чек-листе:\n")

	listTodo(todos, &msg)

	return msg.String()
}

func listTodo(todos []models.Todo, msg *strings.Builder) {
	for i, todo := range todos {
		msg.WriteString(fmt.Sprintf("%v. %s\n", i+1, todo.Label))
	}
}

func CheckTodo(todos []models.Todo) string {
	var msg strings.Builder
	msg.WriteString("Введи номер задачи, которую ты хочешь отметить выполненной или 'Отмена'," +
		" чтобы отменить действие:\n")

	listTodo(todos, &msg)

	return msg.String()
}

func GetGoals(goals []models.Goal) string {
	if len(goals) == 0 {
		return "Ты выполнил все свои цели!"
	}
	var msg strings.Builder
	msg.WriteString("На данный момент у тебя поставлены следующие цели:\n")

	for _, goal := range goals {
		msg.WriteString(goalView(&goal))
	}

	return msg.String()
}

func goalView(goal *models.Goal) string {
	return fmt.Sprintf("*%s*\nТрек: %s\n\t%s\n", goal.Name, goal.Track, goal.Description)
}
