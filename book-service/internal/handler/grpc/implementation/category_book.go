package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryBook struct {
	pb.UnimplementedCategoryBookHandlerServer
	mu              sync.Mutex
	serviceCategory service.CategoryBookService
}

func NewCategoryBook(serviceCategory service.CategoryBookService) *CategoryBook {
	return &CategoryBook{serviceCategory: serviceCategory}
}

func (u *CategoryBook) Add(ctx context.Context, in *pb.AddCategoryBookRequest) (*emptypb.Empty, error) {
	category := &model.CategoryBook{
		CategoryID: in.CategoryId,
		BookID:     in.BookId,
	}
	_, err := u.serviceCategory.Add(ctx, category)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *CategoryBook) Delete(ctx context.Context, in *pb.DeleteCategoryBookRequest) (*emptypb.Empty, error) {
	err := u.serviceCategory.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
