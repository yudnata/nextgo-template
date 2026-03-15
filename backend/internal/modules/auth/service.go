package auth

import (
	"context"
	"errors"
	"time"

	"backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
}

func NewService(repo *Repository, cfg *config.Config) *Service {
	return &Service{repo: repo, cfg: cfg}
}

func (s *Service) Register(input RegisterReq) (*User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	user := &User{
		ID:        uuid.New().String(),
		Email:     input.Email,
		Name:      input.Name,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}
	return user, s.repo.Create(context.Background(), user)
}

func (s *Service) Login(input LoginReq) (string, error) {
	user, err := s.repo.FindByEmail(context.Background(), input.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return "", errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.cfg.JWTSecret))
}
