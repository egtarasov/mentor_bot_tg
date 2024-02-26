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
	query := `select 
    			id, name, surname, telegram_id, occupation_id, first_working_day, adaptation_end_at 
				from employees where telegram_id = $1`

	err := pgxscan.Get(ctx, a.pool, &user, query, tag)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, repository.ErrNoUser
	}
	if err != nil {
		return nil, err
	}

	return convert.ToUserFromRepo(&user), nil
}

func (a *userPostgres) GetUsersOnAdaptation(ctx context.Context) ([]models.User, error) {
	var users []repoModels.User
	query := `select 
    			id, name, surname, telegram_id, occupation_id, first_working_day, adaptation_end_at 
				from employees
				where adaptation_end_at > now()`

	err := pgxscan.Select(ctx, a.pool, &users, query)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(users, convert.ToUserFromRepo), nil
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
	if errors.As(err, &pgx.ErrNoRows) {
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
	if errors.As(err, &pgx.ErrNoRows) {
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
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(commands, convert.ToCommandFromRepo), err
}

func (c *commandPostgres) GetImagePath(ctx context.Context, commandId int64) (*string, error) {
	var path string
	query := `select path from pictures where command_id = $1`

	err := pgxscan.Get(ctx, c.pool, &path, query, commandId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, repository.ErrNoImage
	}
	if err != nil {
		return nil, err
	}

	return &path, nil
}

type tasksPostgres struct {
	pool *pgxpool.Pool
}

func NewTasksRepo(pool *pgxpool.Pool) repository.TasksRepo {
	return &tasksPostgres{
		pool: pool,
	}
}

func (t *tasksPostgres) GetTasksById(ctx context.Context, employeeId int64) ([]models.Task, error) {
	var tasks []repoModels.Task
	query := `select id, name, description, story_points, employee_id, created_at, completed_at 
				from tasks 
				where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &tasks, query, employeeId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(tasks, convert.ToTaskFromRepo), nil
}

func (t *tasksPostgres) GetTodoListById(ctx context.Context, employeeId int64) ([]models.Todo, error) {
	var todos []repoModels.Todo
	query := `select id, label, priority, employee_id, completed from todo_list where employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &todos, query, employeeId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(todos, convert.ToTodoFromRepo), nil
}

func (t *tasksPostgres) CheckTodo(ctx context.Context, todo *models.Todo) error {
	query := `update todo_list set completed = true where id = $1`
	_, err := t.pool.Exec(ctx, query, todo.Id)
	return err
}

func (t *tasksPostgres) GetGoalsById(ctx context.Context, employeeId int64) ([]models.Goal, error) {
	var goals []repoModels.Goal
	query := `select g.id, g.name, g.description, g.employee_id, t.track
	from goals g
    join goal_tracks t on g.track_id = t.id
    where g.employee_id = $1`

	err := pgxscan.Select(ctx, t.pool, &goals, query, employeeId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(goals, convert.ToGoalFromRepo), nil
}

func (t *tasksPostgres) CreateQuestion(ctx context.Context, question *models.Question) (int64, error) {
	query := `insert into questions (text, user_id) values ($1, $2)  returning id`
	var id int64
	err := t.pool.QueryRow(ctx, query, question.Text, question.UserId).Scan(&id)
	return id, err
}
