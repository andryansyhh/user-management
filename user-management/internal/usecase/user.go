package usecase

import (
	"context"
	"user-management/internal/domain/model"
	"user-management/internal/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]*model.User, error) {
	return u.repo.Fetch(ctx)
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.User) error {
	return u.repo.Store(ctx, user)
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *model.User) error {
	
	return u.repo.Update(ctx, user)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
