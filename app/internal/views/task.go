package views

import (
	"fmt"
	"strings"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/models"
	"time"
)

func GetTodo(uncompletedTodos []models.Todo, user *models.User, total int) *models.Message {
	var msg strings.Builder
	msg.WriteString("*Чек-лист*\n\n")

	percentage := float64(total-len(uncompletedTodos)) / float64(total)
	if total == 0 {
		percentage = 1
	}
	listTodo(uncompletedTodos, &msg)
	msg.WriteString("\n\n")
	progressBar(&msg, percentage)
	msg.WriteString("\n\n")
	motivationMessage(&msg, percentage, user)

	return models.NewMessageWithPhotoPath(msg.String(), user.TelegramId, config.Cfg.Tasks.PhotoPathTodos)
}

func listTodo(todos []models.Todo, msg *strings.Builder) {
	for i, todo := range todos {
		msg.WriteString(fmt.Sprintf("%d. %s\n", i+1, todo.Label))
	}
}

func CheckTodo(uncompletedTodos []models.Todo, chatId int64) *models.Message {
	var msg strings.Builder
	msg.WriteString("Введи номер задачи, которую ты хочешь отметить выполненной или 'Отмена'," +
		" чтобы отменить действие:\n\n")

	listTodo(uncompletedTodos, &msg)

	return models.NewMessage(msg.String(), chatId)
}

func CheckTask(uncompletedTasks []models.Task, chatId int64) *models.Message {
	var msg strings.Builder
	msg.WriteString("Введи номер задачи, которую ты хочешь отметить выполненной или 'Отмена'," +
		" чтобы отменить действие:\n\n")

	for i, task := range uncompletedTasks {
		msg.WriteString(fmt.Sprintf("%d. %s\n", i+1, task.Name))
	}

	return models.NewMessage(msg.String(), chatId)
}

func GetGoals(goals []models.Goal, chatId int64) *models.Message {
	if len(goals) == 0 {
		return models.NewMessage("Ты выполнил все свои цели!", chatId)
	}
	var msg strings.Builder
	msg.WriteString("**Твои Цели**\n\n")

	for _, goal := range goals {
		msg.WriteString(goalView(&goal))
		msg.WriteString("--------------------------\n\n")
	}

	return models.NewMessageWithPhotoPath(msg.String(), chatId, config.Cfg.Tasks.PhotoPathGoals)
}

func goalView(goal *models.Goal) string {
	return fmt.Sprintf(
		"Цель: %s\n\n"+
			"Трек: %s\n\n"+
			"Описание: %s\n\n", goal.Name, goal.Track, goal.Description)
}

const (
	marked   = "\U0001F315" // 🌕
	unmarked = "\U0001F311" // 🌑
)

func progressBar(msg *strings.Builder, percentage float64) {
	// Count progress.
	markedCount := int(float64(config.Cfg.Tasks.BarCount) * percentage)

	// Add to string.
	msg.WriteString("Твой текущий прогресс:\n\n")
	msg.WriteString(strings.Repeat(marked, markedCount))
	msg.WriteString(strings.Repeat(unmarked, config.Cfg.Tasks.BarCount-markedCount))
	msg.WriteString(fmt.Sprintf(" %.2f%%", percentage*100))
}

// Assume that the average percentage is equal to days employee work divided by adaptation duration.
func motivationMessage(msg *strings.Builder, percentage float64, user *models.User) {
	daysPast := time.Now().Sub(user.StartWork).Hours()
	adaptationLast := user.AdaptationEnds.Sub(user.StartWork).Hours()
	average := daysPast / adaptationLast

	if percentage >= average {
		msg.WriteString("Ты отлично справляешься! Продолжай в том же духе \U0001F525")
		return
	}

	msg.WriteString("Поторопсиь, время идет \u231B")
}

func GetTasks(tasks []models.Task, user *models.User) *models.Message {
	var msg strings.Builder
	msg.WriteString("*Задачи*\n\n")

	completed := 0
	for i, task := range tasks {
		if task.CompletedAt != nil {
			completed++
			continue
		}
		msg.WriteString(fmt.Sprintf("%d. ", i+1))
		msg.WriteString(taskView(&task))
		msg.WriteString("\n\n")
	}

	// Progress bar and motivation message.
	percentage := float64(completed) / float64(len(tasks))
	if len(tasks) == 0 {
		percentage = 1
	}
	progressBar(&msg, percentage)
	msg.WriteString("\n\n")
	motivationMessage(&msg, percentage, user)

	return models.NewMessageWithPhotoPath(msg.String(), user.TelegramId, config.Cfg.Tasks.PhotoPathTasks)
}

func taskView(task *models.Task) string {
	if task.Deadline == nil {
		return fmt.Sprintf(
			"%s\n"+
				"    _Описание_: %s\n"+
				"    _Сторипоинты_: %d\n"+
				"    _Создана_: %s\n",
			task.Name,
			task.Description,
			task.StoryPoints,
			task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	return fmt.Sprintf(
		"%s\n"+
			"    _Описание_: %s\n"+
			"    _Сторипоинты_: %d\n"+
			"    _Создана_: %s\n"+
			"    _Дедлайн_: %s",
		task.Name,
		task.Description,
		task.StoryPoints,
		task.CreatedAt.Format("2006-01-02 15:04:05"),
		task.Deadline.Format("2006-01-02 15:04:05"))
}

func TasksDeadlineNotification(tasks []models.Task, chatId int64) *models.Message {
	var builder strings.Builder
	builder.WriteString("У тебя приблежается дедлайн по следующим задачам:\n\n")
	for i, task := range tasks {
		builder.WriteString(fmt.Sprintf("%v. ", i+1))
		builder.WriteString(taskView(&task))
		builder.WriteString("\n\n")
	}
	builder.WriteString("Поторопись!")

	return models.NewMessage(builder.String(), chatId)
}
