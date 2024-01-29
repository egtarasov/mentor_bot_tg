package api

import (
	"context"
	"errors"
	"log"
	"sync"
	"telegrambot_new_emploee/internal/bot"
	"telegrambot_new_emploee/internal/commands"
	"telegrambot_new_emploee/internal/config"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/updates"
)

type app struct {
	ctx       context.Context
	wgUpdates sync.WaitGroup

	// Store the queues of updates for each user. The lock is used to safely delete and add queues.
	usersLock sync.Mutex
	users     map[int64]*updates.Queue

	container *container.DiContainer

	// Commands available for a bot.
	getCmd     commands.Cmd
	subDirCmd  commands.Cmd
	complexCmd map[string]commands.Cmd
}

func newApp() (*app, error) {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	ct, err := container.NewDiContainer(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &app{
		ctx:       ctx,
		wgUpdates: sync.WaitGroup{},

		usersLock: sync.Mutex{},
		users:     make(map[int64]*updates.Queue),

		container: ct,

		getCmd:     nil,
		subDirCmd:  nil,
		complexCmd: nil,
	}, nil
}

func addCommands(app *app) {
	app.getCmd = commands.NewGetDataCmd(app.container)
	app.subDirCmd = commands.NewSubDirCmd(app.container)

	// All complex commands must be registered there. Note, that the name in database and name in map must be the same,
	// otherwise the command will node function.
	app.complexCmd = make(map[string]commands.Cmd)
	complexCmd := []struct {
		key string
		cmd commands.Cmd
	}{
		{key: "Показать чек-лист", cmd: commands.NewShowTodoListCmd(app.container)},
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

	app.run()
}

func (a *app) run() {
	updatesCh := a.container.Bot().Start(a.ctx)

	// Process updates from the bot.
	for update := range updatesCh {
		a.processUpdate(update)
	}

	a.wgUpdates.Wait()
}

// processUpdate process the incoming update. Each update is placed into the updateQueue for the appropriate user, where
// it will be processed later on. If the queue did not exist (the user did not communicate recently with the bot), a new
// queue will be created and the processing goroutine for this queue will start.
func (a *app) processUpdate(update *models.Update) {
	// Authenticate the user.
	user := a.authenticate(update)
	if user == nil {
		// TODO return the message that user is unknown.
		return
	}

	// Get the current state of the user.
	a.usersLock.Lock()
	defer a.usersLock.Unlock()
	queue, ok := a.users[user.UserId]
	if !ok {
		// Create a new update's queue and start the processing goroutine.
		queue := updates.newQueue(user)
		// There is no need for synchronization because only this goroutine has access to the queue.
		queue.AddUpdate(update)
		a.users[user.UserId] = queue
		go a.processQueue(queue)
		return
	}

	// Just add a new update to the queue and signal if any goroutine is waiting for a new update.
	queue.AddUpdate(update)
}

// processQueue process the queue of updates for a given user. If the queue becomes empty,
// it will be removed and the processing ceased.
func (a *app) processQueue(queue *updates.updatesQueue) {
	for {
		update := a.getUpdate(queue)
		if update == nil {
			return
		}

		job, ok := a.newJob(queue, update)
		// TODO a better response to the unknown command.
		if !ok {
			_ = a.bot.SendMessage(a.ctx, bot.Message{
				Message: "Unknown command",
				ChatId:  update.ChatId,
			})
			continue
		}

		a.wgUpdates.Add(1)
		a.processJob(job)
	}
}

// getUpdate gets an update concurrently from the queue. If there is no updates to obtain the queue will be removed
// from the application's users map.
func (a *app) getUpdate(queue *updates.updatesQueue) *Update {
	// The lock is to safely delete a queue from the map (if a new update will come, and we will try to delete a map).
	a.usersLock.Lock()
	queue.lock.Lock()
	defer a.usersLock.Unlock()
	defer queue.lock.Unlock()

	// Remove the queue if there is no more updates to process.
	if len(queue.updates) == 0 {
		delete(a.users, queue.user.UserId)
		return nil
	}

	update := queue.updates[0]
	queue.updates = queue.updates[1:]

	return update
}

// newJob create a new job from the update.
func (a *app) newJob(queue *updates.updatesQueue, update *Update) (*job, bool) {
	job := &job{
		command: nil,
		update:  update,
		queue:   queue,
	}

	if ok := a.getCommand(job); !ok {
		return nil, false
	}

	return job, true
}

// processJob process a job based on the job's command.
func (a *app) processJob(job *job) {
	defer a.wgUpdates.Done()

	var cmd commands.Cmd

	switch job.command.Action {
	case GetDataCmd:
		cmd = a.getCmd
	case GetSubsectionsCmd:
		cmd = a.subDirCmd
	case ComplexCmd:
		cmd = a.complexCmd[job.command.Name]
	default:
		// TODO log unknown command type
		return
	}

	err := cmd.Execute(a.ctx, job)
	if err != nil {
		// TODO logging
	}
}

// getCommand gets a command for a given job. Returns true if command is determined, false otherwise.
func (a *app) getCommand(job *job) bool {
	command, err := a.container.CmdRepo().GetCommand(a.ctx, job.update.Message)
	if err != nil {
		// TODO logging
		return false
	}

	job.command = ToCommand(command)

	return true
}

// authenticate gets a user for a given update.
func (a *app) authenticate(update *models.Update) *models.User {
	user, err := a.container.UserRepo().GetUserByTag(a.ctx, update.UpdateUserId)
	if errors.Is(err, repository.ErrNoUser) {
		return nil
	}
	if err != nil {
		// TODO processJob an unexpected error
		return nil
	}
	return ToUser(user)
}