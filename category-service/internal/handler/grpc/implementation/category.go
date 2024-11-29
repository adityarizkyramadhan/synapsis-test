package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/category-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/service"
)

type Category struct {
	pb.UnimplementedCategoryHandlerServer
	mu            sync.Mutex
	serviceAuthor service.CategoryService
}

func NewCategory(serviceAuthor service.CategoryService) pb.CategoryHandlerServer {
	return &Category{serviceAuthor: serviceAuthor}
}

func (u *Category) GetByID(ctx context.Context, arg *pb.GetByIDRequest) (*pb.Category, error) {
	category, err := u.serviceAuthor.GetByID(ctx, uint32(arg.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:        category.ID,
		Name:      category.Name,
		CreatedAt: timestamppb.New(category.CreatedAt),
		UpdatedAt: timestamppb.New(category.UpdatedAt),
	}, nil
}

func (u *Category) Create(ctx context.Context, arg *pb.Category) (*emptypb.Empty, error) {
	category := &model.Category{
		Name: arg.Name,
	}
	err := u.serviceAuthor.Create(ctx, category)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Category) Update(ctx context.Context, arg *pb.UpdateCategoryRequest) (*emptypb.Empty, error) {
	category := &model.Category{
		ID:   uint32(arg.Id),
		Name: arg.Category.Name,
	}
	err := u.serviceAuthor.Update(ctx, category)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Category) Delete(ctx context.Context, arg *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := u.serviceAuthor.Delete(ctx, uint32(arg.Id))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *Category) ListAll(ctx context.Context, arg *emptypb.Empty) (*pb.ListCategoriesResponse, error) {
	categories, err := u.serviceAuthor.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:        category.ID,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		})
	}

	return &pb.ListCategoriesResponse{Categories: pbCategories}, nil
}
