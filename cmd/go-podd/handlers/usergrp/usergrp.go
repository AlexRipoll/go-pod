package usergrp

import (
	"context"
	"errors"
	"github.com/AlexRipoll/go-pod/internal/core/user"
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"github.com/AlexRipoll/go-pod/web"
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

	var nu user.NewUser
	if err := web.Decode(r, &nu); err != nil {
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
