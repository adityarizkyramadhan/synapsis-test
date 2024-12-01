package service

import (
	"context"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/repository"
)

type RecommendationService interface {
	GetRecommendationUserByAuthor(ctx context.Context, userID uint32) ([]model.Book, error)
	GetRecommendationUserByCategory(ctx context.Context, userID uint32) ([]model.Book, error)
	GetRecommendationUserByTitle(ctx context.Context, userID uint32) ([]model.Book, error)
}

type recommendationService struct {
	repo repository.RecommendationRepository
}

func NewRecommendation(repo repository.RecommendationRepository) RecommendationService {
	return &recommendationService{repo}
}

func (s *recommendationService) GetRecommendationUserByAuthor(ctx context.Context, userID uint32) ([]model.Book, error) {
	return s.repo.GetRecommendationUserByAuthor(ctx, userID)
}

func (s *recommendationService) GetRecommendationUserByCategory(ctx context.Context, userID uint32) ([]model.Book, error) {
	return s.repo.GetRecommendationUserByCategory(ctx, userID)
}

func (s *recommendationService) GetRecommendationUserByTitle(ctx context.Context, userID uint32) ([]model.Book, error) {
	return s.repo.GetRecommendationUserByTitle(ctx, userID)
}
