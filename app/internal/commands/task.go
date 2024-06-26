package commands

import (
	"context"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/views"
)

type showTodoListCmd struct {
}

func NewShowTodoListCmd() Command {
	return &showTodoListCmd{}
}

func (c *showTodoListCmd) Execute(ctx context.Context, job *Job) error {
	// Get user's uncompleted todos.
	todos, err := container.Container.TaskRepo().GetTodoListById(ctx, job.User.Id)
	if err != nil {
		return err
	}

	// Show the user their uncompleted todos.
	return container.Container.Bot().
		SendMessage(
			ctx,
			views.GetTodo(
				convert.UncompletedTodo(todos),
				job.User,
				len(todos),
			),
		)
}

type checkTodoCmd struct {
}

func NewCheckTodoCmd() Command {
	return &checkTodoCmd{}
}

func (c *checkTodoCmd) Execute(ctx context.Context, job *Job) error {
	// Get user's uncompleted todos.
	todos, err := container.Container.TaskRepo().GetTodoListById(ctx, job.User.Id)
	if err != nil {
		return err
	}
	uncompletedTodos := convert.UncompletedTodo(todos)
	if len(uncompletedTodos) == 0 {
		return container.Container.Bot().SendMessage(ctx, models.NewMessage("Все задачи выполнены!", job.GetChatId()))
	}
	// Ask the user for a button.
	err = container.Container.Bot().SendMessage(ctx, views.CheckTodo(uncompletedTodos, job.GetChatId()))
	if err != nil {
		return err
	}

	number, err := getNumber(ctx, job, len(uncompletedTodos))
	if err != nil {
		return err
	}

	// Check the todo as completed.
	err = container.Container.TaskRepo().CheckTodo(ctx, &uncompletedTodos[number-1])
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(
		ctx,
		models.NewMessage("Задача была отмечена выполненой!", job.GetChatId()))
}

type showGoalsCmd struct {
}

func NewShowGoalCmd() Command {
	return &showGoalsCmd{}
}

func (c *showGoalsCmd) Execute(ctx context.Context, job *Job) error {
	goals, err := container.Container.TaskRepo().GetGoalsByUser(ctx, job.User)
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, views.GetGoals(goals, job.GetChatId()))
}

type showTasksCmd struct {
}

func NewShowTasksCmd() Command {
	return &showTasksCmd{}
}

func (c *showTasksCmd) Execute(ctx context.Context, job *Job) error {
	tasks, err := container.Container.TaskRepo().GetTasksById(ctx, job.User.Id)
	if err != nil {
		return err
	}
	return container.Container.Bot().SendMessage(ctx, views.GetTasks(tasks, job.User))
}

func uncompletedTask(ctx context.Context, tasks []models.Task) []models.Task {
	uncompletedTasks := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		if task.CompletedAt != nil {
			continue
		}
		uncompletedTasks = append(uncompletedTasks, task)
	}
	return uncompletedTasks
}

type occupationMaterialCmd struct {
}

func NewOccupationMaterialCmd() Command {
	return &occupationMaterialCmd{}
}

func (c *occupationMaterialCmd) Execute(ctx context.Context, job *Job) error {
	occupation, err := container.Container.TaskRepo().GetOccupationMaterial(ctx, job.User.OccupationId)
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage(occupation.Material, job.GetChatId()))
}

type checkTaskCmd struct {
}

func NewCheckTask() Command {
	return &checkTaskCmd{}
}

func (c *checkTaskCmd) Execute(ctx context.Context, job *Job) error {
	// Get user's uncompleted todos.
	tasks, err := container.Container.TaskRepo().GetTasksById(ctx, job.User.Id)
	if err != nil {
		return err
	}
	uncompletedTasks := uncompletedTask(ctx, tasks)
	if len(uncompletedTasks) == 0 {
		return container.Container.Bot().SendMessage(ctx, models.NewMessage("Все задачи выполнены!", job.GetChatId()))
	}
	// Ask the user for a button.
	err = container.Container.Bot().SendMessage(ctx, views.CheckTask(uncompletedTasks, job.GetChatId()))
	if err != nil {
		return err
	}

	number, err := getNumber(ctx, job, len(uncompletedTasks))
	if err != nil {
		return err
	}

	// Check the todo as completed.
	err = container.Container.TaskRepo().CheckTask(ctx, &uncompletedTasks[number-1])
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(
		ctx,
		models.NewMessage("Задача была отмечена выполненой!", job.GetChatId()))
}
