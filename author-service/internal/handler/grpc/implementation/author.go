package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/author-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/author-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/author-service/internal/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Author struct {
	pb.UnimplementedAuthorHandlerServer
	mu            sync.Mutex
	serviceAuthor service.AuthorService
}

func NewAuthor(serviceAuthor service.AuthorService) pb.AuthorHandlerServer {
	return &Author{serviceAuthor: serviceAuthor}
}

func (u *Author) GetByID(ctx context.Context, arg *pb.GetByIDRequest) (*pb.Author, error) {
	author, err := u.serviceAuthor.ReadByID(ctx, int(arg.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Author{
		Id:        author.ID,
		Name:      author.Name,
		CreatedAt: timestamppb.New(author.CreatedAt),
		UpdatedAt: timestamppb.New(author.UpdatedAt),
	}, nil
}

func (u *Author) Create(ctx context.Context, arg *pb.Author) (*emptypb.Empty, error) {
	author := &model.Author{
		Name:  arg.Name,
		Email: arg.Email,
		Bio:   arg.Bio,
	}
	err := u.serviceAuthor.Create(ctx, author)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Author) Update(ctx context.Context, arg *pb.UpdateAuthorRequest) (*emptypb.Empty, error) {
	author := &model.Author{
		ID:    arg.Id,
		Name:  arg.Author.Name,
		Bio:   arg.Author.Bio,
		Email: arg.Author.Email,
	}
	err := u.serviceAuthor.Update(ctx, author)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Author) Delete(ctx context.Context, arg *pb.DeleteAuthorRequest) (*emptypb.Empty, error) {
	_, err := u.serviceAuthor.ReadByID(ctx, int(arg.Id))
	if err != nil {
		return nil, err
	}
	err = u.serviceAuthor.Delete(ctx, int(arg.Id))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Author) ListAll(ctx context.Context, in *emptypb.Empty) (*pb.ListAuthorsResponse, error) {
	authors, err := u.serviceAuthor.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	var authorResponses []*pb.Author
	for _, author := range authors {
		authorResponses = append(authorResponses, &pb.Author{
			Id:        author.ID,
			Name:      author.Name,
			CreatedAt: timestamppb.New(author.CreatedAt),
			UpdatedAt: timestamppb.New(author.UpdatedAt),
		})
	}
	return &pb.ListAuthorsResponse{Authors: authorResponses}, nil
}
