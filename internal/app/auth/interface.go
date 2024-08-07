package auth

import (
	"context"
	"github.com/raffops/auth/internal/app/auth/model"
	"github.com/raffops/auth/internal/app/user/models"
	"github.com/raffops/chat/pkg/errs"
	"net/http"
)

type Controller interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

type Service interface {
	SignUp(
		ctx context.Context,
		username, email string,
		authType user.AuthTypeId,
		role auth.RoleId,
	) (string, errs.ChatError)
	Login(ctx context.Context, username, email string) (string, errs.ChatError)
	Refresh(ctx context.Context, token string) errs.ChatError
}
