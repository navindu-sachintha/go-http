package api

import (
	"encoding/json"
	"fem/go/http/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
