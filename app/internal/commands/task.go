package commands

import (
	"context"
	"sort"
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

	// Show the user their uncompleted todos.
	return container.Container.Bot().SendMessage(ctx, views.ShowTodo(todos, job.GetChatId()))
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
	err = container.Container.Bot().SendMessage(ctx, views.CheckTodo(todos, job.GetChatId()))
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
		models.NewMessage("Задача была отмечена выполненой!", job.GetChatId()))
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

	return container.Container.Bot().SendMessage(ctx, views.GetGoals(goals, job.GetChatId()))
}

type showTasksCmd struct {
}

func NewShowTasksCmd() Cmd {
	return &showTasksCmd{}
}

func (c *showTasksCmd) Execute(ctx context.Context, job *Job) error {
	tasks, err := getUncompletedTask(ctx, job.User.Id)
	if err != nil {
		return err
	}
	return container.Container.Bot().SendMessage(ctx, views.GetTasks(tasks, job.GetChatId()))
}

func getUncompletedTask(ctx context.Context, userId int64) ([]models.Task, error) {
	tasks, err := container.Container.TaskRepo().GetTasksById(ctx, userId)
	if err != nil {
		return nil, err
	}
	uncompletedTasks := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		if task.CompletedAt != nil {
			continue
		}
		uncompletedTasks = append(uncompletedTasks, task)
	}
	return uncompletedTasks, nil
}

type occupationMaterialCmd struct {
}

func NewOccupationMaterialCmd() Cmd {
	return &occupationMaterialCmd{}
}

func (c *occupationMaterialCmd) Execute(ctx context.Context, job *Job) error {
	occupation, err := container.Container.TaskRepo().GetOccupationMaterial(ctx, job.User.OccupationId)
	if err != nil {
		return err
	}

	return container.Container.Bot().SendMessage(ctx, models.NewMessage(occupation.Material, job.GetChatId()))
}
