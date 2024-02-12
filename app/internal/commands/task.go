package commands

import (
	"context"
	"sort"
	"strconv"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
)

type showTodoListCmd struct {
}

func NewShowTodoListCmd() Cmd {
	return &showTodoListCmd{}
}

func (c *showTodoListCmd) Execute(ctx context.Context, job *Job) error {
	// Get user's uncompleted todos.
	todos, err := getUncompletedTodo(ctx, job.User.Id)
	if err != nil {
		return err
	}
	msg := views.ShowTodo(todos)

	// Show the user their uncompleted todos.
	return container.Container.Bot().SendMessage(ctx, models.NewMessage(msg, job.Update.ChatId))
}

func getUncompletedTodo(ctx context.Context, userId int64) ([]models.Todo, error) {
	todos, err := container.Container.TaskRepo().GetTodoListById(ctx, userId)
	if err != nil {
		return nil, err
	}
	uncompletedTodos := make([]models.Todo, 0, len(todos))
	for _, todo := range todos {
		if todo.Completed {
			continue
		}
		uncompletedTodos = append(uncompletedTodos, todo)
	}

	sort.Slice(uncompletedTodos, func(i, j int) bool {
		return todos[i].Priority > todos[j].Priority
	})

	return uncompletedTodos, nil
}

type checkTodoCmd struct {
}

func NewCheckTodoCmd() Cmd {
	return &checkTodoCmd{}
}

func (c *checkTodoCmd) Execute(ctx context.Context, job *Job) error {
	// Get user's uncompleted todos.
	todos, err := getUncompletedTodo(ctx, job.User.Id)
	if err != nil {
		return err
	}

	// Ask the user for a button.
	msg := views.CheckTodo(todos)
	err = container.Container.Bot().SendMessage(ctx, models.NewMessage(msg, job.Update.ChatId))
	if err != nil {
		return err
	}

	number, err := getNumber(ctx, job, len(todos))
	if err != nil {
		return err
	}

	// Check the todo as completed.
	err = container.Container.TaskRepo().CheckTodo(ctx, &todos[number])
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(
		ctx,
		models.NewMessage("Задача была отмечена выполненой!", job.Update.ChatId))
}

// getNumber waits for a number from 1 to limit from the user. If the limit less or equal zero, there is no upper bound.
// Assuming that the message with the request for a number has been already sent. If 'Отмена' will be sent, cancel the
// processing with the corresponding error.
func getNumber(ctx context.Context, job *Job, limit int) (number int, err error) {
	for {
		update := job.Queue.WaitForUpdate()
		if update.Message == CancelMessage {
			return 0, ErrCanceled
		}
		number, err = strconv.Atoi(update.Message)
		if err != nil {
			err = container.Container.Bot().SendMessage(
				ctx,
				models.NewMessage(
					"Это не число, попробуйте снова!",
					job.Update.ChatId))
			if err != nil {
				return 0, err
			}
			continue
		}

		if (limit > 0 && number >= limit) || number < 0 {
			err = container.Container.Bot().SendMessage(
				ctx,
				models.NewMessage(
					"Введенное число превышает допустимое значение, попробуйет снова!",
					job.Update.ChatId))
			if err != nil {
				return 0, err
			}
			continue
		}
		break
	}

	return number, err
}

type showGoalsCmd struct {
}

func NewShowGoalCmd() Cmd {
	return &showGoalsCmd{}
}

func (c *showGoalsCmd) Execute(ctx context.Context, job *Job) error {
	goals, err := container.Container.TaskRepo().GetGoalsById(ctx, job.User.Id)
	if err != nil {
		return err
	}

	msg := views.GetGoals(goals)

	return container.Container.Bot().SendMessage(ctx, models.NewMessage(msg, job.Update.ChatId))
}
