package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Borrowing struct {
	pb.UnimplementedBorrowingHandlerServer
	mu               sync.Mutex
	serviceBorrowing service.Borrowing
}

// mustEmbedUnimplementedBorrowingHandlerServer implements grpc.BorrowingHandlerServer.
func (u *Borrowing) mustEmbedUnimplementedBorrowingHandlerServer() {
	panic("unimplemented")
}

func NewBorrowing(serviceBorrowing service.Borrowing) *Borrowing {
	return &Borrowing{serviceBorrowing: serviceBorrowing}
}

func (u *Borrowing) Borrow(ctx context.Context, in *pb.BorrowRequest) (*emptypb.Empty, error) {
	borrow := &model.Borrowing{
		BookID: in.BookId,
		UserID: in.UserId,
		Amount: in.Amount,
	}
	_, err := u.serviceBorrowing.Borrow(ctx, borrow)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Borrowing) Return(ctx context.Context, in *pb.ReturnRequest) (*emptypb.Empty, error) {
	borrow := &model.Borrowing{
		ID:     in.BorrowingId,
		UserID: in.UserId,
	}
	_, err := u.serviceBorrowing.Return(ctx, borrow)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
