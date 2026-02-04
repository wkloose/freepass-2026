package services

import (
	"errors"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     models.UserRole
}

type LoginInput struct {
	Email    string
	Password string
}

type UserService interface {
	Register(input RegisterInput) (*models.User, error)
	Login(input LoginInput) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) Register(input RegisterInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	err = s.repository.Create(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (s *userService) Login(input LoginInput) (*models.User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}