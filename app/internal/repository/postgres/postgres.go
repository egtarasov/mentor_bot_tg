package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"telegrambot_new_emploee/internal/config"
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

func (a *userPostgres) AddTasks(ctx context.Context, tasks *repository.AddTasks) error {
	tx, err := a.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, task := range tasks.Tasks {
		_, err := tx.Exec(ctx, `insert into tasks (name, description, story_points, employee_id) values ($1, $2, $3, $4)`,
			task.Name, task.Description, task.StoryPoints, tasks.EmployeeId)
		if err != nil {
			return err
		}
	}

	for _, todo := range tasks.Todos {
		_, err := tx.Exec(ctx, `insert into todo_list (label, employee_id) values ($1, $2)`, todo, tasks.EmployeeId)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (a *userPostgres) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	var user repoModels.User
	query := `select 
    			id, name, surname, telegram_id, occupation_id, first_working_day, adaptation_end_at, is_admin, grade
				from employees where id = $1`

	err := pgxscan.Get(ctx, a.pool, &user, query, userId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, repository.ErrNoUser
	}
	if err != nil {
		return nil, err
	}

	return convert.ToUserFromRepo(&user), nil
}

func (a *userPostgres) GetUserByTag(ctx context.Context, tag int64) (*models.User, error) {
	var user repoModels.User
	query := `select 
    			id, name, surname, telegram_id, occupation_id, first_working_day, adaptation_end_at, is_admin, grade 
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
	query := `select c.id, c.name, c.action_id, c.parent_id, c.is_admin
			  from public.commands c 
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
	query := `select c.id, c.name, c.action_id, c.parent_id 
				from public.commands c
				where parent_id = $1
				order by c.id`
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

func (c *commandPostgres) GetCommandsWithMaterials(ctx context.Context) ([]models.CommandWithMaterial, error) {
	var commands []repoModels.CommandWithMaterial
	query := `select c.id, c.name, m.message, c.action_id
			  from commands c
			  right join public.materials m on c.id = m.command_id`

	err := pgxscan.Select(ctx, c.pool, &commands, query)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(commands, convert.ToCommandWithMaterialFromRepo), nil
}

func (c *commandPostgres) UpdateCommand(ctx context.Context, commandName string, material *models.Material) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}

	// Update the command.
	query := `update commands
			  set name = case when $1 = '' then name else $1 end
			  where id = $2`
	res, err := tx.Exec(ctx, query, commandName, material.CommandId)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return repository.ErrNoCommand
	}

	// Update the material.
	query = `update materials
			  set message = case when $1 = '' then message else $1 end
			  where command_id= $2;`

	res, err = tx.Exec(ctx, query, material.Message, material.CommandId)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return repository.ErrNoMaterial
	}

	return tx.Commit(ctx)
}

