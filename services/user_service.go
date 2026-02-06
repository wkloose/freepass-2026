package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/Hisyam/freepass-2026/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginInput struct {
	Email    string
	Password string
}

type UpdateProfileInput struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
}

type UserService interface {
	Register(input RegisterInput) (*models.User, error)
	Login(input LoginInput) (string, error)
	GetProfile(id uuid.UUID) (*models.User, error)
	UpdateProfile(id uuid.UUID, input UpdateProfileInput) (*models.User, error)
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
		Role:     "CUSTOMER",
	}

	err = s.repository.Create(&newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email is registered")
		}
		return nil, err
	}

	return &newUser, nil
}

func (s *userService) Login(input LoginInput) (string, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID.String(),
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * time.Duration(utils.GetEnvAsInt("JWT_EXPIRY_HOUR", 12))).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func (s *userService) GetProfile(id uuid.UUID) (*models.User, error) {
	return s.repository.FindByID(id)
}

func (s *userService) UpdateProfile(id uuid.UUID, input UpdateProfileInput) (*models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email

	err = s.repository.Update(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email has been used by another user")
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUserByAdmin(input CreateUserInput) (*models.User, error) {
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

	err = s.repository.Create(&newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email has been used by another user")
		}
		return nil, err
	}

	return &newUser, nil
}
