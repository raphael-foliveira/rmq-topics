package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, status int, content interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if status == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(w).Encode(content)
}
