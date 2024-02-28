package admin

type AnswerQuestionRequest struct {
	QuestionId int64  `json:"question_id"`
	Answer     string `json:"answer"`
}

type UpdateMaterialRequest struct {
	CommandId int64  `json:"command_id"`
	Message   string `json:"message"`
}

type AddCommandRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
	Message  string `json:"message"`
	ActionId int    `json:"action_id"`
}
