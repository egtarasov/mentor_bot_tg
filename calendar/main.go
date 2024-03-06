package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Meeting struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	StartTime   time.Duration `json:"start_time"`
}

const userIdParam = "id"

func main() {
	mux := http.ServeMux{}
	meetings := []Meeting{
		{
			Name:        "Стендап",
			Description: "Ежедневный стендап комманды",
			StartTime:   time.Hour * 13,
		},
		{
			Name:        "Созвон с поставщиком по",
			Description: "Обсудить вопросики",
			StartTime:   time.Hour * 18,
		},
	}
	data, err := json.Marshal(&meetings)
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("GET /calendar", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		if q.Get(userIdParam) == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		_, _ = writer.Write(data)
	})

	if err := http.ListenAndServe(":8000", &mux); err != nil {
		log.Println(err)
	}
}
