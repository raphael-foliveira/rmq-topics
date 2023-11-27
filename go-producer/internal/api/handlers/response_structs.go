package handlers

type healthCheck struct {
	Status string `json:"status"`
}

type standardMessage struct {
	Message string `json:"message"`
}
