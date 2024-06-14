package auth

import (
	"context"
	"github.com/raffops/auth/internal/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, key, value string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, key, value string) error
	ListUsers(ctx context.Context) ([]models.User, error)
}
