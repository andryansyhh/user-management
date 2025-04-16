package usecase

import (
	"context"
	"errors"
	"time"
	"user-management/internal/repository"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, token string) error
}

type authUsecase struct {
	repo  repository.AuthRepository
	redis *redis.Client
}

func NewAuthUsecase(r repository.AuthRepository, redis *redis.Client) AuthUsecase {
	return &authUsecase{
		repo:  r,
		redis: redis,
	}
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	token := uuid.NewString()

	// store token in Redis
	err = u.redis.Set(ctx, token, user.ID, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *authUsecase) Logout(ctx context.Context, token string) error {
	return u.redis.Del(ctx, token).Err()
}