func (c *commandPostgres) AddCommand(ctx context.Context, command *models.Command, message string) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}

	// Insert the command and get its id.
	query := `insert into commands (name, action_id, parent_id)
			  values ($1, $2, $3) returning id`

	var id int64
	err = tx.QueryRow(ctx, query, command.Name, command.ActionId, command.ParentId).Scan(&id)
	if errors.As(err, &pgx.ErrTxInFailure) {
		return repository.ErrTxFail
	}
	if err != nil {
		return err
	}

	// Insert the material.
	query = `insert into materials (message, command_id) 
			 values ($1, $2)`

	_, err = tx.Exec(ctx, query, message, id)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
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
	query := `select id, name, description, story_points, employee_id, created_at, completed_at, deadline 
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
	query := `select id, label, priority, employee_id, completed from todo_list
        	  where employee_id = $1
        	  order by priority`

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

func (t *tasksPostgres) GetGoalsByUser(ctx context.Context, user *models.User) ([]models.Goal, error) {
	var goals []repoModels.Goal
	query := `select g.id, g.name, g.description, g.grade, gt.track, o.name as occupation_name
			  from goals g
			  join public.goal_tracks gt on gt.id = g.track_id
			  join occupations o on o.id = g.occupation_id
			  where g.grade = $1 and g.occupation_id = $2`

	err := pgxscan.Select(ctx, t.pool, &goals, query, user.Grade, user.OccupationId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(goals, convert.ToGoalFromRepo), nil
}

func (t *tasksPostgres) GetOccupationMaterial(ctx context.Context, occupationId int64) (*models.Occupation, error) {
	var occupation repoModels.Occupation
	query := `select id, name, material from occupations where id = $1`

	err := pgxscan.Get(ctx, t.pool, &occupation, query, occupationId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, repository.ErrNoData
	}
	if err != nil {
		return nil, err
	}

	return convert.ToOccupationFromRepo(&occupation), nil
}

type questionsPostgres struct {
	pool *pgxpool.Pool
}

func NewQuestionRepo(pool *pgxpool.Pool) repository.QuestionRepo {
	return &questionsPostgres{pool: pool}
}

func (q *questionsPostgres) CreateQuestion(ctx context.Context, question *models.Question) (int64, error) {
	query := `insert into questions (text, user_id) values ($1, $2)  returning id`
	var id int64
	err := q.pool.QueryRow(ctx, query, question.Text, question.UserId).Scan(&id)
	return id, err
}

func (q *questionsPostgres) GetUnansweredQuestions(ctx context.Context) ([]models.Question, error) {
	var questions []repoModels.Question
	query := `select id, text, user_id, created_at, answered_at, answered_by, answer
			  from questions where answered_at is null
			  order by created_at`

	err := pgxscan.Select(ctx, q.pool, &questions, query)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return convert.ToArray(questions, convert.ToQuestionFromRepo), nil
}

func (q *questionsPostgres) AnswerQuestion(ctx context.Context, question *models.Question) error {
	tx, err := q.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `update questions
			  set answered_at = now(), answer = $2, answered_by = $3
			  where id = $1 and answered_at is null`
	res, err := tx.Exec(ctx, query, question.Id, question.Answer, question.AnsweredBy)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return repository.ErrQuestionNotExist
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (q *questionsPostgres) GetQuestionById(ctx context.Context, questionId int64) (*models.Question, error) {
	var question repoModels.Question
	query := `select id, text, user_id, created_at, answered_at, answered_by, answer
			  from questions where id = $1`

	err := pgxscan.Get(ctx, q.pool, &question, query, questionId)
	if errors.As(err, &pgx.ErrNoRows) {
		return nil, repository.ErrQuestionNotExist
	}
	if err != nil {
		return nil, err
	}

	return convert.ToQuestionFromRepo(&question), nil
}

type adminPostgres struct {
	pool        *pgxpool.Pool
	commandRepo repository.CommandRepo
}

func NewAdminPostgres(pool *pgxpool.Pool, repo repository.CommandRepo) repository.AdminRepository {
	return &adminPostgres{pool: pool, commandRepo: repo}
}

func (f *adminPostgres) GetFAQSections(ctx context.Context) ([]string, error) {
	var sections []string
	query := `select name from commands
			  where parent_id = (select id from commands where name = $1)`

	err := pgxscan.Select(ctx, f.pool, &sections, query, config.Cfg.Admin.FAQCommandName)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (f *adminPostgres) UpdateFAQ(ctx context.Context, faq *repository.UpdateFaq) error {
	command, err := f.commandRepo.GetCommand(ctx, faq.SectionName)
	if err != nil {
		return err
	}
	tx, err := f.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var message string
	err = pgxscan.Get(ctx, tx, &message, `select message from materials where command_id = $1`, command.Id)
	if err != nil {
		return err
	}

	message += fmt.Sprintf("\n\nQ: *%s*\nA: %s", faq.Question, faq.Answer)
	_, err = tx.Exec(ctx, `update materials set message = $1 where command_id = $2`, message, command.Id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (f *adminPostgres) GetNotifications(ctx context.Context) ([]models.Notification, error) {
	var notifications []repoModels.Notification
	query := `select id, message, photo_path, notification_time, day_of_week, repeat_time from notifications`
	err := pgxscan.Select(ctx, f.pool, &notifications, query)
	if err != nil {
		return nil, err
	}

	return convert.ToArray(notifications, convert.ToNotificationFromRepo), nil
}
