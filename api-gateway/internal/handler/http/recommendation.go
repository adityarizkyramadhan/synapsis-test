package http

import (
	"net/http"
	"os"

	pbRecom "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/recommendation/grpc"
	"github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RecommendationRoutes struct {
	RecommendationHandler pbRecom.RecommendationHandlerClient
}

func NewRecommendationRoutes() *RecommendationRoutes {
	return &RecommendationRoutes{}
}

func (r *RecommendationRoutes) Init(router *gin.RouterGroup) error {
	conn, err := grpc.NewClient(os.Getenv("URL_RECOMMENDATION"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	r.RecommendationHandler = pbRecom.NewRecommendationHandlerClient(conn)

	router.GET("/author", r.GetRecommendationUserByAuthor)
	router.GET("/category", r.GetRecommendationUserByCategory)
	router.GET("/title", r.GetRecommendationUserByTitle)

	return nil
}

func (r *RecommendationRoutes) GetRecommendationUserByAuthor(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(float64)
	userIDInt := int64(userID)
	_, err := r.RecommendationHandler.GetRecommendationUserByAuthor(ctx, &pbRecom.GetRecommendationRequest{UserId: uint32(userIDInt)})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Recommendation user by author")
}

func (r *RecommendationRoutes) GetRecommendationUserByCategory(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(float64)
	userIDInt := int64(userID)
	_, err := r.RecommendationHandler.GetRecommendationUserByCategory(ctx, &pbRecom.GetRecommendationRequest{UserId: uint32(userIDInt)})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Recommendation user by category")
}

func (r *RecommendationRoutes) GetRecommendationUserByTitle(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(float64)
	userIDInt := int64(userID)
	_, err := r.RecommendationHandler.GetRecommendationUserByTitle(ctx, &pbRecom.GetRecommendationRequest{UserId: uint32(userIDInt)})
	if err != nil {
		utils.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Recommendation user by book")
}
