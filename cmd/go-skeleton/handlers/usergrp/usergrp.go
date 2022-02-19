package usergrp

import (
	"context"
	"encoding/json"
	"github.com/AlexRipoll/go-skeleton/internal/core/user"
	"io"
	"net/http"
)

type Handler struct {
	User user.Core 	
}


func (h Handler) Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid body provided"))
		return
	}

	var nu user.NewUser
	if err := json.Unmarshal(body, &nu); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid body provided"))
		return
	}

	u, err := h.User.Create(context.Background(), nu)
	if err != nil {
		w.WriteHeader(000)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resp))
	return
}