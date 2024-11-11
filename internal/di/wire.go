//go:build wireinject
// +build wireinject

package di

import (
	"jibas-template/external/email"
	"jibas-template/internal/delivery/http"
	"jibas-template/internal/repository"
	"jibas-template/internal/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) http.UserHandlerInterface {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		http.NewUserHandler,
		email.NewEmailService,
	)
	return nil
}
