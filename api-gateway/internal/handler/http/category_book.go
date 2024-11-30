package http

import (
	"net/http"
	"os"
	"strconv"

	pbCategoryBook "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/category_book/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/middleware"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryBook struct {
	categoryBookClient pbCategoryBook.CategoryBookHandlerClient
}

func NewCategoryBook() *CategoryBook {
	return &CategoryBook{}
}

func (c *CategoryBook) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient(os.Getenv("URL_BOOK"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.categoryBookClient = pbCategoryBook.NewCategoryBookHandlerClient(conn)
	router.POST("/", middleware.JWTMiddleware(), c.Add)
	router.DELETE("/:id", middleware.JWTMiddleware(), c.Delete)

	return nil
}

func (c *CategoryBook) Add(ctx *gin.Context) {
	var input dto.CategoryBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	_, err := c.categoryBookClient.Add(ctx, &pbCategoryBook.AddCategoryBookRequest{
		CategoryId: input.CategoryID,
		BookId:     input.BookID,
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Category Book created successfully")
}

func (c *CategoryBook) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = c.categoryBookClient.Delete(ctx, &pbCategoryBook.DeleteCategoryBookRequest{
		Id: uint32(id),
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Category Book deleted successfully")
}
