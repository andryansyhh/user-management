package handler

import (
	"context"

	"user-management/internal/usecase"
	pbAuth "user-management/proto/auth"

	"google.golang.org/grpc"
)

type AuthHandler struct {
	pbAuth.UnimplementedAuthServiceServer
	uc usecase.AuthUsecase
}

func RegisterAuthService(server grpc.ServiceRegistrar, uc usecase.AuthUsecase) {
	pbAuth.RegisterAuthServiceServer(server, &AuthHandler{uc: uc})
}

func (h *AuthHandler) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	token, err := h.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return &pbAuth.LoginResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &pbAuth.LoginResponse{
		Status:  true,
		Message: "Successfully",
		Data: &pbAuth.LoginResponse_Data{
			AccessToken: token,
		},
	}, nil
}

func (h *AuthHandler) Logout(ctx context.Context, req *pbAuth.LogoutRequest) (*pbAuth.GenericResponse, error) {
	err := h.uc.Logout(ctx, req.AccessToken)
	if err != nil {
		return &pbAuth.GenericResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &pbAuth.GenericResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
