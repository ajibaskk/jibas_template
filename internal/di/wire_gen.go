// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"gorm.io/gorm"
	"jibas-template/external/email"
	"jibas-template/internal/delivery/http"
	"jibas-template/internal/repository"
	"jibas-template/internal/usecase"
)

// Injectors from wire.go:

func InitializeUserHandler(db *gorm.DB) http.UserHandlerInterface {
	userRepository := repository.NewUserRepository(db)
	emailService := email.NewEmailService()
	userUsecase := usecase.NewUserUsecase(userRepository, emailService)
	userHandlerInterface := http.NewUserHandler(userUsecase)
	return userHandlerInterface
}