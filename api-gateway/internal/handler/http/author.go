package http

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	pbAuthor "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/author/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/middleware"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthorRoutes struct {
	authorClient pbAuthor.AuthorHandlerClient
}

func NewAuthorRoutes() *AuthorRoutes {
	return &AuthorRoutes{}
}

func (a *AuthorRoutes) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient(os.Getenv("URL_AUTHOR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	a.authorClient = pbAuthor.NewAuthorHandlerClient(conn)

	router.POST("/", middleware.JWTMiddleware(), a.Create)
	router.PUT("/:id", middleware.JWTMiddleware(), a.Update)
	router.DELETE("/:id", middleware.JWTMiddleware(), a.Delete)
	router.GET("/:id", middleware.JWTMiddleware(), a.GetByID)
	router.GET("/", middleware.JWTMiddleware(), a.GetAll)
	return nil
}

func (a *AuthorRoutes) Create(ctx *gin.Context) {
	var input dto.AuthorInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	_, err := a.authorClient.Create(ctx, &pbAuthor.Author{
		Name:  input.Name,
		Email: input.Email,
		Bio:   input.Bio,
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Author created successfully")
}

func (a *AuthorRoutes) Update(ctx *gin.Context) {
	var input dto.AuthorInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	idString := ctx.Param("id")
	if idString == "" {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("author id is required"))
		return
	}

	idUint, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = a.authorClient.Update(ctx, &pbAuthor.UpdateAuthorRequest{
		Id: uint32(idUint),
		Author: &pbAuthor.Author{
			Name:  input.Name,
			Email: input.Email,
			Bio:   input.Bio,
		},
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Author updated successfully")
}

func (a *AuthorRoutes) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	if idString == "" {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("author id is required"))
		return
	}

	idUint, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = a.authorClient.Delete(ctx, &pbAuthor.DeleteAuthorRequest{Id: uint32(idUint)})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Author deleted successfully")
}

func (a *AuthorRoutes) GetByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	if idString == "" {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("author id is required"))
		return
	}

	idUint, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	author, err := a.authorClient.GetByID(ctx, &pbAuthor.GetByIDRequest{Id: uint32(idUint)})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	authorOutput := dto.AuthorOutput{
		ID:        author.Id,
		Name:      author.Name,
		Email:     author.Email,
		Bio:       author.Bio,
		CreatedAt: author.CreatedAt.AsTime().Format(time.RFC3339),
		UpdatedAt: author.UpdatedAt.AsTime().Format(time.RFC3339),
	}

	utils.ResponseSuccess(ctx, http.StatusOK, authorOutput)
}

func (a *AuthorRoutes) GetAll(ctx *gin.Context) {
	authors, err := a.authorClient.ListAll(ctx, nil)
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	var author []dto.AuthorOutput
	for _, v := range authors.Authors {
		author = append(author, dto.AuthorOutput{
			ID:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			Bio:       v.Bio,
			CreatedAt: v.CreatedAt.AsTime().Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.AsTime().Format(time.RFC3339),
		})
	}
	utils.ResponseSuccess(ctx, http.StatusOK, author)
}
