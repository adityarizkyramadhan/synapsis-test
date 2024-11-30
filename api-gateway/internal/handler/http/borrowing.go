package http

import (
	"net/http"
	"strconv"

	pbBorrowing "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/borrowing/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/middleware"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Borrowing struct {
	borrowingClient pbBorrowing.BorrowingHandlerClient
}

func NewBorrowingRoutes() *Borrowing {
	return &Borrowing{}
}

func (b *Borrowing) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	b.borrowingClient = pbBorrowing.NewBorrowingHandlerClient(conn)
	router.POST("/", middleware.JWTMiddleware(), b.Borrow)
	router.DELETE("/:id", middleware.JWTMiddleware(), b.Return)

	return nil
}

func (b *Borrowing) Borrow(ctx *gin.Context) {
	var input dto.BorrowingInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	userId := ctx.MustGet("id").(string)
	userIdInt, _ := strconv.ParseUint(userId, 10, 64)

	_, err := b.borrowingClient.Borrow(ctx, &pbBorrowing.BorrowRequest{
		BookId: input.BookID,
		UserId: uint32(userIdInt),
		Amount: input.Amount,
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Borrowing created successfully")
}

func (b *Borrowing) Return(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	userId := ctx.MustGet("id").(string)
	userIdInt, _ := strconv.ParseUint(userId, 10, 64)

	_, err = b.borrowingClient.Return(ctx, &pbBorrowing.ReturnRequest{
		BorrowingId: uint32(id),
		UserId:      uint32(userIdInt),
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Borrowing returned successfully")
}
