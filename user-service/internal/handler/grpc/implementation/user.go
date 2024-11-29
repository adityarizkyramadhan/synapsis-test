package implementation

import (
	"context"
	"sync"

	pb "github.com/adityarizkyramadhan/synapsis-test/user-service/internal/handler/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/service"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	pb.UnimplementedUserHandlerServer
	mu          sync.Mutex
	serviceUser service.UserService
}

func NewUser(serviceUser service.UserService) *User {
	return &User{serviceUser: serviceUser}
}

func (u *User) GetByID(ctx context.Context, ID *pb.GetByIDRequest) (*pb.User, error) {
	user, err := u.serviceUser.GetByID(ctx, ID.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
func (u *User) Create(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:    in.Email,
		Password: string(password),
	}

	err = u.serviceUser.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (u *User) Update(ctx context.Context, in *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:    in.User.Email,
		Password: string(password),
	}

	err = u.serviceUser.Update(ctx, in.Id, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}
func (u *User) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	err := u.serviceUser.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
