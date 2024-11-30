package http

import (
	"errors"
	"net/http"
	"strconv"

	pbCategory "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/category/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/dto"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/middleware"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryRoutes struct {
	categoryClient pbCategory.CategoryHandlerClient
}

func NewCategoryRoutes() *CategoryRoutes {
	return &CategoryRoutes{}
}

func (c *CategoryRoutes) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	c.categoryClient = pbCategory.NewCategoryHandlerClient(conn)
	router.POST("/", middleware.JWTMiddleware(), c.Create)
	router.PUT("/:id", middleware.JWTMiddleware(), c.Update)
	router.DELETE("/:id", middleware.JWTMiddleware(), c.Delete)
	router.GET("/:id", middleware.JWTMiddleware(), c.GetByID)
	router.GET("/", middleware.JWTMiddleware(), c.GetAll)

	return nil
}

func (c *CategoryRoutes) Create(ctx *gin.Context) {
	var input dto.CategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	_, err := c.categoryClient.Create(ctx, &pbCategory.Category{
		Name: input.Name,
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, "Category created successfully")
}

func (c *CategoryRoutes) Update(ctx *gin.Context) {
	var input dto.CategoryInput
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
	_, err = c.categoryClient.Update(ctx, &pbCategory.UpdateCategoryRequest{
		Id: uint32(idInt),
		Category: &pbCategory.Category{
			Name: input.Name,
		},
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Category updated successfully")
}

func (c *CategoryRoutes) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	_, err = c.categoryClient.Delete(ctx, &pbCategory.DeleteCategoryRequest{
		Id: uint32(idInt),
	})

	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Category deleted successfully")
}

func (c *CategoryRoutes) GetAll(ctx *gin.Context) {
	res, err := c.categoryClient.ListAll(ctx, nil)
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, res)
}

func (c *CategoryRoutes) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseError(ctx, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	res, err := c.categoryClient.GetByID(ctx, &pbCategory.GetByIDRequest{
		Id: uint32(idInt),
	})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, res)
}
