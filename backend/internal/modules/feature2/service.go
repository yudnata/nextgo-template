package feature2

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(input CreateReq) (*Data, error) {
	data := &Data{
		ID:        uuid.New().String(),
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}
	return data, s.repo.Create(context.Background(), data)
}
