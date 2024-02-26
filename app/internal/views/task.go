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
	msg.WriteString("**Цели**\n")

	for _, goal := range goals {
		msg.WriteString(goalView(&goal))
	}

	return msg.String()
}

func goalView(goal *models.Goal) string {
	return fmt.Sprintf("*%s*\nТрек: %s\n\t%s\n", goal.Name, goal.Track, goal.Description)
}

func GetTasks(tasks []models.Task) string {
	if len(tasks) == 0 {
		return "Ты выполнил все свои задачи!"
	}
	var msg strings.Builder
	msg.WriteString("*Задачи*:\n\n")

	for i, task := range tasks {
		msg.WriteString(fmt.Sprintf("%d. ", i+1))
		msg.WriteString(taskView(&task))
		msg.WriteString("\n")
	}

	return msg.String()
}

func taskView(task *models.Task) string {
	return fmt.Sprintf(
		"%s\n"+
			"    Описание: %s\n"+
			"    Сторипоинты: %d\n"+
			"    Создана: %s\n", task.Name, task.Description, task.StoryPoints, task.CreatedAt.Format("2006-01-02 15:04:05"))
}
