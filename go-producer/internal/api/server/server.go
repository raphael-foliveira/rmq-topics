package server

import (
	"encoding/json"
	"fmt"
	"go-producer/internal/api/handlers"
	"go-producer/internal/producer"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type server struct {
	port int
}

func NewServer(port int) *server {
	return &server{port}
}

type HttpError struct {
	Status  int
	Message string
}

func (e *HttpError) Error() string {
	return e.Message
}

func (s *server) Start(queueProducer *producer.Publisher) error {
	port := fmt.Sprintf(":%d", s.port)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)

	router.Get("/", wrapHandler(handlers.HealthCheck))

	queueHandler := handlers.NewQueueHandler(queueProducer)
	router.Get("/topic-one", wrapHandler(queueHandler.TopicOne))
	router.Get("/topic-two", wrapHandler(queueHandler.TopicTwo))

	log.Println("Server listening on port", port)
	return http.ListenAndServe(port, router)
}

func wrapHandler(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			httpError, ok := err.(*HttpError)
			if ok {
				w.WriteHeader(httpError.Status)
				json.NewEncoder(w).Encode(map[string]string{"error": httpError.Message})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			return
		}
	}
}
