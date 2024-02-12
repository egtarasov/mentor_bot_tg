package convert

import (
	"telegrambot_new_emploee/internal/models"
	repoModels "telegrambot_new_emploee/internal/repository/models"
)

func ToTaskFromRepo(task *repoModels.Task) *models.Task {
	return &models.Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		StoryPoints: task.StoryPoints,
		Completed:   task.Completed,
		EmployeeId:  task.EmployeeId,
	}
}

func ToCommandFromRepo(cmd *repoModels.Command) *models.Command {
	return &models.Command{
		Id:       cmd.Id,
		Name:     cmd.Name,
		Action:   models.Action(cmd.Action),
		ParentId: cmd.ParentId,
	}
}

func ToTodoFromRepo(todo *repoModels.Todo) *models.Todo {
	return &models.Todo{
		Id:         todo.Id,
		Label:      todo.Label,
		Priority:   todo.Priority,
		EmployeeId: todo.EmployeeId,
		Completed:  todo.Completed,
	}
}

func ToUserFromRepo(user *repoModels.User) *models.User {
	return &models.User{
		Id:           user.Id,
		TelegramId:   user.TelegramId,
		Name:         user.Name,
		OccupationId: user.OccupationId,
	}
}

func ToMaterialFromRepo(material *repoModels.Material) *models.Material {
	return &models.Material{
		Id:        material.Id,
		Message:   material.Message,
		CommandId: material.CommandId,
	}
}

func ToGoalFromRepo(goal *repoModels.Goal) *models.Goal {
	return &models.Goal{
		Id:          goal.Id,
		Name:        goal.Name,
		Description: goal.Description,
		EmployeeId:  goal.EmployeeId,
		Track:       models.Track(goal.Track),
	}
}

func ToArray[From any, To any](arr []From, convert func(from *From) *To) []To {
	res := make([]To, 0, len(arr))
	for _, item := range arr {
		res = append(res, *convert(&item))
	}

	return res
}
