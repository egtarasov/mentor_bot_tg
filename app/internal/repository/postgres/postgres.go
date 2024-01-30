package postgres

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/repository/convert"
	repoModels "telegrambot_new_emploee/internal/repository/models"
)

type userPostgres struct {
	pool *pgxpool.Pool
}

func NewUserPostgres(pool *pgxpool.Pool) repository.UserRepo {
	return &userPostgres{
		pool: pool,
	}
}

func (a *userPostgres) GetUserByTag(ctx context.Context, tag int64) (*models.User, error) {
	var user repoModels.User
	query := "select id, name, surname, telegram_id, occupation_id from employees where telegram_id = $1"

	err := pgxscan.Get(ctx, a.pool, &user, query, tag)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrNoUser
	}
	if err != nil {
		return nil, err
	}

	return convert.ToUserFromRepo(&user), nil
}

type commandPostgres struct {
	pool *pgxpool.Pool
}

func NewCommandPostgres(pool *pgxpool.Pool) repository.CommandRepo {
	return &commandPostgres{
		pool: pool,
	}
}

func (c *commandPostgres) GetCommand(ctx context.Context, command string) (*models.Command, error) {
	var cmd repoModels.Command
	query := `select c.id, c.name, a.action, c.parent_id
			  from public.commands c 
			  join public.actions a 
				on c.action_id = a.id
				where c.name = $1`

	err := pgxscan.Get(ctx, c.pool, &cmd, query, command)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrNoCommand
	}
	if err != nil {
		return nil, err
	}

	return convert.ToCommandFromRepo(&cmd), nil
}

func (c *commandPostgres) GetMaterials(ctx context.Context, cmdId int64) (*models.Material, error) {
	var material repoModels.Material
	query := `select * from materials where command_id = $1`

	err := pgxscan.Get(ctx, c.pool, &material, query, cmdId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrNoMaterial
	}
	if err != nil {
		return nil, err
	}

	return convert.ToMaterialFromRepo(&material), nil
}

func (c *commandPostgres) GetCommands(ctx context.Context, parentId int64) ([]models.Command, error) {
	commands := make([]repoModels.Command, 0)
	query := `select c.id, c.name, a.action, c.parent_id 
				from public.commands c
				join public.actions a on c.action_id = a.id 
				where parent_id = $1`
	err := pgxscan.Select(ctx, c.pool, &commands, query, parentId)
	if err != nil {
		return nil, err
	}

	return convert.ToArray(commands, convert.ToCommandFromRepo), err
}

type tasksPostgres struct {
	pool *pgxpool.Pool
}

func NewTasksRepo(pool *pgxpool.Pool) repository.TasksRepo {
	return &tasksPostgres{
		pool: pool,
	}
}

func (t *tasksPostgres) GetTasks(ctx context.Context, employeeId int64) []models.Task {
	var tasks []repoModels.Task
	query := `select id, name, description, story_points, completed, employee_id from tasks where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &tasks, query, employeeId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return nil
	}

	return convert.ToArray(tasks, convert.ToTaskFromRepo)
}

func (t *tasksPostgres) GetTodoList(ctx context.Context, employeeId int64) []models.Todo {
	var todos []repoModels.Todo
	query := `select id, label, priority, employee_id, completed from todo_list where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &todos, query, employeeId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return nil
	}

	return convert.ToArray(todos, convert.ToTodoFromRepo)
}
