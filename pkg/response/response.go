package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, resp any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
