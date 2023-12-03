package handlers

import (
	"encoding/json"
	"go-producer/internal/api/res"
	"go-producer/internal/producer"
	"io"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return res.Json(w, http.StatusOK, healthCheck{Status: "ok"})
}

type queueHandler struct {
	publisher *producer.Publisher
}

func NewQueueHandler(queueProducer *producer.Publisher) *queueHandler {
	return &queueHandler{queueProducer}
}

type topicOneMessage struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (qh *queueHandler) TopicOne(w http.ResponseWriter, r *http.Request) error {
	contentBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return &res.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	err = qh.publisher.Publish(&producer.Message{
		TopicName: "topic.one",
		Content:   contentBytes,
	})
	if err != nil {
		return &res.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return res.Json(w, http.StatusOK, standardMessage{
		Message: "sent!",
	})
}

type topicTwoMessage struct {
	Log              string `json:"log"`
	CriticalityLevel string `json:"criticality_level"`
}

func (qh *queueHandler) TopicTwo(w http.ResponseWriter, r *http.Request) error {
	content := topicTwoMessage{}
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		return &res.HttpError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	err = qh.publisher.Publish(&producer.Message{
		TopicName: "topic.two",
		Content:   []byte("this is a message to topic one"),
	})
	if err != nil {
		return &res.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return res.Json(w, http.StatusOK, standardMessage{
		Message: "sent!",
	})
}
