package handlers

import (
	"go-producer/internal/api/res"
	"go-producer/internal/producer"
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

func (qh *queueHandler) TopicOne(w http.ResponseWriter, r *http.Request) error {
	qh.publisher.Publish(&producer.Message{
		TopicName: "topic.one",
		Content:   []byte("this is a message to topic one"),
	})
	return res.Json(w, http.StatusOK, standardMessage{
		Message: "sent!",
	})
}

func (qh *queueHandler) TopicTwo(w http.ResponseWriter, r *http.Request) error {
	qh.publisher.Publish(&producer.Message{
		TopicName: "topic.two",
		Content:   []byte("this is a message to topic one"),
	})
	return res.Json(w, http.StatusOK, standardMessage{
		Message: "sent!",
	})
}
