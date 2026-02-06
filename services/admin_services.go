package services

import (
	"errors"
	"strings"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=ADMIN CANTEEN CUSTOMER"`
}

type UpdateUserByAdminInput struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required,oneof=ADMIN CANTEEN CUSTOMER"`
}

type AdminService interface {
	CreateUser(input CreateUserInput) (*models.User, error)
	UpdateUser(id uuid.UUID, input UpdateUserByAdminInput) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type adminService struct {
	userRepository repositories.UserRepository
}

func NewAdminService(userRepository repositories.UserRepository) AdminService {
	return &adminService{userRepository}
}

func (s *adminService) CreateUser(input CreateUserInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     models.UserRole(input.Role),
	}

	err = s.userRepository.Create(&newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email has been used by another user")
		}
		return nil, err
	}

	return &newUser, nil
}

func (s *adminService) UpdateUser(id uuid.UUID, input UpdateUserByAdminInput) (*models.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = models.UserRole(input.Role)

	err = s.userRepository.Update(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email has been used by another user")
		}
		return nil, err
	}

	return user, nil
}

func (s *adminService) DeleteUser(id uuid.UUID) error {
	_, err := s.userRepository.FindByID(id)
    if err != nil {
        return errors.New("user not found")
    }
	return s.userRepository.Delete(id)
}