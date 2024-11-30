package http

import (
	"os"

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
	conn, err := grpc.NewClient(os.Getenv("URL_USER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	u.client = pb.NewUserHandlerClient(conn)

	router.POST("/register", u.Create)
	router.POST("/login", u.Login)

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

func (u *UserRoutes) Login(ctx *gin.Context) {
	var input dto.UserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, 400, err)
		return
	}

	user, err := u.client.Login(ctx, &pb.User{
		Password: input.Password,
		Email:    input.Email,
	})

	if err != nil {
		utils.ResponseError(ctx, 500, err)
		return
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		utils.ResponseError(ctx, 500, err)
		return
	}

	utils.ResponseSuccess(ctx, 200, token)
}
