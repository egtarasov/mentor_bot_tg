package admin

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"telegrambot_new_emploee/internal/config"
	"telegrambot_new_emploee/internal/repository"
)

var TempId int64 = 1

func StartServer() {
	mux := http.NewServeMux()
	s := server{
		ctx:             context.Background(),
		questionService: NewQuestionService(),
		commandsService: NewCommandsService(),
	}

	mux.HandleFunc("GET /questions", s.GetRequestsHandler)
	mux.HandleFunc("POST /questions", s.AnswerQuestion)

	mux.HandleFunc("GET /commands", s.GetCommands)
	mux.HandleFunc("PUT /commands", s.ChangeCommandMaterial)
	mux.HandleFunc("POST /commands", s.AddCommand)

	if err := http.ListenAndServe(config.Cfg.Admin.Port, mux); err != nil {
		log.Println(err)
	}
}

type server struct {
	ctx             context.Context
	questionService *QuestionService
	commandsService *CommandsService
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
	req := unmarshalBody[AnswerQuestionRequest](w, r)
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

func (s *server) ChangeCommandMaterial(w http.ResponseWriter, r *http.Request) {
	req := unmarshalBody[UpdateMaterialRequest](w, r)
	if req == nil {
		return
	}

	err := s.commandsService.UpdateMaterial(s.ctx, req)
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
	req := unmarshalBody[AddCommandRequest](w, r)
	if req == nil {
		return
	}
	err := s.commandsService.AddCommand(s.ctx, req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
