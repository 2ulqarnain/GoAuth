package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("auth ok"))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("username: %s\n", payload.Email)
	fmt.Printf("password: %s\n", payload.Password)
	w.Write([]byte("Logged in successfully!"))
}
