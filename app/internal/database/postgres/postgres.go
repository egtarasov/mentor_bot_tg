package postgres

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"telegrambot_new_emploee/internal/database"
)

type userPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) database.UserRepo {
	return &userPostgres{
		pool: pool,
	}
}

func (a *userPostgres) GetUserByTag(ctx context.Context, tag string) (*database.User, error) {
	var user database.User
	query := "select id, name, surname, telegram_tag, occupation_id from employees where telegram_tag = $1"

	err := pgxscan.Get(ctx, a.pool, &user, query, tag)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, database.ErrNoUser
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type commandPostgres struct {
	pool *pgxpool.Pool
}

func NewCommandPostgres(pool *pgxpool.Pool) database.CommandRepo {
	return &commandPostgres{
		pool: pool,
	}
}

func (c *commandPostgres) GetCommand(ctx context.Context, command string) (*database.Command, error) {
	var cmd database.Command
	query := `select c.id, c.name, a.action, c.parent_id
			  from public.commands c 
			  join public.actions a 
				on c.action_id = a.id
				where c.name = $1`

	err := pgxscan.Get(ctx, c.pool, &cmd, query, command)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, database.ErrNoCommand
	}
	if err != nil {
		return nil, err
	}

	return &cmd, nil
}

func (c *commandPostgres) GetMaterials(ctx context.Context, cmdId int64) (*database.Material, error) {
	var material database.Material
	query := `select * from materials where command_id = $1`

	err := pgxscan.Get(ctx, c.pool, &material, query, cmdId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, database.ErrNoMaterial
	}
	if err != nil {
		return nil, err
	}

	return &material, nil
}

func (c *commandPostgres) GetCommands(ctx context.Context, parentId int64) ([]database.Command, error) {
	commands := make([]database.Command, 0)
	query := `select c.id, c.name, a.action, c.parent_id 
				from public.commands c
				join public.actions a on c.action_id = a.id 
				where parent_id = $1`
	err := pgxscan.Select(ctx, c.pool, &commands, query, parentId)
	if err != nil {
		return nil, err
	}

	return commands, err
}

type tasksPostgres struct {
	pool *pgxpool.Pool
}

func NewTasksRepo(pool *pgxpool.Pool) database.TasksRepo {
	return &tasksPostgres{
		pool: pool,
	}
}

func (t *tasksPostgres) GetTasks(ctx context.Context, employeeId int64) []database.Task {
	var tasks []database.Task
	query := `select id, name, description, story_points, completed, employee_id from tasks where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &tasks, query, employeeId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return nil
	}

	return tasks
}

func (t *tasksPostgres) GetTodoList(ctx context.Context, employeeId int64) []database.Todo {
	var tasks []database.Todo
	query := `select id, label, priority, employee_id, completed from todo_list where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &tasks, query, employeeId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	if err != nil {
		return nil
	}

	return tasks
}
