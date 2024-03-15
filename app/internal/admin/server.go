package admin

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"telegrambot_new_emploee/internal/config"
	container "telegrambot_new_emploee/internal/di-container"
	"telegrambot_new_emploee/internal/repository"
	"telegrambot_new_emploee/internal/services"
)

var TempId int64 = 1

func StartServer() {
	mux := http.NewServeMux()
	s := server{
		ctx:             context.Background(),
		questionService: services.NewQuestionService(),
		commandsService: services.NewCommandsService(),
	}

	mux.HandleFunc("GET /questions", s.GetRequestsHandler)
	mux.HandleFunc("POST /questions", s.AnswerQuestion)

	mux.HandleFunc("GET /commands", s.GetCommands)
	mux.HandleFunc("PUT /commands", s.ChangeCommand)
	mux.HandleFunc("POST /commands", s.AddCommand)

	mux.HandleFunc("POST /send", s.SendMessage)
	mux.HandleFunc("POST /add/tasks", s.AddEmployee)

	if err := http.ListenAndServe(config.Cfg.Admin.Port, mux); err != nil {
		log.Println(err)
	}
}

type server struct {
	ctx             context.Context
	questionService *services.QuestionService
	commandsService *services.CommandsService
}

func sendMarshalledData[T any](obj *T, w http.ResponseWriter) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(data)
}

func unmarshalBody[T any](w http.ResponseWriter, r *http.Request) *T {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid body"))
		return nil
	}

	return &req
}

func (s *server) AddEmployee(w http.ResponseWriter, r *http.Request) {
	req := unmarshalBody[repository.AddTasks](w, r)
	if req == nil {
		return
	}
	err := container.Container.UserRepo().AddTasks(s.ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) GetRequestsHandler(w http.ResponseWriter, _ *http.Request) {
	questions, err := s.questionService.GetQuestions(s.ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sendMarshalledData(&questions, w)
}

func (s *server) AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	req := unmarshalBody[services.AnswerQuestionRequest](w, r)
	if req == nil {
		return
	}

	err := s.questionService.AnswerQuestion(s.ctx, req)

	if errors.Is(err, repository.ErrQuestionNotExist) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) GetCommands(w http.ResponseWriter, _ *http.Request) {
	commands, err := s.commandsService.GetCommands(s.ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sendMarshalledData(&commands, w)
}

func (s *server) ChangeCommand(w http.ResponseWriter, r *http.Request) {
	req := unmarshalBody[services.UpdateCommandRequest](w, r)
	if req == nil {
		return
	}

	err := s.commandsService.UpdateCommand(s.ctx, req)
	if errors.Is(err, repository.ErrNoMaterial) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("material does not exist"))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) AddCommand(w http.ResponseWriter, r *http.Request) {
	req := unmarshalBody[services.AddCommandRequest](w, r)
	if req == nil {
		return
	}
	err := s.commandsService.AddCommand(s.ctx, req)
	if errors.Is(err, repository.ErrTxFail) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ =
			w.Write([]byte("the command name is unique or action id is not found"))
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewSendMessageRequest(w http.ResponseWriter, r *http.Request) *services.SendMessageRequest {
	var req services.SendMessageRequest

	err := r.ParseMultipartForm(config.Cfg.Admin.MaxPhotoSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("can't process"))
		return nil
	}

	// Read photo.
	file, _, err := r.FormFile(config.Cfg.Admin.PhotoFormKey)
	if err == nil {
		req.Photo, err = io.ReadAll(file)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		}
	}

	// Read the message.
	message := r.FormValue(config.Cfg.Admin.MessageFormKey)
	if message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	req.Message = message

	return &req
}

func (s *server) SendMessage(w http.ResponseWriter, r *http.Request) {
	req := NewSendMessageRequest(w, r)
	if req == nil {
		return
	}

	err := services.SendMessage(s.ctx, req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
