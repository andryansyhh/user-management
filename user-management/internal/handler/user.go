package handler

import (
	"context"
	"strconv"

	"user-management/internal/domain/model"
	"user-management/internal/usecase"
	pbUser "user-management/proto/user"

	"google.golang.org/grpc"
)

type UserHandler struct {
	pbUser.UnimplementedUserServiceServer
	uc usecase.UserUsecase
}

func RegisterUserService(server grpc.ServiceRegistrar, uc usecase.UserUsecase) {
	pbUser.RegisterUserServiceServer(server, &UserHandler{uc: uc})
}

func (h *UserHandler) GetAllUsers(ctx context.Context, _ *pbUser.Empty) (*pbUser.UserListResponse, error) {
	users, err := h.uc.GetUsers(ctx)
	if err != nil {
		return &pbUser.UserListResponse{Status: false, Message: err.Error()}, nil
	}

	var res []*pbUser.UserData

	for _, u := range users {
		res = append(res, &pbUser.UserData{
			Id:         strconv.FormatInt(u.ID, 10),
			Name:       u.Name,
			Email:      u.Email,
			RoleId:     u.RoleID,
			RoleName:   u.RoleName,
			LastAccess: u.LastAccess,
		})
	}

	return &pbUser.UserListResponse{
		Status:  true,
		Message: "Success",
		Data:    res,
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.GenericResponse, error) {
	err := h.uc.CreateUser(ctx, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleID:   req.RoleId,
	})
	if err != nil {
		return &pbUser.GenericResponse{Status: false, Message: err.Error()}, nil
	}
	return &pbUser.GenericResponse{Status: true, Message: "Successfully"}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.GenericResponse, error) {
	if req.Id == "" {
		return &pbUser.GenericResponse{Status: false, Message: "User ID is required"}, nil
	}

	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return &pbUser.GenericResponse{Status: false, Message: "Invalid ID format"}, nil
	}

	err = h.uc.UpdateUser(ctx, &model.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return &pbUser.GenericResponse{Status: false, Message: err.Error()}, nil
	}

	return &pbUser.GenericResponse{Status: true, Message: "Successfully"}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.GenericResponse, error) {
	newId, _ := strconv.Atoi(req.Id)
	err := h.uc.DeleteUser(ctx, int64(newId))
	if err != nil {
		return &pbUser.GenericResponse{Status: false, Message: err.Error()}, nil
	}
	return &pbUser.GenericResponse{Status: true, Message: "Successfully"}, nil
}
