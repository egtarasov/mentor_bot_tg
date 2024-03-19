package commands

import (
	"context"
	"errors"
	"telegrambot_new_emploee/internal/convert"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/models"
	"telegrambot_new_emploee/internal/repository"
)

type getMaterialCmd struct {
}

type subDirCmd struct {
}

func NewGetDataCmd() Cmd {
	return &getMaterialCmd{}
}

func NewSubDirCmd() Cmd {
	return &subDirCmd{}
}

func (c *getMaterialCmd) Execute(ctx context.Context, job *Job) error {
	material, err := container.Container.CmdRepo().GetMaterials(ctx, job.Command.Id)
	if err != nil {
		return err
	}
	photoPath, err := container.Container.CmdRepo().GetImagePath(ctx, job.Command.Id)
	if err != nil && !errors.Is(err, repository.ErrNoImage) {
		return err
	}

	msg := models.NewMessage(material.Message, job.GetChatId())
	msg.PhotoPath = photoPath

	return container.Container.Bot().SendMessage(ctx, msg)
}

func (c *subDirCmd) Execute(ctx context.Context, job *Job) error {
	commands, err := container.Container.CmdRepo().GetCommands(ctx, job.Command.Id)
	if err != nil {
		return err
	}

	material, err := container.Container.CmdRepo().GetMaterials(ctx, job.Command.Id)
	if err != nil {
		return err
	}

	photoPath, err := container.Container.CmdRepo().GetImagePath(ctx, job.Command.Id)
	if err != nil && !errors.Is(err, repository.ErrNoImage) {
		return err
	}

	buttons := convert.CommandsToButtons(commands, convert.DefaultToButtonsCfg(material.Message, job.GetChatId()))
	buttons.Message.PhotoPath = photoPath

	return container.Container.Bot().SendButtons(ctx, buttons)
}

var temp = `Часто сотрудники сталкиваются с различными проблемами. Чтобы тебе проще было разобраться, мы собрали самые частые проблемы в одном месте, смотри:

1. [Оформить болничный](https://best-company.ru/sick)

2. [О командировке](https://best-company.ru/bussines_trip)

3. [Подписи докумнетов](https://best-company.ru/docs)

4. [Об NDA](https://best-company.ru/nda)

Кроме того, как новый сотрудник, тебе предстоит много всего сделать в первые дни. Чтобы ты ни о чем не забыл, мы собрали для тебя все процедуры, которые тебе предстоит пройти:

1. Получить доступ к ресурсам.

2. Взять рабочую технику.

3. Познакомится с коммандой.

4. Получть электронную подпись (если оформлял при устрйостве на работу).

За поробностями обращайся к своему руководителю и HR-отделу!`
