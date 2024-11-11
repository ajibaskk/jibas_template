package usecase

import "jibas-template/internal/domain"

type userUsecase struct {
	userRepo     domain.UserRepository
	emailService domain.EmailService
}

func NewUserUsecase(userRepo domain.UserRepository, emailService domain.EmailService) domain.UserUsecase {
	return &userUsecase{
		userRepo:     userRepo,
		emailService: emailService,
	}
}

// CreateUser implements domain.UserUsecase.
func (u *userUsecase) CreateUser(user *domain.User) error {
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}

	// Send a welcome email after user creation
	return u.emailService.SendEmail(user.Email, "Welcome!", "Thank you for signing up!")
}

// GetAllUsers implements domain.UserUsecase.
func (u *userUsecase) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}
