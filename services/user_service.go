package services

import (
	"errors"
	"os"
	"time"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/Hisyam/freepass-2026/repositories"
	"github.com/Hisyam/freepass-2026/utils"
	"github.com/golang-jwt/jwt/v5"
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
	Login(input LoginInput) (string, error)
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
