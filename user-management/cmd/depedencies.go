package cmd

import (
	"log"
	"user-management/internal/handler"
	"user-management/internal/repository"
	"user-management/internal/usecase"

	"google.golang.org/grpc"
)

type Dependency struct {
	Config *Config
}

func InitDependencies(server *grpc.Server) *Dependency {
	cfg, err := Load()
	if err != nil {
		log.Println("error to load")
	}

	// External dependencies
	dbConn, err := NewClientDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	redisClient := NewClientRedis(cfg)

	// Auth module
	authRepository := repository.NewAuthRepository(dbConn)
	authUsecase := usecase.NewAuthUsecase(authRepository, redisClient)
	handler.RegisterAuthService(server, authUsecase)

	// User module
	userRepository := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	handler.RegisterUserService(server, userUsecase)

	return &Dependency{Config: cfg}
}
