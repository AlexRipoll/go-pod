package usergrp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AlexRipoll/go-pod/internal/core/user"
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"github.com/AlexRipoll/go-pod/web"
	"io"
	"net/http"
)

var (
	errInvalidData = errors.New("invalid body provided")
	errInternal    = errors.New("internal server error")
)

type Handler struct {
	User user.Core
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		web.ErrorResponse(ctx, w, errorFlag.New(errInternal, errorFlag.Internal))
		return
	}

	var nu user.NewUser
	if err := json.Unmarshal(body, &nu); err != nil {
		web.ErrorResponse(ctx, w, errorFlag.New(errInvalidData, errorFlag.InvalidData))
		return
	}

	u, err := h.User.Create(context.Background(), nu)
	if err != nil {
		web.ErrorResponse(ctx, w, err)
		return
	}

	web.Response(ctx, w, u, http.StatusCreated)

	return
}
