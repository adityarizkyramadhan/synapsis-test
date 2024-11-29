package http

import (
	pb "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/user/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserRoutes struct {
	client pb.UserHandlerClient
}

func NewUserRoutes() *UserRoutes {
	return &UserRoutes{}
}

func (u *UserRoutes) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	u.client = pb.NewUserHandlerClient(conn)

	router.POST("/", u.Create)

	return nil
}

func (u *UserRoutes) Create(ctx *gin.Context) {
	var input dto.UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, 400, err)
		return
	}

	_, err := u.client.Create(ctx, &pb.User{
		Password: input.Password,
		Email:    input.Email,
	})

	if err != nil {
		utils.ResponseError(ctx, 500, err)
		return
	}

	utils.ResponseSuccess(ctx, 200, "User created")
}
