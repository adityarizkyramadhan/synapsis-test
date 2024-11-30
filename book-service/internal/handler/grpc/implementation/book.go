package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Book struct {
	pb.UnimplementedBookHandlerServer
	mu          sync.Mutex
	serviceUser service.BookService
}

func NewBook(serviceUser service.BookService) *Book {
	return &Book{serviceUser: serviceUser}
}

func (u *Book) GetByID(ctx context.Context, ID *pb.GetByIDRequest) (*pb.Book, error) {
	book, err := u.serviceUser.GetByID(ctx, ID.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Book{
		Id:          book.ID,
		Title:       book.Title,
		AuthorId:    book.AuthorID,
		Description: book.Description,
		Year:        book.Year,
		Stock:       book.Stock,
		CreatedAt:   timestamppb.New(book.CreatedAt),
		UpdatedAt:   timestamppb.New(book.UpdatedAt),
	}, nil
}

func (u *Book) Create(ctx context.Context, in *pb.Book) (*emptypb.Empty, error) {
	book := &model.Book{
		Title:       in.Title,
		AuthorID:    in.AuthorId,
		Description: in.Description,
		Year:        in.Year,
		Stock:       in.Stock,
	}
	_, err := u.serviceUser.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Book) Update(ctx context.Context, in *pb.UpdateBookRequest) (*emptypb.Empty, error) {
	book := &model.Book{
		ID:          in.Id,
		Title:       in.Book.Title,
		AuthorID:    in.Book.AuthorId,
		Description: in.Book.Description,
		Year:        in.Book.Year,
		Stock:       in.Book.Stock,
	}
	_, err := u.serviceUser.Update(ctx, book)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Book) Delete(ctx context.Context, ID *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	err := u.serviceUser.Delete(ctx, ID.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Book) ListAll(ctx context.Context, _ *emptypb.Empty) (*pb.ListBooksResponse, error) {
	books, err := u.serviceUser.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []*pb.Book
	for _, book := range books {
		data = append(data, &pb.Book{
			Id:          book.ID,
			Title:       book.Title,
			AuthorId:    book.AuthorID,
			Description: book.Description,
			Year:        book.Year,
			Stock:       book.Stock,
			CreatedAt:   timestamppb.New(book.CreatedAt),
			UpdatedAt:   timestamppb.New(book.UpdatedAt),
		})
	}
	return &pb.ListBooksResponse{Books: data}, nil
}
