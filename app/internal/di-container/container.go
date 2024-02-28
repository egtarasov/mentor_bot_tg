package container

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"telegrambot_new_emploee/internal/bot"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/repository/postgres"
)

var Container *DiContainer

type DiContainer struct {
	userRepo     repository.UserRepo
	cmdRepo      repository.CommandRepo
	taskRepo     repository.TasksRepo
	questionRepo repository.QuestionRepo
	bot          bot.Bot

	pool *pgxpool.Pool
}

func NewDiContainer(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, config.Cfg.ConnStr)
	if err != nil {
		return err
	}

	tgBot, err := bot.NewTelegramBot(config.Cfg.TgToken, "markdown")
	if err != nil {
		return err
	}

	Container = &DiContainer{
		bot:  tgBot,
		pool: pool,
	}

	return nil
}

func (c *DiContainer) Bot() bot.Bot {
	return c.bot
}

func (c *DiContainer) UserRepo() repository.UserRepo {
	if c.userRepo == nil {
		c.userRepo = postgres.NewUserPostgres(c.pool)
	}

	return c.userRepo
}

func (c *DiContainer) CmdRepo() repository.CommandRepo {
	if c.cmdRepo == nil {
		c.cmdRepo = postgres.NewCommandPostgres(c.pool)
	}

	return c.cmdRepo
}

func (c *DiContainer) TaskRepo() repository.TasksRepo {
	if c.taskRepo == nil {
		c.taskRepo = postgres.NewTasksRepo(c.pool)
	}

	return c.taskRepo
}

func (c *DiContainer) QuestionRepo() repository.QuestionRepo {
	if c.questionRepo == nil {
		c.questionRepo = postgres.NewQuestionRepo(c.pool)
	}

	return c.questionRepo
}
