package http

import (
	"errors"
	"net/http"
	"strconv"

	pbBook "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/book/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/middleware"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookRoutes struct {
	bookClient pbBook.BookHandlerClient
}

func NewBookRoutes() *BookRoutes {
	return &BookRoutes{}
}

func (b *BookRoutes) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	b.bookClient = pbBook.NewBookHandlerClient(conn)
	router.POST("/", middleware.JWTMiddleware(), b.Create)
	router.PUT("/:id", middleware.JWTMiddleware(), b.Update)
	router.DELETE("/:id", middleware.JWTMiddleware(), b.Delete)
	router.GET("/:id", middleware.JWTMiddleware(), b.GetByID)
	router.GET("/", middleware.JWTMiddleware(), b.GetAll)

	return nil
}

func (b *BookRoutes) Create(ctx *gin.Context) {
	var input dto.BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	_, err := b.bookClient.Create(ctx, &pbBook.Book{
		Title:       input.Title,
		AuthorId:    input.AuthorID,
		Description: input.Description,
		Year:        input.Year,
		Stock:       input.Stock,
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Book created successfully")
}

func (b *BookRoutes) Update(ctx *gin.Context) {
	var input dto.BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	_, err = b.bookClient.Update(ctx, &pbBook.UpdateBookRequest{
		Id: uint32(idInt),
		Book: &pbBook.Book{
			Title:       input.Title,
			AuthorId:    input.AuthorID,
			Description: input.Description,
			Year:        input.Year,
			Stock:       input.Stock,
		},
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Book updated successfully")
}

func (b *BookRoutes) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	_, err = b.bookClient.Delete(ctx, &pbBook.DeleteBookRequest{
		Id: uint32(idInt),
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Book deleted successfully")
}

func (b *BookRoutes) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	book, err := b.bookClient.GetByID(ctx, &pbBook.GetByIDRequest{
		Id: uint32(idInt),
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	bookDto := dto.BookOutput{
		ID:          book.Id,
		Title:       book.Title,
		AuthorID:    book.AuthorId,
		Description: book.Description,
		Year:        book.Year,
		Stock:       book.Stock,
		CreatedAt:   book.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
		UpdatedAt:   book.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
	}

	utils.ResponseSuccess(ctx, http.StatusOK, bookDto)
}

func (b *BookRoutes) GetAll(ctx *gin.Context) {
	res, err := b.bookClient.ListAll(ctx, nil)
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	var books []*dto.BookOutput
	for _, book := range res.Books {
		books = append(books, &dto.BookOutput{
			ID:          book.Id,
			Title:       book.Title,
			AuthorID:    book.AuthorId,
			Description: book.Description,
			Year:        book.Year,
			Stock:       book.Stock,
			CreatedAt:   book.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
			UpdatedAt:   book.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
		})
	}

	utils.ResponseSuccess(ctx, http.StatusOK, books)
}
