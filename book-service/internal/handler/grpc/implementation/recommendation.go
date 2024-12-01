package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/service"
)

type Recommendation struct {
	pb.UnimplementedRecommendationHandlerServer
	mu           sync.Mutex
	serviceRecom service.RecommendationService
}

func NewRecommendation(serviceRecom service.RecommendationService) *Recommendation {
	return &Recommendation{serviceRecom: serviceRecom}
}

func (u *Recommendation) GetRecommendationUserByAuthor(ctx context.Context, ID *pb.GetRecommendationRequest) (*pb.RecommendationResponse, error) {
	book, err := u.serviceRecom.GetRecommendationUserByAuthor(ctx, ID.UserId)
	if err != nil {
		return nil, err
	}
	var books []*pb.BookRecommendation
	for _, b := range book {
		books = append(books, &pb.BookRecommendation{
			Id:          b.ID,
			Title:       b.Title,
			AuthorId:    b.AuthorID,
			Description: b.Description,
			Year:        b.Year,
		})
	}
	return &pb.RecommendationResponse{
		Books: books,
	}, nil
}

func (u *Recommendation) GetRecommendationUserByCategory(ctx context.Context, ID *pb.GetRecommendationRequest) (*pb.RecommendationResponse, error) {
	book, err := u.serviceRecom.GetRecommendationUserByCategory(ctx, ID.UserId)
	if err != nil {
		return nil, err
	}
	var books []*pb.BookRecommendation
	for _, b := range book {
		books = append(books, &pb.BookRecommendation{
			Id:          b.ID,
			Title:       b.Title,
			AuthorId:    b.AuthorID,
			Description: b.Description,
			Year:        b.Year,
		})
	}
	return &pb.RecommendationResponse{
		Books: books,
	}, nil
}

func (u *Recommendation) GetRecommendationUserByTitle(ctx context.Context, ID *pb.GetRecommendationRequest) (*pb.RecommendationResponse, error) {
	book, err := u.serviceRecom.GetRecommendationUserByTitle(ctx, ID.UserId)
	if err != nil {
		return nil, err
	}
	var books []*pb.BookRecommendation
	for _, b := range book {
		books = append(books, &pb.BookRecommendation{
			Id:          b.ID,
			Title:       b.Title,
			AuthorId:    b.AuthorID,
			Description: b.Description,
			Year:        b.Year,
		})
	}
	return &pb.RecommendationResponse{
		Books: books,
	}, nil
}
