package repository

import (
	"context"
	"fmt"
	"telegrambot_new_emploee/internal/models"
)

var (
	ErrNoCommand  = fmt.Errorf("no command with this name")
	ErrNoMaterial = fmt.Errorf("no material for this command")
	ErrNoImage    = fmt.Errorf("no image for this command")
	ErrTxFail     = fmt.Errorf("transaction failed")
)

type CommandRepo interface {
	GetCommand(ctx context.Context, command string) (*models.Command, error)
	GetMaterials(ctx context.Context, cmdId int64) (*models.Material, error)
	GetCommands(ctx context.Context, parentId int64) ([]models.Command, error)
	GetImagePath(ctx context.Context, commandId int64) (*string, error)
	GetCommandsWithMaterials(ctx context.Context) ([]models.CommandWithMaterial, error)
	UpdateCommand(ctx context.Context, commandName string, material *models.Material) error
	AddCommand(ctx context.Context, command *models.Command, message string) error
}
