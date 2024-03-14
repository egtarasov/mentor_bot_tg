package api

import (
	"context"
	"errors"
	"log"
	"sync"
	"telegrambot_new_emploee/internal/admin"
	"telegrambot_new_emploee/internal/commands"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/daemons"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/updates"
	updatesqueue "telegrambot_new_emploee/internal/updates/updates-queue"
)

type app struct {
	ctx       context.Context
	wgUpdates sync.WaitGroup

	// Store the queues of updates for each user. The lock is used to safely delete and add queues.
	users updates.Map

	// Commands available for a bot.
	getCmd     commands.Cmd
	subDirCmd  commands.Cmd
	complexCmd map[string]commands.Cmd
}

func newApp() (*app, error) {
	ctx := context.Background()

	err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	err = container.NewDiContainer(ctx)
	if err != nil {
		return nil, err
	}

	return &app{
		ctx:       ctx,
		wgUpdates: sync.WaitGroup{},

		users: updatesqueue.NewMap(),

		getCmd:     nil,
		subDirCmd:  nil,
		complexCmd: nil,
	}, nil
}

func addCommands(app *app) {
	app.getCmd = commands.NewGetDataCmd()
	app.subDirCmd = commands.NewSubDirCmd()

	// All complex commands must be registered there. Note, that the name in database and name in map must be the same,
	// otherwise the command will node function.
	app.complexCmd = make(map[string]commands.Cmd)
	complexCmd := []struct {
		key string
		cmd commands.Cmd
	}{
		{key: "Показать чек-лист", cmd: commands.NewShowTodoListCmd()},
		{key: "Отметить задачу в чек-листе", cmd: commands.NewCheckTodoCmd()},
		{key: "Цели", cmd: commands.NewShowGoalCmd()},
		{key: "Показать задачи", cmd: commands.NewShowTasksCmd()},
		{key: "Задать вопрос", cmd: commands.NewAskQuestionCmd()},
		{key: "Показать задачт", cmd: commands.NewShowTasksCmd()},
		{key: "Полезные матерьялы для меня", cmd: commands.NewOccupationMaterialCmd()},
	}
	if config.Cfg.CalendarUrl != nil {
		app.complexCmd["Календарь"] = commands.NewCalendarCmd()
	}

	for _, node := range complexCmd {
		app.complexCmd[node.key] = node.cmd
	}
}

func Run() {
	app, err := newApp()
	if err != nil {
		log.Fatal(err)
	}
	addCommands(app)

	app.runDaemons()
	app.runServer()

	app.run()
}

func (a *app) runServer() {
	go admin.StartServer()
}

func (a *app) runDaemons() {
	go daemons.NewFeedbackDaemon(a.ctx)
	go daemons.NewHrMeetupDaemon(a.ctx)
	go daemons.NewTrainingDaemon(a.ctx)
	go daemons.NewMentorMeetupDaemon(a.ctx)
}

func (a *app) run() {
	updatesCh := container.Container.Bot().Start(a.ctx)

	// Process updates from the bot.
	for update := range updatesCh {
		a.processUpdate(update)
	}

	a.wgUpdates.Wait()
}

func (a *app) changeUpdate(update *models.Update) {
	commandsSubstitutions := map[string]string{
		"В меню":    "/start",
		"/ask":      AskQuestionCmd,
		"/calendar": CalendarCmd,
	}
	if val, ok := commandsSubstitutions[update.Message]; ok {
		update.Message = val
	}
}

// processUpdate process the incoming update. Each update is placed into the updateQueue for the appropriate user, where
// it will be processed later on. If the queue did not exist (the user did not communicate recently with the bot), a new
// queue will be created and the processing goroutine for this queue will start.
func (a *app) processUpdate(update *models.Update) {
	// Authenticate the user.
	user := a.authenticate(update)
	if user == nil {
		_ = container.Container.Bot().
			SendMessage(a.ctx,
				models.NewMessage("Ой, ой! Кажется, вы не являетесь сотрудником!", update.ChatId))
		return
	}

	a.changeUpdate(update)

	// Get or create a queue for a user and put update into it.
	queue, ok := a.users.GetOrCreate(user.Id)
	queue.AddUpdate(update)
	// If the queue did not exist, we need to start processing goroutine.
	if !ok {
		go a.processQueue(queue, user)
	}
}

func (a *app) processQueue(queue updates.Queue, user *models.User) {
	for {
		update := a.users.GetUpdate(user.Id, queue)
		if update == nil {
			return
		}

		job, ok := commands.NewJob(a.ctx, queue, update, user)
		// TODO a better response to the unknown command.
		if !ok {
			_ = container.Container.Bot().SendMessage(
				a.ctx, models.NewMessage(
					"Такой комманды нету :(", update.ChatId))
			continue
		}

		a.wgUpdates.Add(1)
		a.processJob(job)
	}
}

// processJob process a job based on the job's command.
func (a *app) processJob(job *commands.Job) {
	defer a.wgUpdates.Done()

	var cmd commands.Cmd

	switch models.IntToAction(job.Command.ActionId) {
	case models.GetDataCmd:
		cmd = a.getCmd
	case models.GetSubsectionsCmd:
		cmd = a.subDirCmd
	case models.ComplexCmd:
		cmd = a.complexCmd[job.Command.Name]
	default:
		// TODO log unknown command type
		return
	}

	err := cmd.Execute(a.ctx, job)
	if errors.Is(err, commands.ErrCanceled) {
		_ = container.Container.Bot().SendMessage(a.ctx, models.NewMessage(CancelMessage, job.Update.ChatId))
	} else if err != nil {
		_ = container.Container.Bot().SendMessage(a.ctx, models.NewMessage("Простите, что-то пошло не так :(",
			job.Update.ChatId))
		log.Println(err)
	}
}

// authenticate gets a user for a given update.
func (a *app) authenticate(update *models.Update) *models.User {
	user, err := container.Container.UserRepo().GetUserByTag(a.ctx, update.UpdateUserId)
	if errors.Is(err, repository.ErrNoUser) {
		return nil
	}
	if err != nil {
		// TODO processJob an unexpected error
		return nil
	}

	update.User = user
	return user
}
