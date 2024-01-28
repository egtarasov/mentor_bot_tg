package database

import (
	"context"
	"fmt"
)

var (
	ErrNoCommand  = fmt.Errorf("no command with this name")
	ErrNoMaterial = fmt.Errorf("no material for this command")
)

type Command struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Action   string `db:"action"`
	ParentId int64  `db:"parent_id"`
}

type Material struct {
	Id        int64  `db:"id"`
	Message   string `db:"message"`
	CommandId int64  `db:"command_id"`
}

type CommandRepo interface {
	GetCommand(ctx context.Context, command string) (*Command, error)
	GetMaterials(ctx context.Context, cmdId int64) (*Material, error)
	GetCommands(ctx context.Context, parentId int64) ([]Command, error)
}
